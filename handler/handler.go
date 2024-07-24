package handler

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type Handler func(w http.ResponseWriter, r *http.Request) error

type ErrorResp struct {
	Status int
	Msg    string
}

func (e ErrorResp) Error() string {
	return e.Msg
}

func MakeHandler(h Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := h(w, r); err != nil {
			if e, ok := err.(ErrorResp); ok {
				http.Error(w, e.Error(), e.Status)
				return
			}
			// make something different here (may be add logging or idk)
			//http.Error(w, "418 I guess server is a teapot", http.StatusTeapot)
			if e, ok := err.(validator.ValidationErrors); ok {
				fmt.Fprintf(w, "%s", e.Error())
				return
			}
			w.WriteHeader(200)
			w.Write([]byte(err.Error()))
		}
	}
}
