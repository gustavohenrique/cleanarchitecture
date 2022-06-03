package controllers

import (
	"context"

	"{{ .ProjectName }}/src/adapters"
	"{{ .ProjectName }}/src/interfaces"
	pb "{{ .ProjectName }}/src/proto"
	"{{ .ProjectName }}/src/services"
)

type TodoWebController struct {
	pb.UnimplementedTodoRpcServer
	todoService     interfaces.ITodoService
	todoItemAdapter adapters.TodoItemAdapter
	searchAdapter   adapters.SearchAdapter
}

func NewTodoWebController(serviceContainer services.ServiceContainer) pb.TodoRpcServer {
	return &TodoWebController{
		todoService:     serviceContainer.GetTodoService(),
		todoItemAdapter: adapters.NewTodoItemAdapter(),
		searchAdapter:   adapters.NewSearchAdapter(),
	}
}

func (s *TodoWebController) Search(ctx context.Context, req *pb.SearchRequest) (*pb.SearchResponse, error) {
	items, err := s.todoService.ReadAll(ctx)
	if err != nil {
		return nil, err
	}
	resp := &pb.SearchResponse{}
	resp.TodoItems = s.todoItemAdapter.FromEntityListToProtoList(items)
	return resp, nil
}

func (s *TodoWebController) Busca(ctx context.Context, req *pb.TodoItem) (*pb.TodoItem, error) {
	return &pb.TodoItem{}, nil
}

func (s *TodoWebController) Create(ctx context.Context, req *pb.TodoItem) (*pb.TodoItem, error) {
	entity := s.todoItemAdapter.FromProtoToEntity(req)
	saved, err := s.todoService.Create(ctx, entity)
	if err != nil {
		return nil, err
	}
	return s.todoItemAdapter.FromEntityToProto(saved), nil
}

func (s *TodoWebController) Update(ctx context.Context, req *pb.TodoItem) (*pb.TodoItem, error) {
	entity := s.todoItemAdapter.FromProtoToEntity(req)
	saved, err := s.todoService.Update(ctx, entity)
	if err != nil {
		return nil, err
	}
	return s.todoItemAdapter.FromEntityToProto(saved), nil
}

func (s *TodoWebController) Remove(ctx context.Context, req *pb.TodoItem) (*pb.Nothing, error) {
	entity := s.todoItemAdapter.FromProtoToEntity(req)
	nothing := &pb.Nothing{}
	err := s.todoService.Remove(ctx, entity)
	return nothing, err
}
