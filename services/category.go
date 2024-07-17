package services

import (
	"context"
	"errors"
	"strings"

	"github.com/google/uuid"
	"github.com/yosa12978/mdpages/logging"
	"github.com/yosa12978/mdpages/repos"
	"github.com/yosa12978/mdpages/types"
)

type CategoryService interface {
	GetSubcategories(ctx context.Context, parentId string) []types.Category
	GetById(ctx context.Context, id string) (*types.Category, error)
	GetRoots(ctx context.Context) []types.Category

	Create(ctx context.Context, dto types.CategoryCreateDto) error
	Update(ctx context.Context, id, name, parent_id string) error
	Delete(ctx context.Context, id string) error

	Seed(ctx context.Context) error
}

type categoryService struct {
	categoryRepo repos.CategoryRepo
	logger       logging.Logger
}

func NewCategoryService(
	categoryRepo repos.CategoryRepo,
	logger logging.Logger,
) CategoryService {
	return &categoryService{
		categoryRepo: categoryRepo,
		logger:       logger,
	}
}

func (c *categoryService) GetSubcategories(ctx context.Context, parentId string) []types.Category {
	categories, err := c.categoryRepo.GetSubcategories(ctx, parentId)
	if err != nil {
		c.logger.Error(err.Error())
		return nil
	}
	return categories
}

func (c *categoryService) GetById(ctx context.Context, id string) (*types.Category, error) {
	return c.categoryRepo.GetById(ctx, id)
}

func (c *categoryService) GetRoots(ctx context.Context) []types.Category {
	categories, err := c.categoryRepo.GetRootCategories(ctx)
	if err != nil {
		c.logger.Error(err.Error())
		return nil
	}
	return categories
}

func (c *categoryService) Update(ctx context.Context, id, name, parent_id string) error { // change this to CategoryUpdateDto
	return c.categoryRepo.Update(ctx, id, types.Category{Id: id, Name: name, ParentId: parent_id})
}

func (c *categoryService) Delete(ctx context.Context, id string) error {
	return c.categoryRepo.Delete(ctx, id)
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
	if err := c.categoryRepo.Create(ctx, types.Category{
		Id:       "39495288-dd70-43e5-9531-a57c24d0f3a4",
		Name:     "Main Category",
		ParentId: "",
	}); err != nil {
		return err
	}
	return c.categoryRepo.Create(ctx, types.Category{
		Id:       "7b34ea11-57c3-46f0-80ef-e35e743d2889",
		Name:     "Subcategory 1",
		ParentId: "39495288-dd70-43e5-9531-a57c24d0f3a4",
	})
}
