package repos

import (
	"context"
	"database/sql"

	"github.com/yosa12978/mdpages/types"
	"github.com/yosa12978/mdpages/util"
)

type CategoryRepo interface {
	GetRootCategories(ctx context.Context) ([]types.Category, error)
	GetSubcategories(ctx context.Context, id string) ([]types.Category, error)
	Create(ctx context.Context, entity types.Category) error
	Delete(ctx context.Context, id string) error
	GetById(ctx context.Context, id string) (*types.Category, error)
	Update(ctx context.Context, id string, entity types.Category) error
	GetByName(ctx context.Context, name string) (*types.Category, error)
}

type categoryRepo struct {
	db *sql.DB
}

func NewCategoryRepo(db *sql.DB) CategoryRepo {
	return &categoryRepo{
		db: db,
	}
}

func (c *categoryRepo) GetSubcategories(ctx context.Context, id string) ([]types.Category, error) {
	q := `
		WITH RECURSIVE subcategories AS (
			SELECT id, parent_id, 0 AS depth FROM categories 
			WHERE id=$1
			
		UNION ALL
			
			SELECT c.id, c.parent_id, s.depth+1 
			FROM categories c 
			INNER JOIN subcategories s
			ON c.parent_id = s.id
		)
		SELECT c.id, c.name FROM subcategories s 
		INNER JOIN categories c 
		ON c.id = s.id
		WHERE s.depth = 1;
	`
	row, err := c.db.QueryContext(ctx, q, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return []types.Category{}, types.NewErrNotFound("categories not found")
		}
		return nil, err
	}
	categories := []types.Category{}
	for row.Next() {
		category := types.Category{}
		row.Scan(&category.Id, &category.Name)
		category.ParentId = id
		categories = append(categories, category)
	}
	return categories, nil
}

func (c *categoryRepo) GetRootCategories(ctx context.Context) ([]types.Category, error) {
	q := `
		SELECT id, name FROM categories WHERE parent_id IS NULL;
	`
	row, err := c.db.QueryContext(ctx, q)
	if err != nil {
		if err == sql.ErrNoRows {
			return []types.Category{}, types.NewErrNotFound("categories not found")
		}
		return nil, err
	}
	categories := []types.Category{}
	for row.Next() {
		category := types.Category{}
		row.Scan(
			&category.Id,
			&category.Name,
		)
		categories = append(categories, category)
	}
	return categories, nil
}

func (c *categoryRepo) Create(ctx context.Context, entity types.Category) error {
	q := `
		INSERT INTO categories (id, name, parent_id) VALUES ($1, $2, $3);
	`
	_, err := c.db.ExecContext(ctx, q,
		entity.Id,
		entity.Name,
		util.NewNullString(entity.ParentId))
	return err
}

func (c *categoryRepo) Delete(ctx context.Context, id string) error {
	q := `
		DELETE FROM categories WHERE id=$1;
	`
	_, err := c.db.ExecContext(ctx, q, id)
	return err
}

func (c *categoryRepo) GetById(ctx context.Context, id string) (*types.Category, error) {
	q := `
		SELECT id, name FROM categories WHERE id=$1;
	`
	category_row := c.db.QueryRowContext(ctx, q, id)
	category := types.Category{}
	err := category_row.Scan(&category.Id, &category.Name)
	if err == sql.ErrNoRows {
		return nil, types.NewErrNotFound("category not found")
	}
	return &category, err
}

func (c *categoryRepo) GetByName(ctx context.Context, name string) (*types.Category, error) {
	q := `
		SELECT id, name FROM categories WHERE name=$1;
	`
	category_row := c.db.QueryRowContext(ctx, q, name)
	category := types.Category{}
	err := category_row.Scan(&category.Id, &category.Name)
	if err == sql.ErrNoRows {
		return nil, types.NewErrNotFound("category not found")
	}
	return &category, err
}

func (c *categoryRepo) Update(ctx context.Context, id string, entity types.Category) error {
	q := `
		UPDATE categories SET name=$1 WHERE id=$2;
	`
	_, err := c.db.ExecContext(ctx, q, entity.Name, id)
	return err
}
