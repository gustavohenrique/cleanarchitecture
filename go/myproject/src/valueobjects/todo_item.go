package valueobjects

import (
	"myproject/src/shared/customerror"
	"myproject/src/shared/strings"
)

type TodoItemRequest struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	IsDone bool   `json:"isDone"`
}

func (s TodoItemRequest) Validate() error {
	if strings.HasEmpty(s.Title) {
		return customerror.Invalid("Title is required")
	}
	return nil
}

type TodoItemResponse struct {
	TodoItemRequest
	CreatedAt string `json:"created_at"`
}

type TodoItemTable struct {
	ID        string `db:"id"`
	CreatedAt string `db:"created_at"`
	Title     string `db:"title"`
	IsDone    bool   `db:"is_done"`
}
