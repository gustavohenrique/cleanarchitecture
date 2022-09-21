package todo

import (
	"context"

	"{{ .ProjectName }}/src/domain/entities"
	"{{ .ProjectName }}/src/interfaces"
	"{{ .ProjectName }}/src/shared/uuid"
)

type TodoService struct {
	todoGateway interfaces.ISqlTodoGateway
}

func New(gateways interfaces.IGateway) interfaces.ITodoService {
	return &TodoService{
		todoGateway: gateways.TodoGateway().Sqlite(),
	}
}

func (s *TodoService) ReadAll(ctx context.Context) ([]entities.Todo, error) {
	return s.todoGateway.ReadAll(ctx)
}

func (s *TodoService) ReadOne(ctx context.Context, item entities.Todo) (entities.Todo, error) {
	return s.todoGateway.ReadOne(ctx, item)
}

func (s *TodoService) Create(ctx context.Context, item entities.Todo) (entities.Todo, error) {
	item.ID = uuid.NewV4()
	return s.todoGateway.Create(ctx, item)
}

func (s *TodoService) Update(ctx context.Context, item entities.Todo) (entities.Todo, error) {
	return s.todoGateway.Update(ctx, item)
}

func (s *TodoService) Remove(ctx context.Context, item entities.Todo) error {
	return s.todoGateway.Remove(ctx, item)
}
