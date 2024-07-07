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
	Create(ctx context.Context, dto types.CategoryCreateDto) error
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
func (c *categoryService) Create(ctx context.Context, dto types.CategoryCreateDto) error {
	nameTrimmed := strings.TrimSpace(dto.Name)
	if nameTrimmed == "" || len(nameTrimmed) > 40 {
		return errors.New("len(name) is either >40 or =0")
	}
	// check here if parent category exist
	return c.categoryRepo.Create(ctx, types.Category{
		Id:       uuid.NewString(),
		Name:     nameTrimmed,
		ParentId: dto.ParentId,
	})
}

// Seed implements CategoryService.
func (c *categoryService) Seed(ctx context.Context) error {
	mainId := uuid.NewString()
	childId := uuid.NewString()
	if err := c.categoryRepo.Create(ctx, types.Category{
		Id:       mainId,
		Name:     "Main Category",
		ParentId: "",
	}); err != nil {
		return err
	}
	err := c.categoryRepo.Create(ctx, types.Category{
		Id:       childId,
		Name:     "Subcategory 1",
		ParentId: mainId,
	})
	return err
}
