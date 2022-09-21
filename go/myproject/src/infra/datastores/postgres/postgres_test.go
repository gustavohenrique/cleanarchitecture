package postgres_test

import (
	"fmt"
	"testing"

	"{{ .ProjectName }}/src/infra/conf"
	"{{ .ProjectName }}/src/infra/datastores/postgres"
)

func TestBulkInsert(t *testing.T) {
	columns := []string{"id", "title"}
	var rows [][]interface{}
	for i := 0; i < 10; i++ {
		record := []interface{}{
			fmt.Sprintf("Todo ID %d", i),
			fmt.Sprintf("Todo Title %d", i),
		}
		rows = append(rows, record)
	}
	pconf := conf.Get().Store.Postgres
	store := postgres.New(postgres.Config{
		URL:             pconf.URL,
		MaxOpenConns:    pconf.MaxOpenConns,
		MaxIdleConns:    pconf.MaxIdleConns,
		MaxConnLifetime: pconf.MaxConnLifetime,
	})
	err := store.BulkInsert("todo_items", columns, rows)
	if err != nil {
		t.Errorf("Error: %s", err)
	}
}
