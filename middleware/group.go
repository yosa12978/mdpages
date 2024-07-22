package middleware

import (
	"net/http"

	"github.com/yosa12978/mdpages/session"
)

func GroupFilter(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sessionInfo, err := session.GetInfo(r)
		if err != nil || sessionInfo == nil {
			http.Error(w, "forbidden", http.StatusForbidden)
			return
		}
		w.Header().Set("Cache-Control", "no-cache")
		next.ServeHTTP(w, r)
	})
}
