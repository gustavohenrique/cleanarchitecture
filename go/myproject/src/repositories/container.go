package repositories

import (
	"context"

	"{{ .ProjectName }}/src/infra"
	"{{ .ProjectName }}/src/interfaces"
	todoPostgres "{{ .ProjectName }}/src/repositories/todo/postgres"
	todoSqlite "{{ .ProjectName }}/src/repositories/todo/sqlite"
	"{{ .ProjectName }}/src/shared/strings"
)

const (
	CONTEXT_KEY = "db"
	SQLITE      = "sqlite"
	POSTGRES    = "postgres"
	DEFAULT     = SQLITE
)

var ENGINES = []string{SQLITE, POSTGRES}

type RepositoryContainer struct {
	todoRepositories map[string]interfaces.ITodoRepository
}

func New(infraContainer infra.InfraContainer) RepositoryContainer {
	todoRepositories := map[string]interfaces.ITodoRepository{
		SQLITE:   todoSqlite.NewRepository(infraContainer),
		POSTGRES: todoPostgres.NewRepository(infraContainer),
	}
	return RepositoryContainer{
		todoRepositories: todoRepositories,
	}
}

func getDbEngineFrom(ctx context.Context) string {
	v := ctx.Value(CONTEXT_KEY)
	if v == nil {
		return DEFAULT
	}
	engine := v.(string)
	if strings.SliceContains(ENGINES, engine) {
		return engine
	}
	return DEFAULT
}

func (c RepositoryContainer) GetTodoRepository(ctx context.Context) interfaces.ITodoRepository {
	return c.todoRepositories[getDbEngineFrom(ctx)]
}
