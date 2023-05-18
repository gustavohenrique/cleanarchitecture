package controllers

import (
	"context"

	"{{ .ProjectName }}/pb"
	"{{ .ProjectName }}/src/domain/ports"
)

type TodoGrpcWebController struct {
	pb.UnimplementedTodoRpcServer
	todoRepository ports.TodoRepository
}

func NewTodoGrpcWebController(repos ports.Repositories) pb.TodoRpcServer {
	return &TodoGrpcWebController{
		todoRepository: repos.TodoRepository(),
	}
}

func (s *TodoGrpcWebController) Search(ctx context.Context, req *pb.SearchRequest) (*pb.SearchResponse, error) {
	resp := &pb.SearchResponse{}
	return resp, nil
}

func (s *TodoGrpcWebController) Create(ctx context.Context, req *pb.TodoItem) (*pb.TodoItem, error) {
	return req, nil
}

func (s *TodoGrpcWebController) Update(ctx context.Context, req *pb.TodoItem) (*pb.TodoItem, error) {
	return req, nil
}

func (s *TodoGrpcWebController) Remove(ctx context.Context, req *pb.TodoItem) (*pb.Nothing, error) {
	return nil, nil
}
