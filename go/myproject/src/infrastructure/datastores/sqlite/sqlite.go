package sqlite

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"

	"{{ .ProjectName }}/src/infrastructure/datastores/db"
)

type Config struct {
	Address string
}

type SqliteStore struct {
	connection *sqlx.DB
	config     Config
	ctx        context.Context
}

func New(config Config) db.SqlDataStore {
	return &SqliteStore{
		config: config,
		ctx:    context.Background(),
	}
}

func (store *SqliteStore) Connect() error {
	conn, err := sqlx.ConnectContext(store.getCtx(), "sqlite3", store.config.Address)
	if err != nil {
		return err
	}
	store.connection = conn
	return nil
}

func (store *SqliteStore) WithContext(ctx context.Context) db.SqlDataStore {
	store.ctx = ctx
	return store
}

func (store *SqliteStore) ApplySchemaAndDropData(schema string) error {
	conn, err := store.getConnection()
	if err != nil {
		return err
	}
	_, err = conn.Exec(schema)
	return err
}

func (store *SqliteStore) Get(query string, found interface{}, args ...interface{}) error {
	conn, err := store.getConnection()
	if err != nil {
		return err
	}
	return conn.GetContext(store.getCtx(), found, query, args...)
}

func (store *SqliteStore) QueryOne(query string, found interface{}, args ...interface{}) error {
	return store.Get(query, found, args)
}

func (store *SqliteStore) Query(query string, found interface{}, args ...interface{}) error {
	conn, err := store.getConnection()
	if err != nil {
		return err
	}
	err = conn.QueryRowxContext(store.getCtx(), query, args...).StructScan(found)
	return err
}

func (store *SqliteStore) QueryAll(query string, found interface{}, args ...interface{}) error {
	conn, err := store.getConnection()
	if err != nil {
		return err
	}
	err = conn.SelectContext(store.getCtx(), found, query, args...)
	return err
}

func (store *SqliteStore) Exec(query string, args ...interface{}) error {
	conn, err := store.getConnection()
	if err != nil {
		return err
	}
	result, err := conn.ExecContext(store.getCtx(), query, args...)
	if err != nil {
		return err
	}
	rows, err := result.RowsAffected()
	if rows == 0 || err != nil {
		return fmt.Errorf("no rows affected")
	}
	return nil
}

func (store *SqliteStore) ExecAndReturnID(query string, args ...interface{}) (string, error) {
	conn, err := store.getConnection()
	if err != nil {
		return "", err
	}
	result, err := conn.ExecContext(store.getCtx(), query, args...)
	if err != nil {
		return "", err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%d", id), nil
}

func (store *SqliteStore) ExecAndReturnRowsAffected(query string, args ...interface{}) (int64, error) {
	conn, err := store.getConnection()
	if err != nil {
		return 0, err
	}
	result, err := conn.ExecContext(store.getCtx(), query, args...)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func (store *SqliteStore) BulkInsert(table string, columns []string, rows [][]interface{}) error {
	// Not implemented
	return nil
}

func (store *SqliteStore) getConnection() (*sqlx.DB, error) {
	var err error
	if store.connection == nil {
		if err := store.Connect(); err != nil {
			return store.connection, err
		}
	}
	return store.connection, err
}

func (store *SqliteStore) getCtx() context.Context {
	if store.ctx != nil {
		return store.ctx
	}
	return context.Background()
}
