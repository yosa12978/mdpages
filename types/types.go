package types

type Article struct {
	Id            string
	CategoryId    string
	CategoryName  string
	Title         string
	Body          string
	CommitCreated string
	CommitId      string
	CommitAuthor  string
}

type Commit struct {
	Id        string
	Title     string
	Body      string
	Created   string
	Author    string
	ArticleId string
}

type Account struct {
	Username string
	Password string
	Salt     string
	Created  string
	Role     string
}

type Category struct {
	Id       string
	Name     string
	ParentId string
}

// Create Dtos

type ArticleCreateDto struct {
	Title      string
	Body       string
	Author     string
	CategoryId string
}

type CommitCreateDto struct {
	Title     string
	Body      string
	ArticleId string
	Author    string
}

type CategoryCreateDto struct {
	Name     string
	ParentId string
}

type AccountCreateDto struct {
	Username string
	Password string
	Role     string
}
