package interfaces

import "context"

type IDataStore interface {
	New() IDataStore
	SQL() ISqlDataStore
	Postgres() ISqlDataStore
	Sqlite() ISqlDataStore
	Dgraph() IGraphDataStore
}

type ISqlDataStore interface {
	WithContext(ctx context.Context) ISqlDataStore
	Connect() error
	ApplySchemaAndDropData(schema string) error
	Get(query string, found interface{}, args ...interface{}) error
	Query(query string, found interface{}, args ...interface{}) error
	QueryOne(query string, found interface{}, args ...interface{}) error
	QueryAll(query string, found interface{}, args ...interface{}) error
	Exec(query string, args ...interface{}) error
	ExecAndReturnID(query string, args ...interface{}) (string, error)
	ExecAndReturnRowsAffected(query string, args ...interface{}) (int64, error)
	BulkInsert(table string, columns []string, rows [][]interface{}) error
}

type IGraphDataStore interface {
	WithContext(ctx context.Context) IGraphDataStore
	Connect() error
	ApplySchemaAndDropData(schema string) error
	Query(dest interface{}, dql string) error
	QueryWithVars(dql string, vars map[string]string) ([]byte, error)
	Mutate(dql string) error
	MutateReturnUids(dql string) ([]string, error)
	MutateReturnUid(dql string) (string, error)
}
