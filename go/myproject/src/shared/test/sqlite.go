package test

import (
	"context"
	"io/ioutil"
	"log"
	"os"
	"testing"

	"myproject/src/infra/sqlite"
	"myproject/src/shared/conf"
)

type SqliteFN func(t *testing.T, store *sqlite.SqliteStore, ctx context.Context)

func WithSqlite(ts *testing.T, name string, fn SqliteFN) {
	store := sqlite.NewSqliteStore()
	store.Connect()
	key := "SQLITE_SCHEMA_FILE"
	file := os.Getenv(key)
	if file == "" {
		file = conf.Get().Store.Sqlite.Schema
	}
	b, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatalf("[ERROR] Cannot read schema file %s for Sqlite: %s", file, err)
	}
	if err := store.ApplySchemaAndDropData(string(b)); err != nil {
		log.Fatalf("[ERROR] Cannot apply schema. %s", err)
	}
	ctx := context.WithValue(context.Background(), "db", "sqlite")
	ts.Run(name, func(t *testing.T) {
		fn(t, store, ctx)
	})
}
