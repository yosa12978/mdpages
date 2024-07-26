package types

import (
	"errors"
	"strings"
)

// update dtos  *** id sent as url path value

type ArticleUpdateDto struct {
	CategoryId string
}

type CategoryUpdateDto struct {
	Name     string
	ParentId string
}

type AccountUpdateDto struct {
	OldPassword string
	NewPassword string
}

func (m *AccountUpdateDto) Validate() error {
	if strings.Contains(m.NewPassword, " ") {
		return errors.New("password can't contain spaces")
	}
	if len(m.NewPassword) < 4 {
		return errors.New("length of your password can't be less then 4 characters")
	}
	return nil
}

type GroupUpdateDto struct {
	Name string
}
