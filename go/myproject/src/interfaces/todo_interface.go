package interfaces

import (
	"context"

	"myproject/src/entities"
)

type ITodoRepository interface {
	ReadAll(ctx context.Context) ([]entities.TodoItemEntity, error)
	Create(ctx context.Context, entity entities.TodoItemEntity) (entities.TodoItemEntity, error)
	Update(ctx context.Context, entity entities.TodoItemEntity) (entities.TodoItemEntity, error)
	Remove(ctx context.Context, entity entities.TodoItemEntity) error
}

type ITodoService interface {
	ReadAll(ctx context.Context) ([]entities.TodoItemEntity, error)
	Create(ctx context.Context, entity entities.TodoItemEntity) (entities.TodoItemEntity, error)
	Update(ctx context.Context, entity entities.TodoItemEntity) (entities.TodoItemEntity, error)
	Remove(ctx context.Context, entity entities.TodoItemEntity) error
}
