package controllers

import (
	"context"

	"{{ .ProjectName }}/pb"
	"{{ .ProjectName }}/src/domain/ports"
)

type TodoGrpcController struct {
	pb.UnimplementedTodoRpcServer
	todoRepository ports.TodoRepository
}

func NewTodoGrpcController(repos ports.Repositories) pb.TodoRpcServer {
	return &TodoGrpcController{
		todoRepository: repos.TodoRepository(),
	}
}

func (s *TodoGrpcController) Search(ctx context.Context, req *pb.SearchRequest) (*pb.SearchResponse, error) {
	resp := &pb.SearchResponse{}
	return resp, nil
}

func (s *TodoGrpcController) Create(ctx context.Context, req *pb.TodoItem) (*pb.TodoItem, error) {
	return req, nil
}

func (s *TodoGrpcController) Update(ctx context.Context, req *pb.TodoItem) (*pb.TodoItem, error) {
	return req, nil
}

func (s *TodoGrpcController) Remove(ctx context.Context, req *pb.TodoItem) (*pb.Nothing, error) {
	return nil, nil
}
