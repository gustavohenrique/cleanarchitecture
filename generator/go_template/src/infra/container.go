package infra

import (
	"{{ .ProjectName }}/src/infra/dgraph"
	"{{ .ProjectName }}/src/infra/postgres"
	"{{ .ProjectName }}/src/infra/sqlite"
)

type InfraContainer struct {
	PostgresStore *postgres.PostgresStore
	DgraphStore   *dgraph.DgraphStore
	SqliteStore   *sqlite.SqliteStore
}

func New() InfraContainer {
	return InfraContainer{
		PostgresStore: postgres.NewPostgresStore(),
		DgraphStore:   dgraph.NewDgraphStore(),
		SqliteStore:   sqlite.NewSqliteStore(),
	}
}
