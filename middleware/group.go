package middleware

import (
	"net/http"

	"github.com/yosa12978/mdpages/session"
)

func GroupFilter(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, err := session.GetSession(r)
		if err != nil {
			http.Error(w, "forbidden", http.StatusForbidden)
			return
		}
		w.Header().Set("Cache-Control", "no-cache")
		h(w, r)
	}
}

func RootOnly(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		usr, err := session.GetSession(r)
		if err != nil {
			http.Error(w, "forbidden", http.StatusForbidden)
			return
		}
		w.Header().Set("Cache-Control", "no-cache")
		for _, v := range usr.Groups {
			if v.Id == "root" {
				h(w, r)
				return
			}
		}
		http.Error(w, "forbidden", http.StatusForbidden)
	}
}
