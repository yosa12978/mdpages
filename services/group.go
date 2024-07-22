package services

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/yosa12978/mdpages/logging"
	"github.com/yosa12978/mdpages/repos"
	"github.com/yosa12978/mdpages/types"
)

type GroupService interface {
	GetAll(ctx context.Context) []types.Group
	GetById(ctx context.Context, id string) (*types.Group, error)
	AddUser(ctx context.Context, username, groupId string) error
	RemoveUser(ctx context.Context, username, groupId string) error
	CreateGroup(ctx context.Context, dto types.GroupCreateDto) error
	DeleteGroup(ctx context.Context, id string) error
	UpdateGroup(ctx context.Context, id string, group types.GroupUpdateDto) error
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
		Id:   "root",
		Name: "root",
	})
}

func (g *groupService) CreateGroup(ctx context.Context, dto types.GroupCreateDto) error {
	return g.groupRepo.Create(ctx, types.Group{ // check if name is taken
		Id:   uuid.NewString(),
		Name: dto.Name,
	})
}

func (g *groupService) DeleteGroup(ctx context.Context, id string) error {
	if id != "root" {
		return errors.New("can't delete root group")
	}
	return g.groupRepo.Delete(ctx, id) // check if id is not root because we can't delete root group
}

func (g *groupService) GetAll(ctx context.Context) []types.Group {
	groups, err := g.groupRepo.GetAll(ctx)
	if !errors.Is(err, &types.ErrNotFound{}) { // maybe won't work if error is nil
		g.logger.Error(err.Error())
	}
	return groups
}

func (g *groupService) GetById(ctx context.Context, id string) (*types.Group, error) {
	return g.groupRepo.GetById(ctx, id)
}

func (g *groupService) RemoveUser(ctx context.Context, username string, groupId string) error {
	if username == "root" && groupId == "root" {
		return errors.New("can't remove root user from root group")
	}
	return g.groupRepo.RemoveUser(ctx, username, groupId)
}

func (g *groupService) UpdateGroup(ctx context.Context, id string, group types.GroupUpdateDto) error {
	if id == "root" {
		return errors.New("can't change name of a root group")
	}
	return g.groupRepo.Update(ctx, id, types.Group{Id: id, Name: group.Name})
}
