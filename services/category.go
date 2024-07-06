package services

import (
	"context"
	"errors"
	"strings"

	"github.com/google/uuid"
	"github.com/yosa12978/mdpages/repos"
	"github.com/yosa12978/mdpages/types"
)

type CategoryService interface {
	Seed(ctx context.Context) error
	Create(ctx context.Context, name string) error
}

type categoryService struct {
	categoryRepo repos.CategoryRepo
}

func NewCategoryService(categoryRepo repos.CategoryRepo) CategoryService {
	return &categoryService{
		categoryRepo: categoryRepo,
	}
}

// Create implements CategoryService.
func (c *categoryService) Create(ctx context.Context, name string) error {
	if strings.TrimSpace(name) == "" || len(name) > 40 {
		return errors.New("len(name) is either >40 or =0")
	}
	return c.categoryRepo.Create(ctx, types.Category{
		Id:   uuid.NewString(),
		Name: name,
	})
}

// Seed implements CategoryService.
func (c *categoryService) Seed(ctx context.Context) error {
	c.Create(ctx, "first category")
	c.Create(ctx, "second category")
	c.Create(ctx, "third category")
	c.Create(ctx, "fourth category")
	return nil
}
