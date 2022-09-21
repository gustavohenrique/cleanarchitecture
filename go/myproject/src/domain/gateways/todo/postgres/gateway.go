package postgres

import (
	"context"

	"{{ .ProjectName }}/src/domain/entities"
	"{{ .ProjectName }}/src/interfaces"
)

type TodoGateway struct {
	store interfaces.ISqlDataStore
}

func New(store interfaces.IDataStore) interfaces.ISqlTodoGateway {
	return &TodoGateway{
		store: store.Postgres(),
	}
}

func (r *TodoGateway) ReadAll(ctx context.Context) ([]entities.Todo, error) {
	var items []TodoRow
	q := "SELECT id, title, created_at, is_done FROM table_items ORDER BY created_at, is_done LIMIT 100"
	if err := r.store.WithContext(ctx).QueryAll(q, &items); err != nil {
		return []entities.Todo{}, err
	}
	return ToEntities(items), nil
}

func (r *TodoGateway) ReadOne(ctx context.Context, entity entities.Todo) (entities.Todo, error) {
	var item TodoRow
	q := "SELECT id, title, created_at, is_done FROM table_items WHERE id=$1 LIMIT 1"
	err := r.store.WithContext(ctx).QueryAll(q, &item, entity.ID)
	return item.ToEntity(), err
}

func (r *TodoGateway) Create(ctx context.Context, entity entities.Todo) (entities.Todo, error) {
	q := "INSERT INTO todo_items (id, title) VALUES ($1, $2)"
	err := r.store.WithContext(ctx).Exec(q,
		entity.ID,
		entity.Title,
	)
	return entity, err
}

func (r *TodoGateway) Update(ctx context.Context, entity entities.Todo) (entities.Todo, error) {
	q := "UPDATE todo_items SET title=$2, is_done=$3 WHERE id=$1"
	err := r.store.WithContext(ctx).Exec(
		q,
		entity.ID,
		entity.Title,
		entity.IsDone,
	)
	return entity, err
}

func (r *TodoGateway) Remove(ctx context.Context, entity entities.Todo) error {
	q := "DELETE FROM todo_items WHERE id=$1 LIMIT 1"
	return r.store.WithContext(ctx).Exec(q, entity.ID)
}
