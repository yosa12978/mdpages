package services

import "context"

type CommitService interface {
	Create(ctx context.Context, articleId, author, title, body string) error
	Seed(ctx context.Context) error
}
