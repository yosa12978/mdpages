package repos

import (
	"context"
	"database/sql"

	"github.com/yosa12978/mdpages/types"
)

type CommitRepo interface {
	GetArticleCommits(ctx context.Context, articleId string) ([]types.Commit, error)
	GetById(ctx context.Context, id string) (*types.Commit, error)
	Create(ctx context.Context, entity types.Commit) error
	Delete(ctx context.Context, id string) error
}

type commitRepo struct {
	db *sql.DB
}

func NewCommitRepo(db *sql.DB) CommitRepo {
	return &commitRepo{
		db: db,
	}
}

// GetArticleCommits implements CommitRepo.
func (c *commitRepo) GetArticleCommits(ctx context.Context, articleId string) ([]types.Commit, error) {
	q := `
		SELECT c.id, c.title,  c.author, c.created
		FROM commits c WHERE c.article_id=$1;
	`
	row, err := c.db.QueryContext(ctx, q, articleId)
	if err != nil {
		if err == sql.ErrNoRows {
			return []types.Commit{}, types.NewErrNotFound("commits not found")
		}
		return nil, err
	}
	commits := []types.Commit{}
	for row.Next() {
		commit := types.Commit{}
		row.Scan(
			&commit.Id,
			&commit.Title,
			&commit.Author,
			&commit.Created,
		)
		commits = append(commits, commit)
	}
	return commits, nil
}

// Create implements CommitRepo.
func (c *commitRepo) Create(ctx context.Context, entity types.Commit) error {
	q := `
		INSERT INTO commits (id, title, body, article_id, author, created)
		VALUES ($1, $2, $3, $4, $5, $6);
	`
	_, err := c.db.ExecContext(ctx, q,
		entity.Id,
		entity.Title,
		entity.Body,
		entity.ArticleId,
		entity.Author,
		entity.Created,
	)
	return err
}

// Delete implements CommitRepo.
func (c *commitRepo) Delete(ctx context.Context, id string) error {
	q := `
		DELETE FROM commits WHERE id=$1;
	`
	_, err := c.db.ExecContext(ctx, q, id)
	return err
}

// GetById implements CommitRepo.
func (c *commitRepo) GetById(ctx context.Context, id string) (*types.Commit, error) {
	q := `
		SELECT id, title, body, article_id, author, created 
		FROM commits c WHERE c.id = $1;
	`
	commit_row := c.db.QueryRowContext(ctx, q, id)
	commit := types.Commit{}
	err := commit_row.Scan(
		&commit.Id,
		&commit.Title,
		&commit.Body,
		&commit.ArticleId,
		&commit.Created,
	)
	if err == sql.ErrNoRows {
		return nil, types.NewErrNotFound("commit not found")
	}
	return &commit, err
}
