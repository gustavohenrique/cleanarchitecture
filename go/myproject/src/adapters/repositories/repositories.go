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
	{{ if .HasPostgres }}
	repos.todoRepository = NewTodoPostgresRepository(datastores.Postgres())
	{{ else }}
		{{ if .HasSqlite }}
	repos.todoRepository = NewTodoSqliteRepository(datastores.Sqlite())
		{{ else }}
			{{ if .HasDgraph }}
	repos.todoRepository = NewTodoDgraphRepository(datastores.Dgraph())
			{{ end }}
		{{ end }}
	{{ end }}
	return repos
}

func (repos *RepositoriesContainer) TodoRepository() ports.TodoRepository {
	return repos.todoRepository
}
