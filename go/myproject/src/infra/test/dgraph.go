package test

import (
	"context"
	"log"
	"os"
	"testing"

	"{{ .ProjectName }}/src/infra/conf"
	"{{ .ProjectName }}/src/infra/datastores/dgraph"
	"{{ .ProjectName }}/src/interfaces"
)

type DgraphFN func(t *testing.T, store interfaces.IGraphDataStore, ctx context.Context)

func WithDgraph(ts *testing.T, name string, fn DgraphFN) {
	store := dgraph.New(dgraph.Config{
		Address: conf.Get().Store.Dgraph.Address,
	})
	store.Connect()
	schema := os.Getenv("DGRAPH_SCHEMA")
	b, err := os.ReadFile(schema)
	if err != nil {
		log.Fatalf("Cannot read schema file %s for Dgraph: %s", schema, err)
	}
	if err := store.ApplySchemaAndDropData(string(b)); err != nil {
		log.Fatalf("Cannot apply schema. %s", err)
	}
	ctx := context.Background()
	ts.Run(name, func(t *testing.T) {
		fn(t, store, ctx)
	})
}
