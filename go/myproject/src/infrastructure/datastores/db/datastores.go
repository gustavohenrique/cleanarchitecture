package db

import "context"

type SqlDataStore interface {
	WithContext(ctx context.Context) SqlDataStore
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

type GraphDataStore interface {
	WithContext(ctx context.Context) GraphDataStore
	Connect() error
	ApplySchemaAndDropData(schema string) error
	Query(dest interface{}, dql string) error
	QueryWithVars(dql string, vars map[string]string) ([]byte, error)
	Mutate(dql string) error
	MutateReturnUids(dql string) ([]string, error)
	MutateReturnUid(dql string) (string, error)
}
