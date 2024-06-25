package app

import (
	"context"
	"net/http"

	"github.com/yosa12978/mdpages/view"
)

func NewRouter(ctx context.Context) http.Handler {
	router := http.NewServeMux()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		view.Index("me").Render(ctx, w)
	})
	return router
}
