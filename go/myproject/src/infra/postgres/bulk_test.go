package postgres_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"{{ .ProjectName }}/src/infra/postgres"
	"{{ .ProjectName }}/src/shared/conf"
)

var (
	ctx    = context.Background()
	config = conf.Get()
)

func TestBulkInsert(t *testing.T) {
	start := time.Now()
	columns := []string{"id", "title"}
	rows := getRows(10)
	err := postgres.Bulk(ctx, config).Copy("todo_items", columns, rows)
	elapsed := time.Since(start)
	if err != nil {
		t.Errorf("Time elapsed: %fs. Error: %s", elapsed.Seconds(), err)
	}
	fmt.Printf("Time elapsed: %fs. Total rows: %d", elapsed.Seconds(), len(rows))
}

func getRows(total int) [][]interface{} {
	var rows [][]interface{}
	for i := 0; i < total; i++ {
		record := []interface{}{
			fmt.Sprintf("ID %d", i),
			fmt.Sprintf("Title %d", i),
		}
		rows = append(rows, record)
	}
	return rows
}
