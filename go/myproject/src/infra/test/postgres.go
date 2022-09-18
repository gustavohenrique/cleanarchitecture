package test

import (
	"context"
	"log"
	"os"
	"testing"

	"{{ .ProjectName }}/src/infra/conf"
	"{{ .ProjectName }}/src/infra/datastores/postgres"
	"{{ .ProjectName }}/src/interfaces"
)

type PostgresFN func(t *testing.T, store interfaces.ISqlDataStore, ctx context.Context)

func WithPostgres(ts *testing.T, name string, fn PostgresFN) {
	pconf := conf.Get().Store.Postgres
	store := postgres.New(postgres.Config{
		URL:             pconf.URL,
		MaxOpenConns:    pconf.MaxOpenConns,
		MaxIdleConns:    pconf.MaxIdleConns,
		MaxConnLifetime: pconf.MaxConnLifetime,
	})
	store.Connect()
	schema := os.Getenv("POSTGRES_SCHEMA")
	b, err := os.ReadFile(schema)
	if err != nil {
		log.Fatalf("[ERROR] Cannot read schema file %s for Postgres: %s", schema, err)
	}
	if err := store.ApplySchemaAndDropData(string(b)); err != nil {
		log.Fatalf("[ERROR] Cannot apply schema. %s", err)
	}
	ctx := context.Background()
	ts.Run(name, func(t *testing.T) {
		fn(t, store, ctx)
	})
}
