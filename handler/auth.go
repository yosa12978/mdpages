package handler

import (
	"encoding/json"
	"net/http"

	"github.com/yosa12978/mdpages/logging"
	"github.com/yosa12978/mdpages/middleware"
	"github.com/yosa12978/mdpages/services"
	"github.com/yosa12978/mdpages/session"
	"github.com/yosa12978/mdpages/types"
)

type AuthHandler interface {
	Setup(router *http.ServeMux)
	Login() Handler
	Signup() Handler
	Logout() Handler
}

type authHandler struct {
	accountService services.AccountService
	logger         logging.Logger
}

func NewAuthHandler(
	accountService services.AccountService,
	logger logging.Logger,
) AuthHandler {
	return &authHandler{
		accountService: accountService,
		logger:         logger,
	}
}

func (a *authHandler) Setup(router *http.ServeMux) {
	router.HandleFunc("POST /htmx/login",
		middleware.AnonymousOnly(
			MakeHandler(
				a.Login(),
			),
		),
	)
	router.HandleFunc("POST /htmx/signup",
		middleware.AnonymousOnly(
			MakeHandler(
				a.Signup(),
			),
		),
	)
	router.HandleFunc("POST /htmx/logout",
		MakeHandler(
			a.Logout(),
		),
	)
}

func (h *authHandler) Logout() Handler {
	return func(w http.ResponseWriter, r *http.Request) error {
		if err := session.EndSession(r, w); err != nil {
			return err
		}
		w.Header().Set("HX-Redirect", "/")
		return nil
	}
}

func (h *authHandler) Login() Handler {
	return func(w http.ResponseWriter, r *http.Request) error {
		body := make(map[string]interface{})
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			return err
		}
		account, err := h.accountService.GetByCredentials(
			r.Context(),
			body["username"].(string),
			body["password"].(string),
		)
		if err != nil {
			return err
		}
		if err := session.StartSession(r, w, *account); err != nil {
			return err
		}
		w.Header().Set("HX-Redirect", "/hello")
		return nil
	}
}

func (h *authHandler) Signup() Handler {
	return func(w http.ResponseWriter, r *http.Request) error {
		body := types.AccountCreateDto{}
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			return err
		}
		if err := h.accountService.Create(r.Context(), body); err != nil {
			return err
		}
		w.Header().Set("HX-Redirect", "/login")
		return nil
	}
}
