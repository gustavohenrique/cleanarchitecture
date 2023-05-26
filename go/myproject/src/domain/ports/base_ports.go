package ports

import (
	"context"

	"{{ .ProjectName }}/src/domain/models"
)

type TodoUseCase interface {
	Execute(ctx context.Context, model models.TodoModel) (models.TodoModel, error)
}

type TodoRepository interface {
	Create(ctx context.Context, model models.TodoModel) (models.TodoModel, error)
	ReadOne(ctx context.Context, model models.TodoModel) (models.TodoModel, error)
	Update(ctx context.Context, model models.TodoModel) (models.TodoModel, error)
	Delete(ctx context.Context, model models.TodoModel) (models.TodoModel, error)
}
