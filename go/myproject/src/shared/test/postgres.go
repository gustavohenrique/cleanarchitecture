package test

import (
	"context"
	"io/ioutil"
	"log"
	"os"
	"testing"

	"{{ .ProjectName }}/src/infra/postgres"
	"{{ .ProjectName }}/src/shared/conf"
)

type PostgresFN func(t *testing.T, store *postgres.PostgresStore, ctx context.Context)

func WithPostgres(ts *testing.T, name string, fn PostgresFN) {
	store := postgres.NewPostgresStore()
	store.Connect()
	key := "POSTGRES_SCHEMA_FILE"
	file := os.Getenv(key)
	if file == "" {
		file = conf.Get().Store.Postgres.Schema
	}
	b, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatalf("[ERROR] Cannot read schema file %s for Postgres: %s", file, err)
	}
	if err := store.ApplySchemaAndDropData(string(b)); err != nil {
		log.Fatalf("[ERROR] Cannot apply schema. %s", err)
	}
	ctx := context.WithValue(context.Background(), DB, "postgres")
	ts.Run(name, func(t *testing.T) {
		fn(t, store, ctx)
	})
}
