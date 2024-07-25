package handler

import (
	"net/http"

	"github.com/yosa12978/mdpages/services"
	"github.com/yosa12978/mdpages/util"
)

type CategoryHandler interface {
	Setup(router *http.ServeMux)
	GetRootCategories() Handler
	GetSubcategories() Handler
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

func (c *categoryHandler) Setup(router *http.ServeMux) {
	router.HandleFunc("GET /htmx/categories", MakeHandler(c.GetRootCategories()))
}

func (c *categoryHandler) GetRootCategories() Handler {
	return func(w http.ResponseWriter, r *http.Request) error {
		categories := c.categoryService.GetRoots(r.Context())
		if len(categories) == 0 {
			return nil
		}
		return util.RenderBlock(w, "categories", categories)
	}
}

func (c *categoryHandler) GetSubcategories() Handler {
	return func(w http.ResponseWriter, r *http.Request) error {
		return nil
	}
}
