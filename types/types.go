package types

type Article struct {
	Id            string
	CategoryId    string
	CategoryName  string
	Title         string
	Body          string
	CommitCreated string
	CommitId      string
}

type Commit struct {
	Id             string
	Title          string
	Body           string
	Created        string
	AuthorId       string
	AuthorUsername string
	ArticleId      string
}

type Account struct {
	Id       string
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
