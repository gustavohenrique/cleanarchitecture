package postgres

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type BulkStore struct {
	connection *pgxpool.Pool
	config     Config
	ctx        context.Context
}

func Bulk(ctx context.Context, config Config) *BulkStore {
	pgxconf, _ := pgxpool.ParseConfig(config.URL)
	maxOpenConns := config.MaxOpenConns
	maxConnLifetime := config.MaxConnLifetime
	maxIdleConns := config.MaxIdleConns
	if maxOpenConns > 0 {
		pgxconf.MaxConns = int32(maxOpenConns)
	}
	if maxConnLifetime > 0 {
		pgxconf.MaxConnLifetime = time.Duration(maxConnLifetime) * time.Second
	}
	if maxIdleConns > 0 {
		pgxconf.MaxConnIdleTime = time.Duration(maxIdleConns) * time.Second
	}
	connection, _ := pgxpool.ConnectConfig(ctx, pgxconf)
	if ctx == nil {
		ctx = context.Background()
	}
	return &BulkStore{
		connection: connection,
		config:     config,
		ctx:        ctx,
	}
}

func (db *BulkStore) Copy(table string, columns []string, rows [][]interface{}) error {
	if len(rows) == 0 || len(columns) == 0 {
		return fmt.Errorf("columns ou rows com length zero")
	}
	if len(columns) != len(rows[0]) {
		return fmt.Errorf("the total of cols is different from total of rows")
	}
	conn := db.connection
	if conn == nil {
		return fmt.Errorf("cannot connect to database %s", db.config.URL)
	}
	_, err := db.connection.CopyFrom(
		db.ctx,
		pgx.Identifier{table},
		columns,
		pgx.CopyFromRows(rows),
	)
	return err
}
