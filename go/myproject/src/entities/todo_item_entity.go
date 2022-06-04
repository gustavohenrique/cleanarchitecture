package entities

import (
	"fmt"
	"strings"

	"{{ .ProjectName }}/src/shared/datetime"
	"{{ .ProjectName }}/src/shared/uuid"
)

type TodoItemEntity struct {
	Base
	Title  string `json:"title"`
	IsDone bool   `json:"isDone"`
}

func NewTodoItemEntity() TodoItemEntity {
	me := TodoItemEntity{}
	me.ID = uuid.NewV4()
	me.CreatedAt = datetime.Now()
	return me
}

func (s TodoItemEntity) ValidateInsertRequest() error {
	if len(strings.TrimSpace(s.Title)) < 1 {
		return fmt.Errorf("Title should have at least 1 characters")
	}
	return nil
}
