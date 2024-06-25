package repos

import (
	"context"
)

type CRUD[T any] interface {
	GetById(ctx context.Context, id string) (*T, error)
	GetAll(ctx context.Context) ([]T, error)
	Create(ctx context.Context, entity T) error
	Update(ctx context.Context, id string, entity T) error
	Delete(ctx context.Context, id string) error
}
