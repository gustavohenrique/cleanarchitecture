package postgres_test

import (
	"context"
	"testing"

	"{{ .ProjectName }}/src/domain/entities"
	db "{{ .ProjectName }}/src/domain/gateways/todo/postgres"
	"{{ .ProjectName }}/src/infra/conf"
	"{{ .ProjectName }}/src/infra/datastores"
	"{{ .ProjectName }}/src/infra/test"
	"{{ .ProjectName }}/src/infra/test/assert"
	"{{ .ProjectName }}/src/interfaces"
	"{{ .ProjectName }}/src/shared/uuid"
)

func TestTodoItemRepositoryUpdate(ts *testing.T) {
	datastore := datastores.With(conf.Get()).New()

	item := db.TodoRow{}
	item.ID = uuid.NewV4()
	item.Title = "Gustavo"
	item.IsDone = true
	insertTodoItemQuery := "insert into todo_items (id, title, is_done) values ($1, $2, $3)"

	test.WithPostgres(ts, "Should update fields", func(t *testing.T, store interfaces.ISqlDataStore, ctx context.Context) {
		assert.Nil(t, store.Exec(insertTodoItemQuery, item.ID, item.Title, item.IsDone))

		changed := entities.Todo{}
		changed.ID = item.ID
		changed.Title = "Update the todo list"
		changed.IsDone = true
		gateway := db.New(datastore)
		saved, err := gateway.Update(ctx, changed)
		assert.Nil(t, err)

		var found db.TodoRow
		query := "select id, title, is_done from todo_items where id=$1"
		assert.Nil(t, store.Query(query, &found, saved.ID))
		assert.Equal(t, found.ID, saved.ID)
		assert.Equal(t, found.Title, saved.Title)
		assert.Equal(t, found.IsDone, saved.IsDone)
	})

	test.WithPostgres(ts, "Should fail when title is empty", func(t *testing.T, store interfaces.ISqlDataStore, ctx context.Context) {
		assert.Nil(t, store.Exec(insertTodoItemQuery, item.ID, item.Title, item.IsDone))

		changed := entities.Todo{}
		changed.ID = item.ID
		changed.IsDone = false
		changed.Title = ""
		gateway := db.New(datastore)
		_, err := gateway.Update(ctx, changed)
		assert.NotNil(t, err)

		var found db.TodoRow
		query := "select id, title, is_done from todo_items where id=$1"
		assert.Nil(t, store.Query(query, &found, item.ID))
		assert.Equal(t, found.ID, item.ID)
		assert.Equal(t, found.Title, item.Title)
		assert.Equal(t, found.IsDone, item.IsDone)
	})
}
