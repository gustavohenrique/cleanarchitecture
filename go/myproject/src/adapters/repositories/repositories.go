package repositories

import (
	"{{ .ProjectName }}/src/domain/ports"
	"{{ .ProjectName }}/src/infrastructure/datastores"
)

type RepositoriesContainer struct {
	todoRepository ports.TodoRepository
}

func New(datastores datastores.Stores) ports.Repositories {
	repos := &RepositoriesContainer{}
	repos.todoRepository = NewTodoRepository(datastores.Postgres())
	return repos
}

func (repos *RepositoriesContainer) TodoRepository() ports.TodoRepository {
	return repos.todoRepository
}
