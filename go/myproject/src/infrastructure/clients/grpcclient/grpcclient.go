package grpcclient

import (
	"context"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"{{ .ProjectName }}/pb"
)

type Config struct {
	Address string
	Timeout time.Duration
}

type grpcClient struct {
	ctx    context.Context
	config Config
	conn   *grpc.ClientConn
	{{ range .Models }}
	{{ .LowerCaseName }}Client pb.{{ .CamelCaseName }}RpcClient
	{{ end }}
}

func New(config Config) GrpcClient {
	ctx, _ := context.WithTimeout(context.Background(), config.Timeout*time.Second)
	return &grpcClient{
		ctx:    ctx,
		config: config,
	}
}
{{ range .Models }}
func (c *grpcClient) {{ .CamelCaseName}}Client() pb.{{ .CamelCaseName }}RpcClient {
	if c.{{ .LowerCaseName }}Client != nil {
		return c.{{ .LowerCaseName }}Client
	}
	conn, err := c.getConnection()
	if err != nil {
		return nil
	}
	client := pb.New{{ .CamelCaseName }}RpcClient(conn)
	c.{{ .LowerCaseName }}Client = client
	return client
}
{{ end }}

func (c *grpcClient) getConnection() (*grpc.ClientConn, error) {
	if c.conn != nil {
		return c.conn, nil
	}
	opts := grpc.WithTransportCredentials(insecure.NewCredentials())
	conn, err := grpc.DialContext(c.ctx, c.config.Address, opts)
	if err != nil {
		return nil, err
	}
	c.conn = conn
	return conn, nil
}
