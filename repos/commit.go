package repos

import "github.com/yosa12978/mdpages/types"

type CommitRepo interface {
	CRUD[types.Commit]
}
