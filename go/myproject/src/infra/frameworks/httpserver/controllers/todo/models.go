package todo

import (
	"{{ .ProjectName }}/src/domain/entities"
	"{{ .ProjectName }}/src/shared/customerror"
	"{{ .ProjectName }}/src/shared/strings"
)

type TodoJsonRequest struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	IsDone bool   `json:"isDone"`
}

func (s TodoJsonRequest) Validate() error {
	if strings.HasEmpty(s.Title) {
		return customerror.Invalid("Title is required")
	}
	return nil
}

func (in TodoJsonRequest) ToEntity() entities.Todo {
	out := entities.Todo{}
	out.ID = in.ID
	out.Title = in.Title
	out.IsDone = in.IsDone
	return out
}

type TodoJsonResponse struct {
	TodoJsonRequest
	CreatedAt string `json:"created_at"`
}
