package usecases

import (
	"context"

	"{{ .ProjectName }}/src/domain/models"
	"{{ .ProjectName }}/src/domain/ports"
)

type CreateTodoUseCase struct {
	todoRepository ports.TodoRepository
}

func NewCreateTodoUseCase(repos ports.Repositories) CreateTodoUseCase {
	return CreateTodoUseCase{
		todoRepository: repos.TodoRepository(),
	}
}

func (u CreateTodoUseCase) Execute(ctx context.Context, model models.TodoModel) (models.TodoModel, error) {
	return u.todoRepository.Create(ctx, model)
}
