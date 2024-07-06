package services

import "context"

type AccountService interface {
	Create(ctx context.Context, category, title, body, author string) error
	Seed(ctx context.Context) error
}
