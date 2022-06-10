package postgres_test

import (
	"context"
	"fmt"
	"testing"

	"{{ .ProjectName }}/src/entities"
	"{{ .ProjectName }}/src/infra"
	"{{ .ProjectName }}/src/infra/postgres"
	db "{{ .ProjectName }}/src/repositories/todo/postgres"
	"{{ .ProjectName }}/src/shared/test"
	"{{ .ProjectName }}/src/shared/test/assert"
	"{{ .ProjectName }}/src/valueobjects"
)

func TestTodoItemRepositoryCreate(ts *testing.T) {
	test.WithPostgres(ts, "Should create required fields", func(t *testing.T, store *postgres.PostgresStore, ctx context.Context) {
		item := entities.NewTodoItemEntity()
		item.Title = "My todo_item"
		repo := db.NewRepository(infra.New())
		_, err := repo.Create(ctx, item)
		assert.Nil(t, err, fmt.Sprintf("%s", err))

		var found valueobjects.TodoItemTable
		query := "select * from todo_items where id=$1"
		assert.Nil(t, store.Query(query, &found, item.ID))

		assert.Equal(t, found.ID, item.ID)
		assert.Equal(t, found.Title, item.Title)
		assert.False(t, found.IsDone)
	})
}
