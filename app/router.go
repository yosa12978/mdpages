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

	logger := logging.NewLoggerWithSource(os.Stdout)

	accountService := services.NewAccountService(accountRepo, logger)
	commitService := services.NewCommitService(commitRepo, logger)
	articleService := services.NewArticleService(articleRepo, logger)
	categoryService := services.NewCategoryService(categoryRepo, logger)

	if err := accountService.Seed(ctx); err != nil {
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

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		view.Index("world").Render(ctx, w)
	})
	return router
}
