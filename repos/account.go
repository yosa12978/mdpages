package repos

import (
	"context"
	"database/sql"
	"errors"

	"github.com/yosa12978/mdpages/types"
)

type AccountRepo interface {
	CRUD[types.Account]
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
			id, 
			username, 
			password, 
			salt, 
			created, 
			role
		) VALUES ($1, $2, $3, $4, $5);
	`
	_, err := a.db.ExecContext(ctx, q,
		account.Id,
		account.Username,
		account.Password,
		account.Salt,
		account.Created,
		account.Role,
	)
	return err
}

func (a *accountRepo) Delete(ctx context.Context, accountId string) error {
	q := `
		DELETE FROM accounts WHERE id=$1;
	`
	_, err := a.db.ExecContext(ctx, q, accountId)
	return err
}

func (a *accountRepo) GetById(ctx context.Context, id string) (*types.Account, error) {
	q := `
		SELECT id, username, password, salt, created, role FROM accounts WHERE id = $1;
	`
	user_row := a.db.QueryRowContext(ctx, q, id)
	user := types.Account{}
	err := user_row.Scan(&user.Id, &user.Username, &user.Password, &user.Salt, &user.Created, &user.Role)
	return &user, err
}

func (a *accountRepo) GetByUsername(ctx context.Context, username string) (*types.Account, error) {
	q := `
		SELECT id, username, password, salt, created, role FROM accounts WHERE username = $1;
	`
	user_row := a.db.QueryRowContext(ctx, q, username)
	user := types.Account{}
	err := user_row.Scan(&user.Id, &user.Username, &user.Password, &user.Salt, &user.Created, &user.Role)
	return &user, err
}

func (a *accountRepo) Update(ctx context.Context, accountId string, account types.Account) error {
	q := `
		UPDATE accounts SET username=$1, password=$2, salt=$3, role=$4 WHERE id=$5;
	`
	_, err := a.db.ExecContext(ctx, q,
		account.Username,
		account.Password,
		account.Salt,
		account.Role,
		accountId,
	)
	return err
}
