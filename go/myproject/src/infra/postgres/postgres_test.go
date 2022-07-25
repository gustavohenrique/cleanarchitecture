package postgres_test

import (
	"fmt"
	"testing"

	"{{ .ProjectName }}/src/infra/postgres"
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
	store := postgres.NewPostgresStore()
	err := store.BulkInsert("todo_items", columns, rows)
	if err != nil {
		t.Errorf("Error: %s", err)
	}
}
