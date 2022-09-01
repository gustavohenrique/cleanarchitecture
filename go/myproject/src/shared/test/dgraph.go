package test

import (
	"context"
	"io/ioutil"
	"log"
	"os"
	"testing"

	"{{ .ProjectName }}/src/infra/dgraph"
)

type DgraphFN func(t *testing.T, store *dgraph.DgraphStore, ctx context.Context)

func WithDgraph(ts *testing.T, name string, fn DgraphFN) {
	store := dgraph.NewDgraphStore()
	store.Connect()
	file := os.Getenv("DGRAPH_MIGRATION_FILE")
	if file == "" {
		log.Fatalln("Please set DGRAPH_MIGRATION_FILE variable.")
	}
	b, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatalf("Cannot read schema file %s for Dgraph: %s", file, err)
	}
	if err := store.ApplySchemaAndDropData(string(b)); err != nil {
		log.Fatalf("Cannot apply schema. %s", err)
	}
	ctx := context.WithValue(context.Background(), DB, "dgraph")
	ts.Run(name, func(t *testing.T) {
		fn(t, store, ctx)
	})
}
