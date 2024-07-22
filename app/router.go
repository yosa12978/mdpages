package app

import (
	"context"
	"net/http"
	"os"

	"github.com/yosa12978/mdpages/data"
	"github.com/yosa12978/mdpages/logging"
	"github.com/yosa12978/mdpages/repos"
	"github.com/yosa12978/mdpages/services"
	"github.com/yosa12978/mdpages/view"
)

func NewRouter(ctx context.Context) http.Handler {
	router := http.NewServeMux()

	accountRepo := repos.NewAccountRepo(data.Postgres())
	articleRepo := repos.NewArticleRepo(data.Postgres())
	commitRepo := repos.NewCommitRepo(data.Postgres())
	categoryRepo := repos.NewCategoryRepo(data.Postgres())
	groupRepo := repos.NewGroupRepo(data.Postgres())

	logger := logging.NewLogger(os.Stdout)

	groupService := services.NewGroupService(groupRepo, logger)
	accountService := services.NewAccountService(accountRepo, groupService, logger)
	commitService := services.NewCommitService(commitRepo, logger)
	articleService := services.NewArticleService(articleRepo, logger)
	categoryService := services.NewCategoryService(categoryRepo, logger)

	// Seeding
	if err := groupService.Seed(ctx); err != nil {
		logger.Error(err.Error())
	}
	if err := accountService.Seed(ctx, os.Getenv("ROOT_PASSWORD")); err != nil {
		logger.Error(err.Error())
	}
	if err := categoryService.Seed(ctx); err != nil {
		logger.Error(err.Error())
	}
	if err := articleService.Seed(ctx); err != nil {
		logger.Error(err.Error())
	}
	if err := commitService.Seed(ctx); err != nil {
		logger.Error(err.Error())
	}

	router.Handle("/assets/", http.StripPrefix("/assets", http.FileServer(http.Dir("assets"))))

	router.HandleFunc("GET /{$}", func(w http.ResponseWriter, r *http.Request) {
		view.Index("mdpages - home").Render(ctx, w)
	})

	router.HandleFunc("GET /login", func(w http.ResponseWriter, r *http.Request) {
		view.Login().Render(ctx, w)
	})

	router.HandleFunc("GET /signup", func(w http.ResponseWriter, r *http.Request) {
		view.Signup().Render(ctx, w)
	})

	router.HandleFunc("GET /htmx/home", func(w http.ResponseWriter, r *http.Request) {
		article, err := articleService.GetHomePage(r.Context())
		if err != nil {
			logger.Error(err.Error())
			http.Error(w, err.Error(), 500)
			return
		}
		view.Page(*article).Render(ctx, w)
	})

	return router
}
