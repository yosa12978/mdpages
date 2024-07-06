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
	Id   string
	Name string
}
