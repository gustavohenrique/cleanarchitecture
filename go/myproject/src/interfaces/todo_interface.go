package interfaces

import (
	"context"

	"{{ .ProjectName }}/src/domain/entities"
)

type ITodoGateway interface {
	Postgres() ISqlTodoGateway
	Sqlite() ISqlTodoGateway
}

type ISqlTodoGateway interface {
	ReadAll(ctx context.Context) ([]entities.Todo, error)
	ReadOne(ctx context.Context, entity entities.Todo) (entities.Todo, error)
	Create(ctx context.Context, entity entities.Todo) (entities.Todo, error)
	Update(ctx context.Context, entity entities.Todo) (entities.Todo, error)
	Remove(ctx context.Context, entity entities.Todo) error
}

type ITodoService interface {
	ReadAll(ctx context.Context) ([]entities.Todo, error)
	ReadOne(ctx context.Context, entity entities.Todo) (entities.Todo, error)
	Create(ctx context.Context, entity entities.Todo) (entities.Todo, error)
	Update(ctx context.Context, entity entities.Todo) (entities.Todo, error)
	Remove(ctx context.Context, entity entities.Todo) error
}
