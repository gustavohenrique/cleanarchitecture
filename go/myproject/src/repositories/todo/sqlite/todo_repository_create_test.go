package sqlite_test

import (
	"context"
	"fmt"
	"testing"

	"{{ .ProjectName }}/src/entities"
	"{{ .ProjectName }}/src/infra"
	"{{ .ProjectName }}/src/infra/sqlite"
	db "{{ .ProjectName }}/src/repositories/todo/sqlite"
	"{{ .ProjectName }}/src/shared/test"
	"{{ .ProjectName }}/src/shared/test/assert"
	"{{ .ProjectName }}/src/valueobjects"
)

func TestTodoItemRepositoryCreate(ts *testing.T) {
	test.WithSqlite(ts, "Should create required fields", func(t *testing.T, store *sqlite.SqliteStore, ctx context.Context) {
		item := entities.NewTodoItemEntity()
		item.Title = "My todo_item"
		repo := db.NewRepository(infra.New())
		_, err := repo.Create(ctx, item)
		assert.Nil(t, err, fmt.Sprintf("%s", err))

		var found valueobjects.TodoItemTable
		query := "select * from todo_items where id=?"
		assert.Nil(t, store.Query(query, &found, item.ID))

		assert.Equal(t, found.ID, item.ID)
		assert.Equal(t, found.Title, item.Title)
		assert.False(t, found.IsDone)
	})
}
