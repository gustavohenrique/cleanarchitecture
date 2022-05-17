package sqlite

import (
	"context"
	"fmt"

	"myproject/src/adapters"
	"myproject/src/entities"
	"myproject/src/infra"
	st "myproject/src/infra/sqlite"
	"myproject/src/interfaces"
	"myproject/src/valueobjects"
)

const TODOITEMS = "todo_items"

type TodoRepository struct {
	store           *st.SqliteStore
	todoItemAdapter adapters.TodoItemAdapter
}

func NewRepository(infraContainer infra.InfraContainer) interfaces.ITodoRepository {
	return &TodoRepository{
		store:           infraContainer.SqliteStore,
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
	q := "INSERT INTO " + TODOITEMS + " (id, title) VALUES (?, ?)"
	err := r.store.WithContext(ctx).Exec(q,
		entity.ID,
		entity.Title,
	)
	return entity, err
}

func (r *TodoRepository) Update(ctx context.Context, entity entities.TodoItemEntity) (entities.TodoItemEntity, error) {
	q := "UPDATE " + TODOITEMS + " SET title=?, is_done=? WHERE id=?"
	err := r.store.WithContext(ctx).Exec(
		q,
		entity.Title,
		entity.IsDone,
		entity.ID,
	)
	return entity, err
}

func (r *TodoRepository) Remove(ctx context.Context, entity entities.TodoItemEntity) error {
	q := "DELETE FROM " + TODOITEMS + " WHERE id=?"
	return r.store.WithContext(ctx).Exec(q, entity.ID)
}
