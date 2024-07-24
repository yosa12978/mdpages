package middleware

import (
	"net/http"

	"github.com/yosa12978/mdpages/session"
)

func AnonymousOnly(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, err := session.GetSession(r)
		if err == nil {
			http.Redirect(w, r, "/", http.StatusPermanentRedirect)
			return
		}
		w.Header().Set("Cache-Control", "no-cache")
		h(w, r)
	}
}
