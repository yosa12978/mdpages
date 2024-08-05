package handler

import (
	"encoding/json"
	"net/http"

	"github.com/yosa12978/mdpages/services"
	"github.com/yosa12978/mdpages/types"
	"github.com/yosa12978/mdpages/util"
)

type CommitHandler interface {
	GetArticleCommits() Handler
	GetCommitById() Handler
	CreateCommit() Handler
	DeleteCommit() Handler
}

type commitHandler struct {
	commitService services.CommitService
}

func NewCommitHandler(commitService services.CommitService) CommitHandler {
	return &commitHandler{
		commitService: commitService,
	}
}

func (c *commitHandler) CreateCommit() Handler {
	return func(w http.ResponseWriter, r *http.Request) error {
		body := types.CommitCreateDto{}
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			// here render red box
			return err
		}
		err := c.commitService.Create(r.Context(), body)
		// here render green box
		return err
	}
}

func (c *commitHandler) DeleteCommit() Handler {
	return func(w http.ResponseWriter, r *http.Request) error {
		// check here if user root or author then delete
		err := c.commitService.Delete(r.Context(), r.PathValue("id"))
		return err
	}
}

func (c *commitHandler) GetArticleCommits() Handler {
	return func(w http.ResponseWriter, r *http.Request) error {
		commits, err := c.commitService.GetArticleCommits(r.Context(), r.PathValue("article_id"))
		if err != nil {
			return err
		}
		return util.RenderBlock(w, "commits", commits)
	}
}

func (c *commitHandler) GetCommitById() Handler {
	return func(w http.ResponseWriter, r *http.Request) error {
		commit, err := c.commitService.GetById(r.Context(), r.PathValue("commit_id"))
		if err != nil {
			return err
		}
		return util.RenderBlock(w, "commit", commit)
	}
}
