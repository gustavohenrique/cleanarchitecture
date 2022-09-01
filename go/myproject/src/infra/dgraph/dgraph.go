package dgraph

import (
	"context"

	"github.com/dgraph-io/dgo/v210"
	"github.com/dgraph-io/dgo/v210/protos/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"{{ .ProjectName }}/src/shared/conf"
	"{{ .ProjectName }}/src/shared/customjson"
)

type DgraphStore struct {
	config *conf.Config
	client *dgo.Dgraph
	ctx    context.Context
}

func NewDgraphStore() *DgraphStore {
	return &DgraphStore{
		config: conf.Get(),
		ctx:    context.Background(),
	}
}

func (db *DgraphStore) Connect() error {
	databaseURL := db.config.Store.Dgraph.Grpc
	conn, err := grpc.Dial(databaseURL, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return err
	}
	db.client = dgo.NewDgraphClient(api.NewDgraphClient(conn))
	return nil
}

func (db *DgraphStore) WithContext(ctx context.Context) *DgraphStore {
	db.ctx = ctx
	return db
}

func (db *DgraphStore) Query(dest interface{}, dql string) error {
	txn := db.getConnection().NewReadOnlyTxn()
	resp, err := txn.Query(db.ctx, dql)
	if err != nil {
		return err
	}
	return customjson.FromDgraphJson(resp.GetJson(), dest)
}

func (db *DgraphStore) QueryWithVars(dql string, vars map[string]string) ([]byte, error) {
	txn := db.getConnection().NewReadOnlyTxn()
	resp, err := txn.QueryWithVars(db.ctx, dql, vars)
	if err != nil {
		return []byte(""), err
	}
	return resp.GetJson(), nil
}

func (db *DgraphStore) MutateReturnUids(dql string) ([]string, error) {
	txn := db.getConnection().NewTxn()
	mu := &api.Mutation{
		SetJson:   []byte(dql),
		CommitNow: true,
	}
	resp, err := txn.Mutate(db.ctx, mu)
	uids := []string{}
	if err != nil {
		return uids, err
	}
	for _, v := range resp.Uids {
		uids = append(uids, v)
	}
	return uids, nil
}

func (db *DgraphStore) MutateReturnUid(dql string) (string, error) {
	uids, err := db.MutateReturnUids(dql)
	if err != nil || len(uids) == 0 {
		return "", err
	}
	return uids[0], nil
}

func (db *DgraphStore) Mutate(dql string) error {
	_, err := db.MutateReturnUids(dql)
	return err
}

func (db *DgraphStore) Delete(dql string) error {
	txn := db.getConnection().NewTxn()
	mu := &api.Mutation{
		DeleteJson: []byte(dql),
		CommitNow:  true,
	}
	_, err := txn.Mutate(db.ctx, mu)
	return err
}

func (db *DgraphStore) ApplySchemaAndDropData(schema string) error {
	op := &api.Operation{
		Schema: schema,
		// DropAll: true,
		DropOp: api.Operation_DATA,
	}
	conn := db.getConnection()
	return conn.Alter(db.ctx, op)
}

func (db *DgraphStore) getConnection() *dgo.Dgraph {
	if db.client == nil {
		db.Connect()
	}
	return db.client
}
