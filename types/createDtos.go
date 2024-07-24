package types

import (
	"errors"
	"strings"
)

type ArticleCreateDto struct {
	Title      string `json:"title"`
	Body       string `json:"body"`
	Author     string `json:"author"`
	CategoryId string `json:"categoryId"`
}

type CommitCreateDto struct {
	Title     string `json:"title"`
	Body      string `json:"body"`
	ArticleId string `json:"articleId"`
	Author    string `json:"author"`
}

type CategoryCreateDto struct {
	Name     string `json:"name"`
	ParentId string `json:"parentId"`
}

func (m *CategoryCreateDto) Validate() error {
	m.Name = strings.TrimSpace(m.Name)
	if len(m.Name) < 3 {
		return errors.New("length of username less then 3 characters")
	}
	return nil
}

type AccountCreateDto struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (m *AccountCreateDto) Validate() error {
	m.Username = strings.TrimSpace(m.Username)
	if len(m.Username) < 4 {
		return errors.New("length of username less then 4 characters")
	}
	if strings.Contains(m.Username, " ") {
		return errors.New("username can't contain spaces")
	}
	if strings.Contains(m.Password, " ") {
		return errors.New("password can't contain spaces")
	}
	if len(m.Password) < 4 {
		return errors.New("length of your password can't be less then 4 characters")
	}
	return nil
}

type GroupCreateDto struct {
	Name string `json:"name"`
}

func (m *GroupCreateDto) Validate() error {
	m.Name = strings.TrimSpace(m.Name)
	if len(m.Name) < 3 {
		return errors.New("length of group name less then 3 characters")
	}
	return nil
}

type LoginDto struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
