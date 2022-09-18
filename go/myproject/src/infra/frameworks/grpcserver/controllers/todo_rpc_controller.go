package controllers

import (
	"context"

	"{{ .ProjectName }}/src/interfaces"
	pb "{{ .ProjectName }}/src/proto"
)

type TodoRpcController struct {
	pb.UnimplementedTodoRpcServer
	todoService interfaces.ITodoService
	adapter     *Adapter
}

func NewTodoRpcController(services interfaces.IService) pb.TodoRpcServer {
	return &TodoRpcController{
		todoService: services.GetTodoService(),
		adapter:     NewAdapter(),
	}
}

func (s *TodoRpcController) Search(ctx context.Context, req *pb.SearchRequest) (*pb.SearchResponse, error) {
	items, err := s.todoService.ReadAll(ctx)
	if err != nil {
		return nil, err
	}
	resp := &pb.SearchResponse{}
	resp.TodoItems = s.adapter.ToProtos(items)
	return resp, nil
}

func (s *TodoRpcController) Create(ctx context.Context, req *pb.TodoItem) (*pb.TodoItem, error) {
	entity := s.adapter.ToEntity(req)
	saved, err := s.todoService.Create(ctx, entity)
	if err != nil {
		return nil, err
	}
	return s.adapter.ToProto(saved), nil
}

func (s *TodoRpcController) Update(ctx context.Context, req *pb.TodoItem) (*pb.TodoItem, error) {
	entity := s.adapter.ToEntity(req)
	saved, err := s.todoService.Update(ctx, entity)
	if err != nil {
		return nil, err
	}
	return s.adapter.ToProto(saved), nil
}

func (s *TodoRpcController) Remove(ctx context.Context, req *pb.TodoItem) (*pb.Nothing, error) {
	entity := s.adapter.ToEntity(req)
	nothing := &pb.Nothing{}
	err := s.todoService.Remove(ctx, entity)
	return nothing, err
}
