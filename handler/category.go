package handler

import (
	"net/http"

	"github.com/yosa12978/mdpages/services"
	"github.com/yosa12978/mdpages/util"
)

type CategoryHandler interface {
	Setup(router *http.ServeMux)
	GetCategories() Handler
}

type categoryHandler struct {
	categoryService services.CategoryService
	articleService  services.ArticleService
}

func NewCategoryHandler(
	categoryService services.CategoryService,
	articleService services.ArticleService,
) CategoryHandler {
	return &categoryHandler{
		categoryService: categoryService,
		articleService:  articleService,
	}
}

func (c *categoryHandler) Setup(router *http.ServeMux) { // remake setup handler to take interface as an argument and be just a regular func not struct method
	router.HandleFunc("GET /htmx/categories", MakeHandler(c.GetCategories()))
	router.HandleFunc("GET /htmx/categories/{parentId}", MakeHandler(c.GetCategories()))
}

func (c *categoryHandler) GetCategories() Handler {
	return func(w http.ResponseWriter, r *http.Request) error {
		parentId := r.PathValue("parentId")
		payload := make(map[string]any)
		parent, err := c.categoryService.GetById(r.Context(), parentId)
		if err != nil && parentId != "" {
			return err
		}
		payload["Parent"] = parent
		categories := c.categoryService.GetCategories(r.Context(), parentId)
		payload["Categories"] = categories
		return util.RenderBlock(w, "categories", payload)
	}
}
