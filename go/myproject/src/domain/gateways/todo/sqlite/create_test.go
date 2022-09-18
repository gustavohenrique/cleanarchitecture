package sqlite_test

import (
	"context"
	"fmt"
	"testing"

	"{{ .ProjectName }}/src/domain/entities"
	db "{{ .ProjectName }}/src/domain/gateways/todo/sqlite"
	"{{ .ProjectName }}/src/infra/conf"
	"{{ .ProjectName }}/src/infra/datastores"
	"{{ .ProjectName }}/src/infra/test"
	"{{ .ProjectName }}/src/infra/test/assert"
	"{{ .ProjectName }}/src/interfaces"
)

func TestTodoItemGatewayCreate(ts *testing.T) {
	datastore := datastores.With(conf.Get()).New()

	test.WithSqlite(ts, "Should create required fields", func(t *testing.T, store interfaces.ISqlDataStore, ctx context.Context) {
		item := entities.NewTodo()
		item.Title = "My todo_item"
		gateway := db.New(datastore)
		_, err := gateway.Create(ctx, item)
		assert.Nil(t, err, fmt.Sprintf("%s", err))

		var found db.TodoRow
		query := "select * from todo_items where id=?"
		assert.Nil(t, store.Query(query, &found, item.ID))

		assert.Equal(t, found.ID, item.ID)
		assert.Equal(t, found.Title, item.Title)
		assert.False(t, found.IsDone)
	})
}
