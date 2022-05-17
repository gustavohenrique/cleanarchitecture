package sqlite_test

import (
	"context"
	"testing"

	"myproject/src/entities"
	"myproject/src/infra"
	"myproject/src/infra/sqlite"
	db "myproject/src/repositories/todo/sqlite"
	"myproject/src/shared/test"
	"myproject/src/shared/test/assert"
	"myproject/src/shared/uuid"
	"myproject/src/valueobjects"
)

func TestTodoItemRepositoryUpdate(ts *testing.T) {
	item := valueobjects.TodoItemTable{}
	item.ID = uuid.NewV4()
	item.Title = "Gustavo"
	item.IsDone = true
	insertTodoItemQuery := "insert into todo_items (id, title, is_done) values (?, ?, ?)"

	test.WithSqlite(ts, "Should update fields", func(t *testing.T, store *sqlite.SqliteStore, ctx context.Context) {
		assert.Nil(t, store.Exec(insertTodoItemQuery, item.ID, item.Title, item.IsDone))

		repo := db.NewRepository(infra.New())
		changed := entities.TodoItemEntity{}
		changed.ID = item.ID
		changed.Title = "Update the todo list"
		changed.IsDone = true
		saved, err := repo.Update(ctx, changed)
		assert.Nil(t, err)

		var found valueobjects.TodoItemTable
		query := "select id, title, is_done from todo_items where id=?"
		assert.Nil(t, store.Query(query, &found, saved.ID))
		assert.Equal(t, found.ID, saved.ID)
		assert.Equal(t, found.Title, saved.Title)
		assert.Equal(t, found.IsDone, saved.IsDone)
	})

	test.WithSqlite(ts, "Should fail when title is empty", func(t *testing.T, store *sqlite.SqliteStore, ctx context.Context) {
		assert.Nil(t, store.Exec(insertTodoItemQuery, item.ID, item.Title, item.IsDone))

		repo := db.NewRepository(infra.New())
		changed := entities.TodoItemEntity{}
		changed.ID = item.ID
		changed.IsDone = false
		changed.Title = ""
		_, err := repo.Update(ctx, changed)
		assert.NotNil(t, err)

		var found valueobjects.TodoItemTable
		query := "select id, title, is_done from todo_items where id=?"
		assert.Nil(t, store.Query(query, &found, item.ID))
		assert.Equal(t, found.ID, item.ID)
		assert.Equal(t, found.Title, item.Title)
		assert.Equal(t, found.IsDone, item.IsDone)
	})
}
