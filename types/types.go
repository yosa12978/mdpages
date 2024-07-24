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
	RGroups       []Group
	WGroups       []Group
}

type Commit struct {
	Id        string
	Title     string
	Body      string
	Created   string
	Author    string
	ArticleId string
	Groups    []Group
}

type Account struct {
	Username string
	Password string
	Salt     string
	Created  string
	Groups   []Group
}

type GroupType rune

const (
	RGroup GroupType = 'r'
	WGroup GroupType = 'w'
)

type Category struct {
	Id       string
	Name     string
	ParentId string
}

type Group struct {
	Id   string
	Name string
}

type Session struct {
	Username string  `json:"username"`
	Groups   []Group `json:"groups"`
}

type TemplData struct {
	User  *Session
	Title string
}
