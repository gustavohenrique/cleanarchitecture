package controllers

import (
	"{{ .ProjectName }}/pb"
	"{{ .ProjectName }}/src/domain/ports"
)

type GrpcControllers interface {
	TodoController() pb.TodoRpcServer
}

type GrpcControllersContainer struct {
	repos ports.Repositories
}

func NewGrpcControllers(repos ports.Repositories) GrpcControllers {
	return &GrpcControllersContainer{repos}
}

func (c *GrpcControllersContainer) TodoController() pb.TodoRpcServer {
	return NewTodoGrpcController(c.repos)
}
