package controllers

import (
	"context"

	"myproject/src/adapters"
	"myproject/src/interfaces"
	pb "myproject/src/proto"
	"myproject/src/services"
)

type TodoRpcController struct {
	pb.UnimplementedTodoRpcServer
	todoService     interfaces.ITodoService
	todoItemAdapter adapters.TodoItemAdapter
	searchAdapter   adapters.SearchAdapter
}

func NewTodoRpcController(serviceContainer services.ServiceContainer) pb.TodoRpcServer {
	return &TodoRpcController{
		todoService:     serviceContainer.GetTodoService(),
		todoItemAdapter: adapters.NewTodoItemAdapter(),
		searchAdapter:   adapters.NewSearchAdapter(),
	}
}

func (s *TodoRpcController) Search(ctx context.Context, req *pb.SearchRequest) (*pb.SearchResponse, error) {
	items, err := s.todoService.ReadAll(ctx)
	if err != nil {
		return nil, err
	}
	resp := &pb.SearchResponse{}
	resp.TodoItems = s.todoItemAdapter.FromEntityListToProtoList(items)
	return resp, nil
}

func (s *TodoRpcController) Create(ctx context.Context, req *pb.TodoItem) (*pb.TodoItem, error) {
	entity := s.todoItemAdapter.FromProtoToEntity(req)
	saved, err := s.todoService.Create(ctx, entity)
	if err != nil {
		return nil, err
	}
	return s.todoItemAdapter.FromEntityToProto(saved), nil
}

func (s *TodoRpcController) Update(ctx context.Context, req *pb.TodoItem) (*pb.TodoItem, error) {
	entity := s.todoItemAdapter.FromProtoToEntity(req)
	saved, err := s.todoService.Update(ctx, entity)
	if err != nil {
		return nil, err
	}
	return s.todoItemAdapter.FromEntityToProto(saved), nil
}

func (s *TodoRpcController) Remove(ctx context.Context, req *pb.TodoItem) (*pb.Nothing, error) {
	entity := s.todoItemAdapter.FromProtoToEntity(req)
	nothing := &pb.Nothing{}
	err := s.todoService.Remove(ctx, entity)
	return nothing, err
}
