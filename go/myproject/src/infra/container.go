package infra

import (
	"myproject/src/infra/dgraph"
	"myproject/src/infra/postgres"
	"myproject/src/infra/sqlite"
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
