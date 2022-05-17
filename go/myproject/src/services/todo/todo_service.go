package todo

import (
	"context"

	"myproject/src/entities"
	"myproject/src/interfaces"
	"myproject/src/repositories"
	"myproject/src/shared/uuid"
)

type TodoService struct {
	repositoryContainer repositories.RepositoryContainer
}

func NewService(repositoryContainer repositories.RepositoryContainer) interfaces.ITodoService {
	return &TodoService{
		repositoryContainer: repositoryContainer,
	}
}

func (s *TodoService) ReadAll(ctx context.Context) ([]entities.TodoItemEntity, error) {
	todoRepository := s.repositoryContainer.GetTodoRepository(ctx)
	return todoRepository.ReadAll(ctx)
}

func (s *TodoService) Create(ctx context.Context, item entities.TodoItemEntity) (entities.TodoItemEntity, error) {
	todoRepository := s.repositoryContainer.GetTodoRepository(ctx)
	item.ID = uuid.NewV4()
	return todoRepository.Create(ctx, item)
}

func (s *TodoService) Update(ctx context.Context, item entities.TodoItemEntity) (entities.TodoItemEntity, error) {
	todoRepository := s.repositoryContainer.GetTodoRepository(ctx)
	return todoRepository.Update(ctx, item)
}

func (s *TodoService) Remove(ctx context.Context, item entities.TodoItemEntity) error {
	todoRepository := s.repositoryContainer.GetTodoRepository(ctx)
	return todoRepository.Remove(ctx, item)
}
