package services

import (
	"{{ .ProjectName }}/src/interfaces"
	"{{ .ProjectName }}/src/repositories"
	"{{ .ProjectName }}/src/services/auth"
	"{{ .ProjectName }}/src/services/todo"
)

type ServiceContainer struct {
	todoService interfaces.ITodoService
	authService interfaces.IAuthService
}

func New(repositoryContainer repositories.RepositoryContainer) ServiceContainer {
	return ServiceContainer{
		todoService: todo.NewService(repositoryContainer),
		authService: auth.NewService(repositoryContainer),
	}
}

func (c ServiceContainer) GetTodoService() interfaces.ITodoService {
	return c.todoService
}

func (c ServiceContainer) GetAuthService() interfaces.IAuthService {
	return c.authService
}
