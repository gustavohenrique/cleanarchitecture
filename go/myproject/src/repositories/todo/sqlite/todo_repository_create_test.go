package sqlite_test

import (
	"context"
	"fmt"
	"testing"

	"myproject/src/entities"
	"myproject/src/infra"
	"myproject/src/infra/sqlite"
	db "myproject/src/repositories/todo/sqlite"
	"myproject/src/shared/test"
	"myproject/src/shared/test/assert"
	"myproject/src/valueobjects"
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
