package repos

import (
	"context"
	"database/sql"

	"github.com/yosa12978/mdpages/types"
	"github.com/yosa12978/mdpages/util"
)

type ArticleRepo interface {
	GetByCategoryId(ctx context.Context, categoryId string) ([]types.Article, error)
	GetById(ctx context.Context, id string) (*types.Article, error)
	GetAll(ctx context.Context) ([]types.Article, error)
	Create(ctx context.Context, entity types.Article) error
	Update(ctx context.Context, id string, entity types.Article) error
	Delete(ctx context.Context, id string) error

	AddRGroup(ctx context.Context, article_id, group_id string) error
	AddWGroup(ctx context.Context, article_id, group_id string) error
	RemoveRGroup(ctx context.Context, article_id, group_id string) error
	RemoveWGroup(ctx context.Context, article_id, group_id string) error
}

type articleRepo struct {
	db *sql.DB
}

func NewArticleRepo(db *sql.DB) ArticleRepo {
	return &articleRepo{
		db: db,
	}
}

func (g *articleRepo) AddRGroup(ctx context.Context, article_id, group_id string) error {
	q := `
		INSERT INTO r_articles_groups(article_id, group_id)
		VALUES ($1, $2)
	`
	_, err := g.db.ExecContext(ctx, q, article_id, group_id)
	return err
}
func (g *articleRepo) AddWGroup(ctx context.Context, article_id, group_id string) error {
	q := `
		INSERT INTO w_articles_groups(article_id, group_id)
		VALUES ($1, $2)
	`
	_, err := g.db.ExecContext(ctx, q, article_id, group_id)
	return err
}

func (g *articleRepo) RemoveRGroup(ctx context.Context, article_id, group_id string) error {
	q := `
		DELETE FROM r_articles_groups WHERE article_id=$1 AND group_id=$2;
	`
	_, err := g.db.ExecContext(ctx, q, article_id, group_id)
	return err
}

func (g *articleRepo) RemoveWGroup(ctx context.Context, article_id, group_id string) error {
	q := `
		DELETE FROM w_articles_groups WHERE article_id=$1 AND group_id=$2;
	`
	_, err := g.db.ExecContext(ctx, q, article_id, group_id)
	return err
}

func (a *articleRepo) GetByCategoryId(ctx context.Context, categoryId string) ([]types.Article, error) {
	q := `
		SELECT
			a.id AS article_id, 
			comm.id AS commit_id, 
			comm.title AS title, 
			comm.created AS last_updated
		FROM articles a WHERE a.category_id = $1
		INNER JOIN commits comm ON comm.id = (
			SELECT id FROM commits WHERE article_id = a.id ORDER BY created DESC LIMIT 1
		)  
		ORDER BY title;
	`
	row, err := a.db.QueryContext(ctx, q, categoryId)
	if err != nil {
		if err == sql.ErrNoRows {
			return []types.Article{}, types.NewErrNotFound("articles not found")
		}
		return nil, err
	}
	articles := []types.Article{}
	for row.Next() {
		article := types.Article{}
		row.Scan(
			&article.Id,
			&article.CommitId,
			&article.Title,
			&article.CommitCreated,
		)
		articles = append(articles, article)
	}
	return articles, nil
}

func (a *articleRepo) Create(ctx context.Context, entity types.Article) error {
	tx, err := a.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()
	q1 := `
		INSERT INTO articles (id, category_id) VALUES ($1, $2);
	`
	q2 := `
		INSERT INTO commits (
			id, 
			title, 
			body, 
			article_id, 
			author, 
			created
		) VALUES ($1, $2, $3, $4, $5, $6);
	`
	_, err = tx.ExecContext(ctx, q1,
		entity.Id,
		util.NewNullString(entity.CategoryId),
	)
	if err != nil {
		return err
	}
	_, err = tx.ExecContext(ctx, q2,
		entity.CommitId,
		entity.Title,
		entity.Body,
		entity.Id,
		entity.CommitAuthor,
		entity.CommitCreated,
	)
	if err != nil {
		return err
	}
	return tx.Commit()
}

func (a *articleRepo) Delete(ctx context.Context, id string) error {
	q := `
		DELETE FROM articles WHERE id=$1;
	`
	_, err := a.db.ExecContext(ctx, q, id)
	return err
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
		if err == sql.ErrNoRows {
			return []types.Article{}, types.NewErrNotFound("articles not found")
		}
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
		LEFT JOIN categories categ ON categ.id = a.category_id
		INNER JOIN commits comm ON comm.article_id = $1 
		ORDER BY last_updated DESC LIMIT 1;
	`
	article_row := a.db.QueryRowContext(ctx, q, id)
	article := types.Article{}
	var categoryId sql.NullString
	var categoryName sql.NullString
	err := article_row.Scan(
		&article.Id,
		&categoryId,
		&categoryName,
		&article.CommitId,
		&article.Title,
		&article.CommitCreated,
		&article.Body,
	)
	article.CategoryId = categoryId.String // instead of this make sql entities and then fetch into them
	article.CategoryName = categoryName.String
	if err == sql.ErrNoRows {
		return nil, types.NewErrNotFound("article not found")
	}
	return &article, err
}

func (a *articleRepo) Update(ctx context.Context, id string, entity types.Article) error {
	q := `
		UPDATE articles SET category=$1 WHERE id=$2
	`
	_, err := a.db.ExecContext(ctx, q, entity.CategoryId, id)
	return err
}
