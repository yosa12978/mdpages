package services

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/yosa12978/mdpages/repos"
	"github.com/yosa12978/mdpages/types"
)

type CommitService interface {
	Create(ctx context.Context, dto types.CommitCreateDto) error
	Seed(ctx context.Context) error
}

type commitService struct {
	commitRepo repos.CommitRepo
}

func NewCommitService(commitRepo repos.CommitRepo) CommitService {
	return &commitService{
		commitRepo: commitRepo,
	}
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
	panic("unimplemented")
}
