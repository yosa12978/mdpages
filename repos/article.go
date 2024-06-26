package repos

import (
	"context"
	"database/sql"

	"github.com/yosa12978/mdpages/types"
)

type ArticleRepo interface {
	CRUD[types.Article]
}

type articleRepo struct {
	db *sql.DB
}

func NewArticleRepo(db *sql.DB) ArticleRepo {
	return &articleRepo{
		db: db,
	}
}

func (a *articleRepo) Create(ctx context.Context, entity types.Article) error {
	panic("unimplemented")
}

func (a *articleRepo) Delete(ctx context.Context, id string) error {
	panic("unimplemented")
}

func (a *articleRepo) GetAll(ctx context.Context) ([]types.Article, error) {
	q := `
		SELECT
			a.id AS article_id, 
			categ.id AS category_id, 
			categ.name AS category_name,
			comm.id AS commit_id, 
			comm.title AS title, 
			comm.created AS last_updated
		FROM articles a 
		INNER JOIN categories categ ON categ.id = a.category_id
		INNER JOIN commits comm ON comm.id = (
			SELECT id FROM commits WHERE article_id = a.id ORDER BY created DESC LIMIT 1
		)  
		ORDER BY title;
	`
	rows, err := a.db.QueryContext(ctx, q)
	if err != nil {
		return nil, err
	}
	articles := []types.Article{}
	for rows.Next() {
		article := types.Article{}
		rows.Scan(
			&article.Id,
			&article.CategoryId,
			&article.CategoryName,
			&article.CommitId,
			&article.Title,
			&article.CommitCreated,
		)
		articles = append(articles, article)
	}
	return articles, nil
}

func (a *articleRepo) GetById(ctx context.Context, id string) (*types.Article, error) {
	q := `
		SELECT
			a.id AS article_id, 
			categ.id AS category_id, 
			categ.name AS category_name,
			comm.id AS commit_id, 
			comm.title AS title, 
			comm.created AS last_updated,
			comm.body AS body
		FROM articles a 
		INNER JOIN categories categ ON categ.id = a.category_id
		INNER JOIN commits comm ON comm.article_id = $1 
		ORDER BY last_updated DESC LIMIT 1;
	`
	article_row := a.db.QueryRowContext(ctx, q, id)
	article := types.Article{}
	err := article_row.Scan(
		&article.Id,
		&article.CategoryId,
		&article.CategoryName,
		&article.CommitId,
		&article.Title,
		&article.CommitCreated,
		&article.Body,
	)
	return &article, err
}

func (a *articleRepo) Update(ctx context.Context, id string, entity types.Article) error {
	q := `
		UPDATE articles SET category=$1 WHERE id=$2
	`
	_, err := a.db.ExecContext(ctx, q, entity.CategoryId, id)
	return err
}
