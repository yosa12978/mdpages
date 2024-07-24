package app

import (
	"context"
	"net/http"
	"os"

	"github.com/yosa12978/mdpages/data"
	"github.com/yosa12978/mdpages/handler"
	"github.com/yosa12978/mdpages/logging"
	"github.com/yosa12978/mdpages/middleware"
	"github.com/yosa12978/mdpages/repos"
	"github.com/yosa12978/mdpages/services"
	"github.com/yosa12978/mdpages/session"
	"github.com/yosa12978/mdpages/types"
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
		usr, _ := session.GetSession(r)
		view.Index(types.TemplData{
			Title: "Home",
			User:  usr,
		}).Render(ctx, w)
	})

	router.HandleFunc("GET /login", middleware.AnonymousOnly(
		func(w http.ResponseWriter, r *http.Request) {
			usr, _ := session.GetSession(r)
			if usr != nil {
				http.Redirect(w, r, "/", http.StatusPermanentRedirect)
				return
			}
			view.Login(types.TemplData{
				Title: "Login",
				User:  usr,
			}).Render(ctx, w)
		}))

	router.HandleFunc("GET /signup", middleware.AnonymousOnly(
		func(w http.ResponseWriter, r *http.Request) {
			usr, _ := session.GetSession(r)
			if usr != nil {
				http.Redirect(w, r, "/", http.StatusPermanentRedirect)
				return
			}
			view.Signup(types.TemplData{
				Title: "Signup",
				User:  usr,
			}).Render(ctx, w)
		}))

	router.HandleFunc("GET /hello", func(w http.ResponseWriter, r *http.Request) {
		usr, err := session.GetSession(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		view.Hello(usr.Username, types.TemplData{
			Title: "Hello Page",
			User:  usr,
		}).Render(ctx, w)
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

	// place to another spot
	authHandler := handler.NewAuthHandler(accountService, logger)
	router.HandleFunc("POST /htmx/login",
		middleware.AnonymousOnly(
			handler.MakeHandler(
				authHandler.Login(),
			),
		),
	)
	router.HandleFunc("POST /htmx/signup",
		middleware.AnonymousOnly(
			handler.MakeHandler(
				authHandler.Signup(),
			),
		),
	)
	router.HandleFunc("POST /htmx/logout", handler.MakeHandler(authHandler.Logout()))

	// make a middleware for groups and etc i.e. only for anon users

	return router
}

func SetupRouter() {

}
