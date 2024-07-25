package handler

import (
	"net/http"

	"github.com/yosa12978/mdpages/logging"
	"github.com/yosa12978/mdpages/services"
	"github.com/yosa12978/mdpages/view"
)

type ArticleHandler interface {
	Setup(router *http.ServeMux)
	GetArticles() Handler
	GetArticleById() Handler
	GetHomePage() Handler
}

type articleHandler struct {
	articleService services.ArticleService
	logger         logging.Logger
}

func NewArticleHandler(
	articleService services.ArticleService,
	logger logging.Logger,
) ArticleHandler {
	return &articleHandler{
		articleService: articleService,
		logger:         logger,
	}
}

func (a *articleHandler) Setup(router *http.ServeMux) {
	router.HandleFunc("/htmx/pages",
		MakeHandler(a.GetArticles()),
	)
	router.HandleFunc("/htmx/pages/{id}",
		MakeHandler(a.GetArticleById()),
	)
	router.HandleFunc("GET /htmx/home",
		MakeHandler(a.GetHomePage()),
	)
}

func (a *articleHandler) GetHomePage() Handler {
	return func(w http.ResponseWriter, r *http.Request) error {
		article, err := a.articleService.GetHomePage(r.Context())
		if err != nil {
			return err
		}
		view.Page(*article).Render(r.Context(), w)
		return err
	}
}

// GetArticles implements ArticleHandler.
func (a *articleHandler) GetArticles() Handler {
	return func(w http.ResponseWriter, r *http.Request) error {
		articles := a.articleService.GetUncategorized(r.Context())
		view.Pages(articles).Render(r.Context(), w)
		return nil
	}
}

// GetArticlesById implements ArticleHandler.
func (a *articleHandler) GetArticleById() Handler {
	return func(w http.ResponseWriter, r *http.Request) error {
		id := r.PathValue("id")
		article, err := a.articleService.GetById(r.Context(), id)
		if err != nil {
			return err
		}
		view.Page(*article).Render(r.Context(), w)
		return nil
	}
}
