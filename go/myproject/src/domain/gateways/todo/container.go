package todo

import (
	"{{ .ProjectName }}/src/domain/gateways/todo/postgres"
	"{{ .ProjectName }}/src/domain/gateways/todo/sqlite"
	"{{ .ProjectName }}/src/interfaces"
)

type TodoGateway struct {
	postgresImpl interfaces.ISqlTodoGateway
	sqliteImpl   interfaces.ISqlTodoGateway
}

func New(store interfaces.IDataStore) interfaces.ITodoGateway {
	return TodoGateway{
		postgresImpl: postgres.New(store),
		sqliteImpl:   sqlite.New(store),
	}
}

func (g TodoGateway) Postgres() interfaces.ISqlTodoGateway {
	return g.postgresImpl
}

func (g TodoGateway) Sqlite() interfaces.ISqlTodoGateway {
	return g.sqliteImpl
}
