package types

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
}

type GroupCreateDto struct {
	Name string
}
