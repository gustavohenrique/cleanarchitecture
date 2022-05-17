package postgres

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"

	"myproject/src/shared/conf"
)

type PostgresStore struct {
	connection *pgxpool.Pool
	pgxconf    *pgxpool.Config
	tx         pgx.Tx
}

type Row pgx.Row

func NewPostgresStore() *PostgresStore {
	config := conf.Get()
	databaseURL := config.Store.Postgres.URL
	pgxconf, err := pgxpool.ParseConfig(databaseURL)
	if err != nil {
		log.Fatalln("Cannot connect to postgres=", databaseURL, " error=", err)
	}
	db := &PostgresStore{}

	pgxconf.MaxConns = int32(config.Store.Postgres.MaxConns)
	pgxconf.MaxConnLifetime = time.Duration(config.Store.Postgres.MaxConnLifetime) * time.Second
	pgxconf.MaxConnIdleTime = time.Duration(config.Store.Postgres.MaxConnIdleTime) * time.Second

	if config.Log.Level != "" {
		level, err := pgx.LogLevelFromString(strings.ToLower(config.Log.Level))
		if err == nil {
			pgxconf.ConnConfig.LogLevel = level
		}
	}
	db.pgxconf = pgxconf
	return db
}

func (db *PostgresStore) SetLogAdapter(l pgx.Logger) {
	db.pgxconf.ConnConfig.Logger = l
}

func (db *PostgresStore) Connect(ctx context.Context) (*pgxpool.Pool, error) {
	return pgxpool.ConnectConfig(ctx, db.pgxconf)
}

func (db *PostgresStore) getConnection(ctx context.Context) (*pgxpool.Pool, error) {
	var err error
	if db.connection == nil {
		db.connection, err = db.Connect(ctx)
	}
	return db.connection, err
}

func (db *PostgresStore) Begin(ctx context.Context) (pgx.Tx, error) {
	conn, err := db.getConnection(ctx)
	if err != nil {
		return nil, err
	}
	tx, err := conn.Begin(ctx)
	db.tx = tx
	return tx, err
}

func (db *PostgresStore) GetTx() pgx.Tx {
	return db.tx
}

func (db *PostgresStore) Rollback(ctx context.Context) error {
	return db.GetTx().Rollback(ctx)
}

func (db *PostgresStore) Commit(ctx context.Context) error {
	return db.GetTx().Commit(ctx)
}

func (db *PostgresStore) Query(ctx context.Context, query string, args ...interface{}) (pgx.Rows, error) {
	conn, err := db.getConnection(ctx)
	if err != nil {
		return nil, err
	}
	return conn.Query(ctx, query, args...)
}

func (db *PostgresStore) QueryRow(ctx context.Context, query string, found interface{}, args ...interface{}) error {
	conn, err := db.getConnection(ctx)
	if err != nil {
		return err
	}
	var row string
	err = conn.QueryRow(ctx, query, args...).Scan(&row)
	if err != nil {
		return err
	}
	return json.Unmarshal([]byte(row), &found)
}

func (db *PostgresStore) QueryOne(ctx context.Context, query string, found interface{}, args ...interface{}) error {
	conn, err := db.getConnection(ctx)
	if err != nil {
		return err
	}
	return conn.QueryRow(ctx, query, args...).Scan(found)
}

func (db *PostgresStore) Exec(ctx context.Context, query string, args ...interface{}) error {
	conn, err := db.getConnection(ctx)
	if err != nil {
		return err
	}
	commandTag, err := conn.Exec(ctx, query, args...)
	if err != nil {
		return err
	}
	if commandTag.RowsAffected() == 0 {
		return fmt.Errorf("No rows affected")
	}
	return nil
}

func (db *PostgresStore) ForEach(rows pgx.Rows, f func(row string)) error {
	defer rows.Close()
	for rows.Next() {
		var row string
		err := rows.Scan(&row)
		if err != nil {
			return err
		}
		f(row)
	}
	if rows.Err() != nil {
		return fmt.Errorf("%v", rows.Err())
	}
	return nil
}
