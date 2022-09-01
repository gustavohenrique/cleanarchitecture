package sqlite

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"

	"{{ .ProjectName }}/src/shared/conf"
	"{{ .ProjectName }}/src/shared/customerror"
)

type SqliteStore struct {
	connection *sqlx.DB
	config     *conf.Config
	ctx        context.Context
}

func NewSqliteStore() *SqliteStore {
	return &SqliteStore{
		config: conf.Get(),
		ctx:    context.Background(),
	}
}

func (db *SqliteStore) Connect() (*sqlx.DB, error) {
	conn, err := sqlx.ConnectContext(db.getCtx(), "sqlite3", db.config.Store.Sqlite.Address)
	if err != nil {
		return nil, err
	}
	return conn, nil
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
		return customerror.NotFound("No rows affected")
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

func (db *SqliteStore) WithContext(ctx context.Context) *SqliteStore {
	db.ctx = ctx
	return db
}

func (db *SqliteStore) getConnection() (*sqlx.DB, error) {
	var err error
	if db.connection == nil {
		db.connection, err = db.Connect()
	}
	return db.connection, err
}
func (db *SqliteStore) getCtx() context.Context {
	if db.ctx != nil {
		return db.ctx
	}
	return context.Background()
}
