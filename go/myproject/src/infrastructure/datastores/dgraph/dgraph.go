package dgraph

import (
	"context"

	"github.com/dgraph-io/dgo/v210"
	"github.com/dgraph-io/dgo/v210/protos/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"{{ .ProjectName }}/src/infrastructure/datastores/db"
	"{{ .ProjectName }}/src/infrastructure/datastores/dgraph/customjson"
)

type Config struct {
	Address string
}

type DgraphStore struct {
	config Config
	client *dgo.Dgraph
	ctx    context.Context
}

func New(config Config) db.GraphDataStore {
	return &DgraphStore{
		config: config,
		ctx:    context.Background(),
	}
}

func (store *DgraphStore) Connect() error {
	databaseURL := store.config.Address
	conn, err := grpc.Dial(databaseURL, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return err
	}
	store.client = dgo.NewDgraphClient(api.NewDgraphClient(conn))
	return nil
}

func (store *DgraphStore) WithContext(ctx context.Context) db.GraphDataStore {
	store.ctx = ctx
	return store
}

func (store *DgraphStore) Query(dest interface{}, dql string) error {
	txn := store.getConnection().NewReadOnlyTxn()
	resp, err := txn.Query(store.ctx, dql)
	if err != nil {
		return err
	}
	return customjson.FromDgraphJson(resp.GetJson(), dest)
}

func (store *DgraphStore) QueryWithVars(dql string, vars map[string]string) ([]byte, error) {
	txn := store.getConnection().NewReadOnlyTxn()
	resp, err := txn.QueryWithVars(store.ctx, dql, vars)
	if err != nil {
		return []byte(""), err
	}
	return resp.GetJson(), nil
}

func (store *DgraphStore) MutateReturnUids(dql string) ([]string, error) {
	txn := store.getConnection().NewTxn()
	mu := &api.Mutation{
		SetJson:   []byte(dql),
		CommitNow: true,
	}
	resp, err := txn.Mutate(store.ctx, mu)
	uids := []string{}
	if err != nil {
		return uids, err
	}
	for _, v := range resp.Uids {
		uids = append(uids, v)
	}
	return uids, nil
}

func (store *DgraphStore) MutateReturnUid(dql string) (string, error) {
	uids, err := store.MutateReturnUids(dql)
	if err != nil || len(uids) == 0 {
		return "", err
	}
	return uids[0], nil
}

func (store *DgraphStore) Mutate(dql string) error {
	_, err := store.MutateReturnUids(dql)
	return err
}

func (store *DgraphStore) Delete(dql string) error {
	txn := store.getConnection().NewTxn()
	mu := &api.Mutation{
		DeleteJson: []byte(dql),
		CommitNow:  true,
	}
	_, err := txn.Mutate(store.ctx, mu)
	return err
}

func (store *DgraphStore) ApplySchemaAndDropData(schema string) error {
	op := &api.Operation{
		Schema: schema,
		// DropAll: true,
		DropOp: api.Operation_DATA,
	}
	conn := store.getConnection()
	return conn.Alter(store.ctx, op)
}

func (store *DgraphStore) getConnection() *dgo.Dgraph {
	if store.client == nil {
		store.Connect()
	}
	return store.client
}
