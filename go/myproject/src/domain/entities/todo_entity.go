package entities

import (
	"fmt"
	"strings"

	"{{ .ProjectName }}/src/shared/datetime"
	"{{ .ProjectName }}/src/shared/uuid"
)

type Todo struct {
	Base
	Title  string
	IsDone bool
}

func NewTodo() Todo {
	me := Todo{}
	me.ID = uuid.NewV4()
	me.CreatedAt = datetime.Now()
	return me
}

func (s Todo) ValidateInsertRequest() error {
	if len(strings.TrimSpace(s.Title)) < 1 {
		return fmt.Errorf("title should have at least 1 characters")
	}
	return nil
}
