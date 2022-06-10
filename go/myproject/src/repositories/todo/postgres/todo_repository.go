package postgres

import (
	"context"
	"fmt"

	"{{ .ProjectName }}/src/adapters"
	"{{ .ProjectName }}/src/entities"
	"{{ .ProjectName }}/src/infra"
	st "{{ .ProjectName }}/src/infra/postgres"
	"{{ .ProjectName }}/src/interfaces"
	"{{ .ProjectName }}/src/valueobjects"
)

const TODOITEMS = "todo_items"

type TodoRepository struct {
	store           *st.PostgresStore
	todoItemAdapter adapters.TodoItemAdapter
}

func NewRepository(infraContainer infra.InfraContainer) interfaces.ITodoRepository {
	return &TodoRepository{
		store:           infraContainer.PostgresStore,
		todoItemAdapter: adapters.NewTodoItemAdapter(),
	}
}

func (r *TodoRepository) String() string {
	return TODOITEMS
}

func (r *TodoRepository) ReadAll(ctx context.Context) ([]entities.TodoItemEntity, error) {
	var items []valueobjects.TodoItemTable
	q := fmt.Sprintf(`SELECT id, title, created_at, is_done FROM %s`, TODOITEMS)
	if err := r.store.WithContext(ctx).QueryAll(q, &items); err != nil {
		return []entities.TodoItemEntity{}, err
	}
	return r.todoItemAdapter.FromTableListToEntityList(items), nil
}

func (r *TodoRepository) Create(ctx context.Context, entity entities.TodoItemEntity) (entities.TodoItemEntity, error) {
	q := "INSERT INTO " + TODOITEMS + " (id, title) VALUES ($1, $2)"
	err := r.store.WithContext(ctx).Exec(q,
		entity.ID,
		entity.Title,
	)
	return entity, err
}

func (r *TodoRepository) Update(ctx context.Context, entity entities.TodoItemEntity) (entities.TodoItemEntity, error) {
	q := "UPDATE " + TODOITEMS + " SET title=$2, is_done=$3 WHERE id=$1"
	err := r.store.WithContext(ctx).Exec(
		q,
		entity.ID,
		entity.Title,
		entity.IsDone,
	)
	return entity, err
}

func (r *TodoRepository) Remove(ctx context.Context, entity entities.TodoItemEntity) error {
	q := "DELETE FROM " + TODOITEMS + " WHERE id=$1"
	return r.store.WithContext(ctx).Exec(q, entity.ID)
}
