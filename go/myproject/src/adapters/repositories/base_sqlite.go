package repositories

import (
	"context"

	"{{ .ProjectName }}/src/domain/models"
	"{{ .ProjectName }}/src/domain/ports"
	"{{ .ProjectName }}/src/infrastructure/datastores/db"
	"{{ .ProjectName }}/src/wire"
)

type TodoSqliteRepository struct {
	store db.SqlDataStore
}

func NewTodoSqliteRepository(store db.SqlDataStore) ports.TodoRepository {
	return &TodoSqliteRepository{store}
}

func (r TodoSqliteRepository) Create(ctx context.Context, model models.TodoModel) (models.TodoModel, error) {
	q := "INSERT INTO todo_items (id, title) VALUES (?, ?)"
	err := r.store.WithContext(ctx).Exec(q,
		model.ID,
		model.Title,
	)
	return model, err
}

func (r TodoSqliteRepository) ReadOne(ctx context.Context, model models.TodoModel) (models.TodoModel, error) {
	var item wire.TodoEntity
	q := "SELECT id, title, created_at, is_done FROM table_items WHERE id=? LIMIT 1"
	err := r.store.WithContext(ctx).QueryAll(q, &item, model.ID)
	return item.ToModel(), err
}

func (r TodoSqliteRepository) Update(ctx context.Context, model models.TodoModel) (models.TodoModel, error) {
	q := "UPDATE todo_items SET title=?, is_done=? WHERE id=?"
	err := r.store.WithContext(ctx).Exec(
		q,
		model.ID,
		model.Title,
		model.IsDone,
	)
	return model, err
}

func (r TodoSqliteRepository) Delete(ctx context.Context, model models.TodoModel) (models.TodoModel, error) {
	q := "DELETE FROM todo_items WHERE id=? LIMIT 1"
	err := r.store.WithContext(ctx).Exec(q, model.ID)
	return model, err
}
