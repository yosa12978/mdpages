package services

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/yosa12978/mdpages/logging"
	"github.com/yosa12978/mdpages/repos"
	"github.com/yosa12978/mdpages/types"
)

type CommitService interface {
	GetById(ctx context.Context, id string) (*types.Commit, error)
	GetArticleCommits(ctx context.Context, articleId string) ([]types.Commit, error)

	Create(ctx context.Context, dto types.CommitCreateDto) error
	Delete(ctx context.Context, id string) error

	Seed(ctx context.Context) error
}

type commitService struct {
	commitRepo repos.CommitRepo
	logger     logging.Logger
}

func NewCommitService(
	commitRepo repos.CommitRepo,
	logger logging.Logger,
) CommitService {
	return &commitService{
		commitRepo: commitRepo,
		logger:     logger,
	}
}

func (c *commitService) GetById(ctx context.Context, id string) (*types.Commit, error) {
	return c.commitRepo.GetById(ctx, id)
}
func (c *commitService) GetArticleCommits(ctx context.Context, articleId string) ([]types.Commit, error) {
	return c.commitRepo.GetArticleCommits(ctx, articleId)
}

func (c *commitService) Delete(ctx context.Context, id string) error {
	return c.commitRepo.Delete(ctx, id)
}

// Create implements CommitService.
func (c *commitService) Create(ctx context.Context, dto types.CommitCreateDto) error {
	return c.commitRepo.Create(ctx, types.Commit{
		Id:        uuid.NewString(),
		ArticleId: dto.ArticleId,
		Title:     dto.Title,
		Body:      dto.Body,
		Author:    dto.Author,
		Created:   time.Now().Format(time.RFC3339),
	})
}

// Seed implements CommitService.
func (c *commitService) Seed(ctx context.Context) error {
	return nil
}
