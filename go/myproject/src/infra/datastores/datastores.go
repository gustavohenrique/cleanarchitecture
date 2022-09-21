package datastores

import (
	"{{ .ProjectName }}/src/infra/conf"
	"{{ .ProjectName }}/src/infra/datastores/dgraph"
	"{{ .ProjectName }}/src/infra/datastores/postgres"
	"{{ .ProjectName }}/src/infra/datastores/sqlite"
	"{{ .ProjectName }}/src/interfaces"
)

type DataStore struct {
	config        *conf.Config
	postgresStore interfaces.ISqlDataStore
	sqliteStore   interfaces.ISqlDataStore
	dgraphStore   interfaces.IGraphDataStore
}

func With(config *conf.Config) interfaces.IDataStore {
	return DataStore{config: config}
}

func (ds DataStore) New() interfaces.IDataStore {
	pconf := ds.config.Store.Postgres
	ds.postgresStore = postgres.New(postgres.Config{
		URL:             pconf.URL,
		MaxOpenConns:    pconf.MaxOpenConns,
		MaxIdleConns:    pconf.MaxIdleConns,
		MaxConnLifetime: pconf.MaxConnLifetime,
	})

	ds.sqliteStore = sqlite.New(sqlite.Config{
		Address: ds.config.Store.Sqlite.Address,
	})
	ds.dgraphStore = dgraph.New(dgraph.Config{
		Address: ds.config.Store.Dgraph.Address,
	})
	return ds
}

func (ds DataStore) SQL() interfaces.ISqlDataStore {
	return ds.Postgres()
}

func (ds DataStore) Postgres() interfaces.ISqlDataStore {
	return ds.postgresStore
}

func (ds DataStore) Sqlite() interfaces.ISqlDataStore {
	return ds.sqliteStore
}

func (ds DataStore) Dgraph() interfaces.IGraphDataStore {
	return ds.dgraphStore
}
