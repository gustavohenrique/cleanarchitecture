package clickhouse

import (
	"context"

	"github.com/ClickHouse/clickhouse-go/v2"

	"clickhousepoc/src/shared/conf"
)

type ClickHouseStore struct {
	connection *clickhouse.Connection
	tx         *clickhouse.Transaction
	config     *conf.Config
	ctx        context.Context
}

func NewClickHouseStore() *ClickHouseStore {
	return &ClickHouseStore{
		config: conf.Get(),
		ctx:    context.Background(),
	}
}

func (db *ClickHouseStore) Connect() (*clickhouse.Connection, error) {
	config := db.config
	clickhouseConfig, err := clickhouse.ParseDSN(config.Store.ClickHouse.URL)
	if err != nil {
		return nil, err
	}
	clickhouseConfig.Debug = config.Debug
	conn := clickhouse.OpenDB(clickhouseConfig)
	conn.SetMaxIdleConns(config.Store.ClickHouse.MaxConnIddleTime)
	conn.SetMaxOpenConns(config.Store.ClickHouse.MaxConns)
	conn.SetConnMaxLifetime(config.Store.ClickHouse.MaxConnLifetime)

	db.connection = conn
	return conn, nil
}

func (db *ClickHouseStore) WithContext(ctx context.Context) *ClickHouseStore {
	db.ctx = clickhouse.Context(ctx, clickhouse.WithSettings(clickhouse.Settings{
		"max_block_size": 10,
	}))
	return db
}

func (db *ClickHouseStore) ApplySchemaAndDropData(schema string) error {
	conn, err := db.getConnection()
	if err != nil {
		return err
	}
	_, err = conn.ExecContext(db.getCtx(), schema)
	return err
}

func (db *ClickHouseStore) Get(query string, found interface{}, args ...interface{}) error {
	conn, err := db.getConnection()
	if err != nil {
		return err
	}
	return db.QueryOne(query, found, args...)
}

func (db *ClickHouseStore) QueryOne(query string, found interface{}, args ...interface{}) error {
	conn, err := db.getConnection()
	if err != nil {
		return err
	}
	err := conn.QueryRow(db.getCtx(), query, args...).ScanStruct(&found)
	return err
}

func (db *ClickHouseStore) Query(query string, found interface{}, args ...interface{}) error {
	conn, err := db.getConnection()
	if err != nil {
		return found, err
	}
	rows, err := conn.QueryContext(db.getCtx(), query, args...)
	if err != nil {
		return found, err
	}
	return rows.ScanStruct(&found)
}

func (db *ClickHouseStore) QueryAll(query string, found interface{}, args ...interface{}) error {
	return db.Query(query, found, args...)
}

func (db *ClickHouseStore) Exec(query string, args ...interface{}) error {
	conn, err := db.getConnection()
	if err != nil {
		return err
	}
	return conn.ExecContext(db.getCtx(), query, args...)
}

func (db *ClickHouseStore) getConnection() (*clickhouse.Connection, error) {
	var err error
	if db.connection == nil {
		db.connection, err = db.Connect()
	}
	return db.connection, err
}

func (db *ClickHouseStore) getCtx() context.Context {
	if db.ctx != nil {
		return db.ctx
	}
	return context.Background()
}
