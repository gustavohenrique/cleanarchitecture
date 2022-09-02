package services

import (
	"{{ .ProjectName }}/src/interfaces"
	"{{ .ProjectName }}/src/repositories"
	"{{ .ProjectName }}/src/services/auth"
	"{{ .ProjectName }}/src/services/todo"
)

type ServiceContainer struct {
	TodoService interfaces.ITodoService
	AuthService interfaces.IAuthService
}

func New(repositoryContainer repositories.RepositoryContainer) ServiceContainer {
	return ServiceContainer{
		TodoService: todo.NewService(repositoryContainer),
		AuthService: auth.NewService(repositoryContainer),
	}
}

func (c ServiceContainer) GetTodoService() interfaces.ITodoService {
	return c.TodoService
}

func (c ServiceContainer) GetAuthService() interfaces.IAuthService {
	return c.AuthService
}
