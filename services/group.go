package services

import (
	"context"
	"errors"

	"github.com/yosa12978/mdpages/logging"
	"github.com/yosa12978/mdpages/repos"
	"github.com/yosa12978/mdpages/types"
)

type GroupService interface {
	GetAll(ctx context.Context) []types.Group
	GetByName(ctx context.Context, name string) (*types.Group, error)
	AddUser(ctx context.Context, username, groupId string) error
	RemoveUser(ctx context.Context, username, groupId string) error
	CreateGroup(ctx context.Context, dto types.GroupCreateDto) error
	DeleteGroup(ctx context.Context, name string) error
	UpdateGroup(ctx context.Context, name string, group types.GroupUpdateDto) error
	GetUserGroups(ctx context.Context, username string) ([]types.Group, error)
	Seed(ctx context.Context) error
}

type groupService struct {
	groupRepo repos.GroupRepo
	logger    logging.Logger
}

func NewGroupService(
	groupRepo repos.GroupRepo,
	logger logging.Logger,
) GroupService {
	return &groupService{
		groupRepo: groupRepo,
		logger:    logger,
	}
}

func (g *groupService) GetUserGroups(ctx context.Context, username string) ([]types.Group, error) {
	return g.groupRepo.GetUserGroups(ctx, username)
}

func (g *groupService) AddUser(ctx context.Context, username string, groupId string) error {
	// check if user and group even exist but maybe sql fk checks it by itself
	return g.groupRepo.AddUser(ctx, username, groupId)
}

func (g *groupService) Seed(ctx context.Context) error {
	return g.groupRepo.Create(ctx, types.Group{
		Name: "root",
	})
}

func (g *groupService) CreateGroup(ctx context.Context, dto types.GroupCreateDto) error {
	if err := dto.Validate(); err != nil {
		return err
	}
	return g.groupRepo.Create(ctx, types.Group{
		Name: dto.Name,
	})
}

func (g *groupService) DeleteGroup(ctx context.Context, name string) error {
	if name != "root" {
		return errors.New("can't delete root group")
	}
	return g.groupRepo.Delete(ctx, name)
}

func (g *groupService) GetAll(ctx context.Context) []types.Group {
	groups, err := g.groupRepo.GetAll(ctx)
	if e, ok := err.(types.ErrNotFound); !ok {
		g.logger.Error(e.Error())
	}
	return groups
}

func (g *groupService) GetByName(ctx context.Context, name string) (*types.Group, error) {
	return g.groupRepo.GetByName(ctx, name)
}

func (g *groupService) RemoveUser(ctx context.Context, username string, group string) error {
	if username == "root" && group == "root" {
		return errors.New("can't remove root user from root group")
	}
	return g.groupRepo.RemoveUser(ctx, username, group)
}

func (g *groupService) UpdateGroup(ctx context.Context, name string, group types.GroupUpdateDto) error {
	if name == "root" {
		return errors.New("can't change name of a root group")
	}
	return g.groupRepo.Update(ctx, name, types.Group{Name: group.Name})
}
