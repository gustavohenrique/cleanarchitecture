package todo_test

import (
	"context"
	"testing"

	"myproject/src/application/httpserver"
	"myproject/src/entities"
	"myproject/src/infra/sqlite"
	"myproject/src/shared/test"
	"myproject/src/shared/test/assert"
	"myproject/src/shared/test/httpclient"
	"myproject/src/valueobjects"
)

func TestReadAll(ts *testing.T) {
	test.WithSqlite(ts, "Should fetch all TODO items", func(t *testing.T, store *sqlite.SqliteStore, ctx context.Context) {
		todoItem1 := entities.TodoItemEntity{}
		todoItem1.ID = "2bbb00bf-b4f5-4746-9544-dc1ff07671ef"
		todoItem1.Title = "Marcar reuniao sobre produto"
		todoItem1.IsDone = true
		sql := "insert into todo_items (id, title, is_done) values (?, ?, ?)"
		err := store.Exec(sql, todoItem1.ID, todoItem1.Title, todoItem1.IsDone)
		assert.Nil(t, err, "Could not insert")

		todoItem2 := entities.TodoItemEntity{}
		todoItem2.ID = "dffa84f1-b0be-4191-ba03-5f5a6edd7979"
		todoItem2.Title = "Verificar PRs pendentes de aprovacao"
		err = store.Exec(sql, todoItem2.ID, todoItem2.Title, todoItem2.IsDone)
		assert.Nil(t, err, "Could not insert")

		req := httpclient.DoGET("/todo/")
		serviceContainer := test.GetServiceContainer()
		resp, body := httpserver.With(&serviceContainer).ServeHTTP(req)
		assert.HttpStatusCode(t, resp, 200)
		var items []valueobjects.TodoItemResponse
		body.To(&items)
		assert.Equal(t, 2, len(items))
	})
}
