package services

import (
	"context"
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

	Seed(ctx context.Context) error
}

type accountService struct {
	accountRepo repos.AccountRepo
	logger      logging.Logger
}

func NewAccountService(
	accountRepo repos.AccountRepo,
	logger logging.Logger,
) AccountService {
	return &accountService{
		accountRepo: accountRepo,
		logger:      logger,
	}
}

func (a *accountService) Delete(ctx context.Context, username string) error {
	return a.accountRepo.Delete(ctx, username)
}

func (a *accountService) GetByUsername(ctx context.Context, username string) (*types.Account, error) {
	return a.accountRepo.GetByUsername(ctx, username)
}

func (a *accountService) GetByCredentials(ctx context.Context, username, password string) (*types.Account, error) {
	acc, err := a.GetByUsername(ctx, username)
	if err != nil {
		return nil, fmt.Errorf("user doesn't exist")
	}
	if util.CheckPasswordHash(password+acc.Salt, acc.Password) {
		return nil, fmt.Errorf("user doesn't exist")
	}
	return acc, nil
}

// Create implements AccountService.
func (a *accountService) Create(ctx context.Context, dto types.AccountCreateDto) error {
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
		Role:     dto.Role,
	})
}

// Seed implements AccountService.
func (a *accountService) Seed(ctx context.Context) error {
	if err := a.Create(ctx, types.AccountCreateDto{
		Username: "admin",
		Password: "admin1234",
		Role:     "ADMIN",
	}); err != nil {
		return err
	}
	return a.Create(ctx, types.AccountCreateDto{
		Username: "user",
		Password: "user1234",
		Role:     "USER",
	})
}
