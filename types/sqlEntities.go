package types

import "database/sql"

type ArticleEntity struct {
	Id            string
	CategoryId    sql.NullString
	CategoryName  sql.NullString
	Title         string
	Body          string
	CommitCreated string
	CommitId      string
	CommitAuthor  string
}

type CommitEntity struct {
	Id        string
	Title     string
	Body      string
	Created   string
	Author    string
	ArticleId string
}

type AccountEntity struct {
	Username string
	Password string
	Salt     string
	Created  string
}

type CategoryEntity struct {
	Id       string
	Name     string
	ParentId sql.NullString
}

type GroupEntity struct {
	Id   string
	Name string
}
