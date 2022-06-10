package postgres_test

import (
	"context"
	"testing"

	"{{ .ProjectName }}/src/entities"
	"{{ .ProjectName }}/src/infra"
	"{{ .ProjectName }}/src/infra/postgres"
	db "{{ .ProjectName }}/src/repositories/todo/postgres"
	"{{ .ProjectName }}/src/shared/test"
	"{{ .ProjectName }}/src/shared/test/assert"
	"{{ .ProjectName }}/src/shared/uuid"
	"{{ .ProjectName }}/src/valueobjects"
)

func TestTodoItemRepositoryUpdate(ts *testing.T) {
	item := valueobjects.TodoItemTable{}
	item.ID = uuid.NewV4()
	item.Title = "Gustavo"
	item.IsDone = true
	insertTodoItemQuery := "insert into todo_items (id, title, is_done) values ($1, $2, $3)"

	test.WithPostgres(ts, "Should update fields", func(t *testing.T, store *postgres.PostgresStore, ctx context.Context) {
		assert.Nil(t, store.Exec(insertTodoItemQuery, item.ID, item.Title, item.IsDone))

		repo := db.NewRepository(infra.New())
		changed := entities.TodoItemEntity{}
		changed.ID = item.ID
		changed.Title = "Update the todo list"
		changed.IsDone = true
		saved, err := repo.Update(ctx, changed)
		assert.Nil(t, err)

		var found valueobjects.TodoItemTable
		query := "select id, title, is_done from todo_items where id=$1"
		assert.Nil(t, store.Query(query, &found, saved.ID))
		assert.Equal(t, found.ID, saved.ID)
		assert.Equal(t, found.Title, saved.Title)
		assert.Equal(t, found.IsDone, saved.IsDone)
	})

	test.WithPostgres(ts, "Should fail when title is empty", func(t *testing.T, store *postgres.PostgresStore, ctx context.Context) {
		assert.Nil(t, store.Exec(insertTodoItemQuery, item.ID, item.Title, item.IsDone))

		repo := db.NewRepository(infra.New())
		changed := entities.TodoItemEntity{}
		changed.ID = item.ID
		changed.IsDone = false
		changed.Title = ""
		_, err := repo.Update(ctx, changed)
		assert.NotNil(t, err)

		var found valueobjects.TodoItemTable
		query := "select id, title, is_done from todo_items where id=$1"
		assert.Nil(t, store.Query(query, &found, item.ID))
		assert.Equal(t, found.ID, item.ID)
		assert.Equal(t, found.Title, item.Title)
		assert.Equal(t, found.IsDone, item.IsDone)
	})
}
