package repositories

import (
	"context"

	"{{ .ProjectName }}/src/domain/models"
	"{{ .ProjectName }}/src/domain/ports"
	"{{ .ProjectName }}/src/infrastructure/datastores/db"
	"{{ .ProjectName }}/src/wire"
)

type TodoRepository struct {
	store db.SqlDataStore
}

func NewTodoRepository(store db.SqlDataStore) ports.TodoRepository {
	return &TodoRepository{store}
}

func (r TodoRepository) Create(ctx context.Context, model models.TodoModel) (models.TodoModel, error) {
	q := "INSERT INTO todo_items (id, title) VALUES ($1, $2)"
	err := r.store.WithContext(ctx).Exec(q,
		model.ID,
		model.Title,
	)
	return model, err
}

func (r TodoRepository) ReadOne(ctx context.Context, model models.TodoModel) (models.TodoModel, error) {
	var item wire.TodoEntity
	q := "SELECT id, title, created_at, is_done FROM table_items WHERE id=$1 LIMIT 1"
	err := r.store.WithContext(ctx).QueryAll(q, &item, model.ID)
	return item.ToModel(), err
}

func (r TodoRepository) Update(ctx context.Context, model models.TodoModel) (models.TodoModel, error) {
	q := "UPDATE todo_items SET title=$2, is_done=$3 WHERE id=$1"
	err := r.store.WithContext(ctx).Exec(
		q,
		model.ID,
		model.Title,
		model.IsDone,
	)
	return model, err
}

func (r TodoRepository) Delete(ctx context.Context, model models.TodoModel) (models.TodoModel, error) {
	q := "DELETE FROM todo_items WHERE id=$1 LIMIT 1"
	err := r.store.WithContext(ctx).Exec(q, model.ID)
	return model, err
}
