package handler

import (
	"net/http"
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
			// if e, ok := err.(ErrorResp); ok {
			// 	http.Error(w, e.Error(), e.Status)
			// 	return
			// }
			// make something different here (may be add logging or idk)
			//http.Error(w, "418 I guess server is a teapot", http.StatusTeapot)
			http.Error(w, err.Error(), 200)
		}
	}
}
