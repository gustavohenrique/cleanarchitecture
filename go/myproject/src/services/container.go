package services

import (
	"myproject/src/interfaces"
	"myproject/src/repositories"
	"myproject/src/services/todo"
)

type ServiceContainer struct {
	todoService interfaces.ITodoService
}

func New(repositoryContainer repositories.RepositoryContainer) ServiceContainer {
	return ServiceContainer{
		todoService: todo.NewService(repositoryContainer),
	}
}

func (c ServiceContainer) GetTodoService() interfaces.ITodoService {
	return c.todoService
}
