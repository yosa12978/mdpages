package app

import (
	"context"
	"net/http"
	"os"

	"github.com/yosa12978/mdpages/config"
	"github.com/yosa12978/mdpages/data"
	"github.com/yosa12978/mdpages/handler"
	"github.com/yosa12978/mdpages/logging"
	"github.com/yosa12978/mdpages/middleware"
	"github.com/yosa12978/mdpages/repos"
	"github.com/yosa12978/mdpages/services"
	"github.com/yosa12978/mdpages/session"
	"github.com/yosa12978/mdpages/util"
)

func NewRouter(ctx context.Context) http.Handler {
	router := http.NewServeMux()
	cfg := config.Get()

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
	if err := accountService.Seed(ctx, cfg.App.RootPassword); err != nil {
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

	SetupRoutes(router)

	authHandler := handler.NewAuthHandler(accountService, logger)
	authHandler.Setup(router)

	articleHandler := handler.NewArticleHandler(articleService, logger)
	articleHandler.Setup(router)

	categoryHandler := handler.NewCategoryHandler(categoryService, articleService)
	categoryHandler.Setup(router)

	return router
}

func SetupRoutes(router *http.ServeMux) {
	router.Handle("GET /assets/", http.StripPrefix("/assets", http.FileServer(http.Dir("assets"))))

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			w.WriteHeader(http.StatusNotFound)
			util.RenderView(w, r, "notfound", nil)
			return
		}
		util.RenderView(w, r, "home", nil)
	})

	router.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(200)
		w.Write([]byte("healthy"))
	})

	router.HandleFunc("GET /login", middleware.AnonymousOnly(
		func(w http.ResponseWriter, r *http.Request) {
			util.RenderView(w, r, "login", nil)
		}))

	router.HandleFunc("GET /signup", middleware.AnonymousOnly(
		func(w http.ResponseWriter, r *http.Request) {
			util.RenderView(w, r, "signup", nil)
		}))

	router.HandleFunc("GET /hello", func(w http.ResponseWriter, r *http.Request) {
		usr, err := session.GetSession(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		util.RenderView(w, r, "hello", usr.Username)
	})

	router.HandleFunc("GET /categories/", func(w http.ResponseWriter, r *http.Request) {
		util.RenderView(w, r, "articles", nil)
	})

	router.HandleFunc("GET /article/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		util.RenderView(w, r, "detail", id)
	})

	router.HandleFunc("GET /categories/{id}", func(w http.ResponseWriter, r *http.Request) {
		util.RenderView(w, r, "articles", r.PathValue("id"))
	})
}
