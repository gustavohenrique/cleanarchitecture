package sqlite

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"

	"{{ .ProjectName }}/src/interfaces"
)

type Config struct {
	Address string
}

type SqliteStore struct {
	connection *sqlx.DB
	config     Config
	ctx        context.Context
}

func New(config Config) interfaces.ISqlDataStore {
	return &SqliteStore{
		config: config,
		ctx:    context.Background(),
	}
}

func (db *SqliteStore) Connect() error {
	conn, err := sqlx.ConnectContext(db.getCtx(), "sqlite3", db.config.Address)
	if err != nil {
		return err
	}
	db.connection = conn
	return nil
}

func (db *SqliteStore) WithContext(ctx context.Context) interfaces.ISqlDataStore {
	db.ctx = ctx
	return db
}

func (db *SqliteStore) ApplySchemaAndDropData(schema string) error {
	conn, err := db.getConnection()
	if err != nil {
		return err
	}
	_, err = conn.Exec(schema)
	return err
}

func (db *SqliteStore) Get(query string, found interface{}, args ...interface{}) error {
	conn, err := db.getConnection()
	if err != nil {
		return err
	}
	return conn.GetContext(db.getCtx(), found, query, args...)
}

func (db *SqliteStore) QueryOne(query string, found interface{}, args ...interface{}) error {
	return db.Get(query, found, args)
}

func (db *SqliteStore) Query(query string, found interface{}, args ...interface{}) error {
	conn, err := db.getConnection()
	if err != nil {
		return err
	}
	err = conn.QueryRowxContext(db.getCtx(), query, args...).StructScan(found)
	return err
}

func (db *SqliteStore) QueryAll(query string, found interface{}, args ...interface{}) error {
	conn, err := db.getConnection()
	if err != nil {
		return err
	}
	err = conn.SelectContext(db.getCtx(), found, query, args...)
	return err
}

func (db *SqliteStore) Exec(query string, args ...interface{}) error {
	conn, err := db.getConnection()
	if err != nil {
		return err
	}
	result, err := conn.ExecContext(db.getCtx(), query, args...)
	if err != nil {
		return err
	}
	rows, err := result.RowsAffected()
	if rows == 0 || err != nil {
		return fmt.Errorf("no rows affected")
	}
	return nil
}

func (db *SqliteStore) ExecAndReturnID(query string, args ...interface{}) (string, error) {
	conn, err := db.getConnection()
	if err != nil {
		return "", err
	}
	result, err := conn.ExecContext(db.getCtx(), query, args...)
	if err != nil {
		return "", err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%d", id), nil
}

func (db *SqliteStore) ExecAndReturnRowsAffected(query string, args ...interface{}) (int64, error) {
	conn, err := db.getConnection()
	if err != nil {
		return 0, err
	}
	result, err := conn.ExecContext(db.getCtx(), query, args...)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func (db *SqliteStore) BulkInsert(table string, columns []string, rows [][]interface{}) error {
	// Not implemented
	return nil
}

func (db *SqliteStore) getConnection() (*sqlx.DB, error) {
	var err error
	if db.connection == nil {
		if err := db.Connect(); err != nil {
			return db.connection, err
		}
	}
	return db.connection, err
}

func (db *SqliteStore) getCtx() context.Context {
	if db.ctx != nil {
		return db.ctx
	}
	return context.Background()
}
