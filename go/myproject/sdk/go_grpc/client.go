package gogrpc

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	pb "myproject/sdk/gogrpc/proto"
)

type Client struct {
	config     *Config
	todoClient pb.TodoRpcClient
}

func NewClient(conn Connection, config *Config) *Client {
	return &Client{
		config:     config,
		todoClient: pb.NewTodoRpcClient(conn.(grpc.ClientConnInterface)),
	}
}

func (c *Client) Auth(ctx context.Context) context.Context {
	md := metadata.Pairs("X-CSRF-Token", c.config.Token, "X-User-Id", "some-user-id")
	return metadata.NewOutgoingContext(ctx, md)
}

func (c *Client) GetTodoClient() pb.TodoRpcClient {
	return c.todoClient
}
