package types

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

type GroupUpdateDto struct {
	Name string
}
