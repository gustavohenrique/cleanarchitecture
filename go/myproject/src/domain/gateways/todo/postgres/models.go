package postgres

import (
	"time"

	"{{ .ProjectName }}/src/domain/entities"
)

type TodoRow struct {
	ID        string `db:"id"`
	CreatedAt string `db:"created_at"`
	Title     string `db:"title"`
	IsDone    bool   `db:"is_done"`
}

func (in *TodoRow) ToEntity() entities.Todo {
	out := entities.Todo{}
	out.ID = in.ID
	out.Title = in.Title
	out.IsDone = in.IsDone
	dt, _ := time.Parse(time.RFC3339, in.CreatedAt)
	out.CreatedAt = &dt
	return out
}

func ToEntities(in []TodoRow) []entities.Todo {
	var out []entities.Todo
	for _, item := range in {
		out = append(out, item.ToEntity())
	}
	return out
}
