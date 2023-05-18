package datastores

import (
	"{{ .ProjectName }}/src/components/configurator"
	"{{ .ProjectName }}/src/infrastructure/datastores/db"
	"{{ .ProjectName }}/src/infrastructure/datastores/dgraph"
	"{{ .ProjectName }}/src/infrastructure/datastores/postgres"
	"{{ .ProjectName }}/src/infrastructure/datastores/sqlite"
)

type Stores interface {
	Postgres() db.SqlDataStore
	Sqlite() db.SqlDataStore
	Dgraph() db.GraphDataStore
}

type DataStoresContainer struct {
	postgres db.SqlDataStore
	sqlite   db.SqlDataStore
	dgraph   db.GraphDataStore
}

func New(config *configurator.Config) Stores {
	datastores := &DataStoresContainer{}
	pconf := config.Store.Postgres
	datastores.postgres = postgres.New(postgres.Config{
		URL:             pconf.URL,
		MaxOpenConns:    pconf.MaxOpenConns,
		MaxIdleConns:    pconf.MaxIdleConns,
		MaxConnLifetime: pconf.MaxConnLifetime,
	})

	datastores.sqlite = sqlite.New(sqlite.Config{
		Address: config.Store.Sqlite.Address,
	})

	datastores.dgraph = dgraph.New(dgraph.Config{
		Address: config.Store.Dgraph.Address,
	})
	return datastores
}

func (d *DataStoresContainer) Postgres() db.SqlDataStore {
	return d.postgres
}

func (d *DataStoresContainer) Sqlite() db.SqlDataStore {
	return d.sqlite
}

func (d *DataStoresContainer) Dgraph() db.GraphDataStore {
	return d.dgraph
}
