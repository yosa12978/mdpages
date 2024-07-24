package repos

import (
	"context"
	"database/sql"

	"github.com/yosa12978/mdpages/types"
)

type GroupRepo interface {
	GetAll(ctx context.Context) ([]types.Group, error)
	GetById(ctx context.Context, id string) (*types.Group, error)
	Create(ctx context.Context, group types.Group) error
	Delete(ctx context.Context, id string) error
	Update(ctx context.Context, id string, group types.Group) error
	AddUser(ctx context.Context, username, groupId string) error
	RemoveUser(ctx context.Context, username, groupId string) error
	GetUserGroups(ctx context.Context, username string) ([]types.Group, error)
}

type groupRepo struct {
	db *sql.DB
}

func NewGroupRepo(db *sql.DB) GroupRepo {
	return &groupRepo{db: db}
}

// SELECT username, password, salt, created, ARRAY(
//     SELECT g.id FROM groups g
//     INNER JOIN accounts_groups ag
//     ON ag.group_id=g.id AND ag.account_id='root'
// ) AS groups FROM accounts WHERE username = 'root';

func (g *groupRepo) GetUserGroups(ctx context.Context, username string) ([]types.Group, error) {
	q := `
		SELECT g.id, g.name FROM groups g 
		INNER JOIN accounts_groups ag 
		ON ag.group_id=g.id AND ag.account_id=$1;
	`
	groups := []types.Group{}
	row, err := g.db.QueryContext(ctx, q, username)
	if err != nil {
		if err == sql.ErrNoRows {
			return groups, types.NewErrNotFound("user not found")
		}
		return groups, err
	}
	for row.Next() {
		group := types.Group{}
		row.Scan(&group.Id, &group.Name)
		groups = append(groups, group)
	}
	return groups, nil
}

func (g *groupRepo) AddUser(ctx context.Context, username, groupId string) error {
	q := `
		INSERT INTO accounts_groups(account_id, group_id) 
		VALUES ($1, $2);
	`
	_, err := g.db.ExecContext(ctx, q, username, groupId)
	return err
}
func (g *groupRepo) RemoveUser(ctx context.Context, username, groupId string) error {
	q := `
		DELETE FROM accounts_groups 
		WHERE account_id=$1 AND group_id=$2;
	`
	_, err := g.db.ExecContext(ctx, q, username, groupId)
	return err
}

func (g *groupRepo) GetAll(ctx context.Context) ([]types.Group, error) {
	q := `
		SELECT id, name FROM groups;
	`
	rows, err := g.db.QueryContext(ctx, q)
	if err != nil {
		if err == sql.ErrNoRows {
			return []types.Group{}, types.NewErrNotFound("groups not found")
		}
		return nil, err
	}
	groups := []types.Group{}
	for rows.Next() {
		group := types.Group{}
		rows.Scan(&group.Id, &group.Name)
		groups = append(groups, group)
	}
	return groups, nil
}

func (g *groupRepo) GetById(ctx context.Context, id string) (*types.Group, error) {
	q := `
		SELECT id, name FROM groups WHERE id=$1;
	`
	row := g.db.QueryRowContext(ctx, q, id)
	group := types.Group{}
	err := row.Scan(&group.Id, &group.Name)
	if err == sql.ErrNoRows {
		return nil, types.NewErrNotFound("group not found")
	}
	return &group, err
}

func (g *groupRepo) Create(ctx context.Context, group types.Group) error {
	q := `
		INSERT INTO groups(id, name) VALUES ($1, $2);
	`
	_, err := g.db.ExecContext(ctx, q, group.Id, group.Name)
	return err
}

func (g *groupRepo) Delete(ctx context.Context, id string) error {
	q := `
		DELETE FROM groups WHERE id=$1;
	`
	_, err := g.db.ExecContext(ctx, q, id)
	return err
}

func (g *groupRepo) Update(ctx context.Context, id string, group types.Group) error {
	q := `
		UPDATE groups SET name=$1 WHERE id=$2
	`
	_, err := g.db.ExecContext(ctx, q, group.Name, id)
	return err
}
