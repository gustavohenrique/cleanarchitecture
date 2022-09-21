package sqlite_test

import (
	"context"
	"testing"

	"{{ .ProjectName }}/src/domain/entities"
	db "{{ .ProjectName }}/src/domain/gateways/todo/sqlite"
	"{{ .ProjectName }}/src/infra/conf"
	"{{ .ProjectName }}/src/infra/datastores"
	"{{ .ProjectName }}/src/infra/test"
	"{{ .ProjectName }}/src/infra/test/assert"
	"{{ .ProjectName }}/src/interfaces"
	"{{ .ProjectName }}/src/shared/uuid"
)

func TestTodoItemGatewayReadAll(ts *testing.T) {
	datastore := datastores.With(conf.Get()).New()

	todoItem := entities.NewTodo()
	todoItem.Title = "My todoitem"
	insertTodoItemQuery := "insert into todo_items (id, title) values (?, ?)"

	test.WithSqlite(ts, "Should return all", func(t *testing.T, store interfaces.ISqlDataStore, ctx context.Context) {
		assert.Nil(t, store.Exec(insertTodoItemQuery, todoItem.ID, todoItem.Title))

		item := db.TodoRow{}
		item.ID = uuid.NewV4()
		item.Title = "TODO 1"
		item.IsDone = true
		query := "insert into todo_items (id, title, is_done) values (?, ?, ?)"
		assert.Nil(t, store.Exec(
			query,
			item.ID, item.Title, item.IsDone,
		), "Cannot insert todo item")

		gateway := db.New(datastore)
		founds, err := gateway.ReadAll(ctx)
		assert.Nil(t, err)
		assert.Equal(t, len(founds), 2)
		assert.Equal(t, founds[1].ID, item.ID)
		assert.Equal(t, founds[1].Title, item.Title)
		assert.Equal(t, founds[1].IsDone, item.IsDone)
	})
}
