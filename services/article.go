package services

import (
	"context"

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
	panic("unimplemented")
}

// Seed implements ArticleService.
func (a *articleService) Seed(ctx context.Context) error {
	panic("unimplemented")
}
