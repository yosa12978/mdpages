package repos

import (
	"context"
	"database/sql"
	"errors"

	"github.com/yosa12978/mdpages/types"
)

type AccountRepo interface {
	GetAll(ctx context.Context) ([]types.Account, error)
	Create(ctx context.Context, account types.Account) error
	Delete(ctx context.Context, username string) error
	Update(ctx context.Context, username string, account types.Account) error
	GetByUsername(ctx context.Context, username string) (*types.Account, error)
}

type accountRepo struct {
	db *sql.DB
}

func NewAccountRepo(db *sql.DB) AccountRepo {
	return &accountRepo{
		db: db,
	}
}

func (a *accountRepo) GetAll(ctx context.Context) ([]types.Account, error) {
	return nil, errors.New("not implemented")
}

func (a *accountRepo) Create(ctx context.Context, account types.Account) error {
	q := `
		INSERT INTO accounts (
			username, 
			password, 
			salt, 
			created 
		) VALUES ($1, $2, $3, $4);
	`
	_, err := a.db.ExecContext(ctx, q,
		account.Username,
		account.Password,
		account.Salt,
		account.Created,
	)
	return err
}

func (a *accountRepo) Delete(ctx context.Context, username string) error {
	q := `
		DELETE FROM accounts WHERE username=$1;
	`
	_, err := a.db.ExecContext(ctx, q, username)
	return err
}

func (a *accountRepo) GetByUsername(ctx context.Context, username string) (*types.Account, error) {
	q := `
		SELECT username, password, salt, created FROM accounts WHERE username = $1;
	`
	user_row := a.db.QueryRowContext(ctx, q, username)
	user := types.Account{}
	err := user_row.Scan(&user.Username, &user.Password, &user.Salt, &user.Created)
	if err == sql.ErrNoRows {
		return nil, types.NewErrNotFound("user not found")
	}
	return &user, err
}

func (a *accountRepo) Update(ctx context.Context, username string, account types.Account) error {
	q := `
		UPDATE accounts SET password=$1, salt=$2 WHERE username=$3;
	`
	_, err := a.db.ExecContext(ctx, q,
		account.Password,
		account.Salt,
		username,
	)
	return err
}
