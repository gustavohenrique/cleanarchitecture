package todo_test

import (
	"context"
	"testing"

	"{{ .ProjectName }}/src/infra/frameworks/httpserver"
	"{{ .ProjectName }}/src/infra/frameworks/httpserver/controllers/todo"
	"{{ .ProjectName }}/src/infra/test"
	"{{ .ProjectName }}/src/infra/test/assert"
	"{{ .ProjectName }}/src/infra/test/httpclient"
	"{{ .ProjectName }}/src/interfaces"
)

func TestReadAll(ts *testing.T) {
	test.WithSqlite(ts, "Should fetch all TODO items", func(t *testing.T, store interfaces.ISqlDataStore, ctx context.Context) {
		todoItem1 := todo.TodoJsonRequest{}
		todoItem1.ID = "2bbb00bf-b4f5-4746-9544-dc1ff07671ef"
		todoItem1.Title = "Schedule a meeting"
		todoItem1.IsDone = true
		sql := "insert into todo_items (id, title, is_done) values (?, ?, ?)"
		err := store.Exec(sql, todoItem1.ID, todoItem1.Title, todoItem1.IsDone)
		assert.Nil(t, err, "Could not insert")

		todoItem2 := todo.TodoJsonRequest{}
		todoItem2.ID = "dffa84f1-b0be-4191-ba03-5f5a6edd7979"
		todoItem2.Title = "Check opened PRs"
		err = store.Exec(sql, todoItem2.ID, todoItem2.Title, todoItem2.IsDone)
		assert.Nil(t, err, "Could not insert")

		req := httpclient.
			WithOauth2Mock().
			DoGET("/v1/todo")
		services := test.NewTestHelper().WithOauth2Mock(t).GetServices()
		resp, body := httpserver.WithStub(services).ServeHTTP(req)
		assert.HttpStatusCode(t, resp, 200)
		var items []todo.TodoJsonResponse
		body.To(&items)
		assert.Equal(t, 2, len(items))
	})
}
