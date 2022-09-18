package postgres_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"{{ .ProjectName }}/src/infra/conf"
	"{{ .ProjectName }}/src/infra/datastores/postgres"
)

var ctx = context.Background()

func TestBulkCopy(t *testing.T) {
	start := time.Now()
	columns := []string{"id", "title"}
	rows := getRows(10)
	pconf := conf.Get().Store.Postgres
	config := postgres.Config{
		URL:             pconf.URL,
		MaxOpenConns:    pconf.MaxOpenConns,
		MaxIdleConns:    pconf.MaxIdleConns,
		MaxConnLifetime: pconf.MaxConnLifetime,
	}
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
