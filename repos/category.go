package repos

import (
	"context"
	"database/sql"

	"github.com/yosa12978/mdpages/types"
)

type CategoryRepo interface {
	GetAll(ctx context.Context) ([]types.Category, error)
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

func (c *categoryRepo) GetAll(ctx context.Context) ([]types.Category, error) {
	q := `
		SELECT id, name FROM categories;
	`
	row, err := c.db.QueryContext(ctx, q)
	if err != nil {
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
		INSERT INTO categories (id, name) VALUES ($1, $2);
	`
	_, err := c.db.ExecContext(ctx, q, entity.Id, entity.Name)
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
	return &category, err
}

func (c *categoryRepo) GetByName(ctx context.Context, name string) (*types.Category, error) {
	q := `
		SELECT id, name FROM categories WHERE name=$1;
	`
	category_row := c.db.QueryRowContext(ctx, q, name)
	category := types.Category{}
	err := category_row.Scan(&category.Id, &category.Name)
	return &category, err
}

func (c *categoryRepo) Update(ctx context.Context, id string, entity types.Category) error {
	q := `
		UPDATE categories SET name=$1 WHERE id=$2;
	`
	_, err := c.db.ExecContext(ctx, q, entity.Name, id)
	return err
}
