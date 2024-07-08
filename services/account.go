package services

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/yosa12978/mdpages/repos"
	"github.com/yosa12978/mdpages/types"
	"github.com/yosa12978/mdpages/util"
)

type AccountService interface {
	Create(ctx context.Context, dto types.AccountCreateDto) error
	Seed(ctx context.Context) error
}

type accountService struct {
	accountRepo repos.AccountRepo
}

func NewAccountService(accountRepo repos.AccountRepo) AccountService {
	return &accountService{
		accountRepo: accountRepo,
	}
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
		Role:     "USER",
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
