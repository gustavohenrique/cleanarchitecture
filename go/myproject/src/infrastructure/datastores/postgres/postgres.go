package postgres

import (
	"context"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"{{ .ProjectName }}/src/infrastructure/datastores/db"
)

type PostgresStore struct {
	connection *sqlx.DB
	config     Config
	ctx        context.Context
	bulk       *BulkStore
}

type Config struct {
	URL             string
	MaxOpenConns    int
	MaxIdleConns    int
	MaxConnLifetime int
}

func New(config Config) db.SqlDataStore {
	return &PostgresStore{
		config: config,
		ctx:    context.Background(),
		bulk:   Bulk(context.Background(), config),
	}
}

func (store *PostgresStore) Connect() error {
	config := store.config
	conn, err := sqlx.ConnectContext(store.getCtx(), "postgres", config.URL)
	if err != nil {
		return err
	}
	if config.MaxOpenConns > 0 {
		conn.SetMaxOpenConns(config.MaxOpenConns)
	}
	if config.MaxIdleConns > 0 {
		conn.SetMaxIdleConns(config.MaxIdleConns)
	}
	if config.MaxConnLifetime > 0 {
		conn.SetConnMaxLifetime(time.Second * time.Duration(config.MaxConnLifetime))
	}
	store.connection = conn
	return nil
}

func (store *PostgresStore) WithContext(ctx context.Context) db.SqlDataStore {
	store.ctx = ctx
	return store
}

func (store *PostgresStore) ApplySchemaAndDropData(schema string) error {
	conn, err := store.getConnection()
	if err != nil {
		return err
	}
	_, err = conn.Exec(schema)
	return err
}

func (store *PostgresStore) Get(query string, found interface{}, args ...interface{}) error {
	conn, err := store.getConnection()
	if err != nil {
		return err
	}
	return conn.GetContext(store.getCtx(), found, query, args...)
}

func (store *PostgresStore) QueryOne(query string, found interface{}, args ...interface{}) error {
	return store.Get(query, found, args...)
}

func (store *PostgresStore) Query(query string, found interface{}, args ...interface{}) error {
	conn, err := store.getConnection()
	if err != nil {
		return err
	}
	err = conn.QueryRowxContext(store.getCtx(), query, args...).StructScan(found)
	return err
}

func (store *PostgresStore) QueryAll(query string, found interface{}, args ...interface{}) error {
	conn, err := store.getConnection()
	if err != nil {
		return err
	}
	err = conn.SelectContext(store.getCtx(), found, query, args...)
	return err
}

func (store *PostgresStore) Exec(query string, args ...interface{}) error {
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

func (store *PostgresStore) ExecAndReturnID(query string, args ...interface{}) (string, error) {
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

func (store *PostgresStore) ExecAndReturnRowsAffected(query string, args ...interface{}) (int64, error) {
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

func (store *PostgresStore) BulkInsert(table string, columns []string, rows [][]interface{}) error {
	return store.bulk.Copy(table, columns, rows)
}

func (store *PostgresStore) getConnection() (*sqlx.DB, error) {
	var err error
	if store.connection == nil {
		if err := store.Connect(); err != nil {
			return store.connection, err
		}
	}
	return store.connection, err
}
func (store *PostgresStore) getCtx() context.Context {
	if store.ctx != nil {
		return store.ctx
	}
	return context.Background()
}
