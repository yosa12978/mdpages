package services

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/yosa12978/mdpages/repos"
	"github.com/yosa12978/mdpages/types"
)

type ArticleService interface {
	Create(ctx context.Context, dto types.ArticleCreateDto) error
	Seed(ctx context.Context) error
}

type articleService struct {
	articleRepo repos.ArticleRepo
}

func NewArticleService(articleRepo repos.ArticleRepo) ArticleService {
	return &articleService{
		articleRepo: articleRepo,
	}
}

// Create implements ArticleService.
func (a *articleService) Create(ctx context.Context, dto types.ArticleCreateDto) error {
	return a.articleRepo.Create(ctx, types.Article{
		Id:            uuid.NewString(),
		CategoryId:    dto.CategoryId,
		Title:         dto.Title,
		Body:          dto.Body,
		CommitCreated: time.Now().Format(time.RFC3339),
		CommitId:      uuid.NewString(),
		CommitAuthor:  dto.Author,
	})
}

// Seed implements ArticleService.
func (a *articleService) Seed(ctx context.Context) error {
	if err := a.articleRepo.Create(ctx, types.Article{
		Id:            "8f5e42b3-9645-4715-b1d2-9ee0b7ae2d74",
		CategoryId:    "39495288-dd70-43e5-9531-a57c24d0f3a4",
		Title:         "first article",
		Body:          "new test body 123",
		CommitCreated: time.Now().Format(time.RFC3339),
		CommitId:      "3aa9c6a6-576b-4cd2-9fbb-eabdb1ee93f8",
		CommitAuthor:  "user",
	}); err != nil {
		return err
	}
	return a.articleRepo.Create(ctx, types.Article{
		Id:            "3dc29d82-f60c-4fdf-9d9c-5b135a30765e",
		CategoryId:    "7b34ea11-57c3-46f0-80ef-e35e743d2889",
		Title:         "second article",
		Body:          "new test body 456",
		CommitCreated: time.Now().Format(time.RFC3339),
		CommitId:      "9d942708-2ae1-44e4-944c-7b6abead63c8",
		CommitAuthor:  "user",
	})
}
