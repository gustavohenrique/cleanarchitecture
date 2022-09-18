package test

import (
	"context"
	"log"
	"os"
	"testing"

	"{{ .ProjectName }}/src/infra/conf"
	"{{ .ProjectName }}/src/infra/datastores/sqlite"
	"{{ .ProjectName }}/src/interfaces"
)

type SqliteFN func(t *testing.T, store interfaces.ISqlDataStore, ctx context.Context)

func WithSqlite(ts *testing.T, name string, fn SqliteFN) {
	store := sqlite.New(sqlite.Config{
		Address: conf.Get().Store.Sqlite.Address,
	})
	store.Connect()
	schema := os.Getenv("SQLITE_SCHEMA")
	b, err := os.ReadFile(schema)
	if err != nil {
		log.Fatalf("[ERROR] Cannot read schema file %s for Sqlite: %s", schema, err)
	}
	if err := store.ApplySchemaAndDropData(string(b)); err != nil {
		log.Fatalf("[ERROR] Cannot apply schema. %s", err)
	}
	ctx := context.Background()
	ts.Run(name, func(t *testing.T) {
		fn(t, store, ctx)
	})
}
