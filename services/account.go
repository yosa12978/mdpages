package services

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/yosa12978/mdpages/logging"
	"github.com/yosa12978/mdpages/repos"
	"github.com/yosa12978/mdpages/types"
	"github.com/yosa12978/mdpages/util"
)

type AccountService interface {
	GetByUsername(ctx context.Context, username string) (*types.Account, error)
	GetByCredentials(ctx context.Context, username, password string) (*types.Account, error)

	Create(ctx context.Context, dto types.AccountCreateDto) error
	Delete(ctx context.Context, username string) error

	Seed(ctx context.Context, rootPassword string) error
}

type accountService struct {
	accountRepo  repos.AccountRepo
	groupService GroupService
	logger       logging.Logger
}

func NewAccountService(
	accountRepo repos.AccountRepo,
	groupService GroupService,
	logger logging.Logger,
) AccountService {
	return &accountService{
		accountRepo:  accountRepo,
		groupService: groupService,
		logger:       logger,
	}
}

func (a *accountService) Delete(ctx context.Context, username string) error {
	if username == "root" {
		return errors.New("can't delete root user")
	}
	return a.accountRepo.Delete(ctx, username)
}

func (a *accountService) GetByUsername(ctx context.Context, username string) (*types.Account, error) {
	acc, err := a.accountRepo.GetByUsername(ctx, username)
	if err != nil {
		return nil, err
	}
	groups, err := a.groupService.GetUserGroups(ctx, username)
	if err != nil {
		return nil, err
	}
	acc.Groups = groups
	return acc, nil
}

func (a *accountService) GetByCredentials(ctx context.Context, username, password string) (*types.Account, error) {
	acc, err := a.GetByUsername(ctx, username)
	if err != nil {
		return nil, fmt.Errorf("username or password is wrong")
	}
	if !util.CheckPasswordHash(password+acc.Salt, acc.Password) { // this slows everything
		return nil, fmt.Errorf("username or password is wrong")
	}
	return acc, nil
}

// Create implements AccountService.
func (a *accountService) Create(ctx context.Context, dto types.AccountCreateDto) error {
	if err := dto.Validate(); err != nil {
		return err
	}

	if _, err := a.accountRepo.GetByUsername(ctx, dto.Username); err == nil {
		return errors.New("username is already taken")
	}
	salt := uuid.NewString()
	hashedPassword, err := util.HashPassword(dto.Password + salt)
	if err != nil {
		return err
	}
	return a.accountRepo.Create(ctx, types.Account{
		Username: dto.Username,
		Password: hashedPassword,
		Salt:     salt,
		Created:  time.Now().Format(time.RFC3339),
	})
}

// Seed implements AccountService.
func (a *accountService) Seed(ctx context.Context, rootPassword string) error {
	if err := a.Create(ctx, types.AccountCreateDto{
		Username: "root",
		Password: rootPassword,
	}); err != nil {
		a.logger.Error(err.Error())
	}
	return a.groupService.AddUser(ctx, "root", "root")
}
