package datastores

import (
	"{{ .ProjectName }}/src/components/configurator"
	"{{ .ProjectName }}/src/infrastructure/datastores/db"
	"{{ .ProjectName }}/src/infrastructure/datastores/dgraph"
	"{{ .ProjectName }}/src/infrastructure/datastores/postgres"
	"{{ .ProjectName }}/src/infrastructure/datastores/sqlite"
)

type Stores interface {
{{ if .HasPostgres }}
	Postgres() db.SqlDataStore
{{ end }}
{{ if .HasSqlite }}
	Sqlite() db.SqlDataStore
{{ end }}
{{ if .HasDgraph }}
	Dgraph() db.GraphDataStore
{{ end }}
}

type DataStoresContainer struct {
{{ if .HasPostgres }}
	postgres db.SqlDataStore
{{ end }}
{{ if .HasSqlite }}
	sqlite   db.SqlDataStore
{{ end }}
{{ if .HasDgraph }}
	dgraph   db.GraphDataStore
{{ end }}
}

func New(config *configurator.Config) Stores {
	datastores := &DataStoresContainer{}
{{ if .HasPostgres }}
	pconf := config.Store.Postgres
	datastores.postgres = postgres.New(postgres.Config{
		URL:             pconf.URL,
		MaxOpenConns:    pconf.MaxOpenConns,
		MaxIdleConns:    pconf.MaxIdleConns,
		MaxConnLifetime: pconf.MaxConnLifetime,
	})
{{ end }}
{{ if .HasSqlite }}
	datastores.sqlite = sqlite.New(sqlite.Config{
		Address: config.Store.Sqlite.Address,
	})
{{ end }}
{{ if .HasDgraph }}
	datastores.dgraph = dgraph.New(dgraph.Config{
		Address: config.Store.Dgraph.Address,
	})
{{ end }}
	return datastores
}
{{ if .HasPostgres }}
func (d *DataStoresContainer) Postgres() db.SqlDataStore {
	return d.postgres
}
{{ end }}
{{ if .HasSqlite }}
func (d *DataStoresContainer) Sqlite() db.SqlDataStore {
	return d.sqlite
}
{{ end }}
{{ if .HasDgraph }}
func (d *DataStoresContainer) Dgraph() db.GraphDataStore {
	return d.dgraph
}
{{ end }}
