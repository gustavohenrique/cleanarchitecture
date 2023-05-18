package controllers

import (
	"{{ .ProjectName }}/pb"
	"{{ .ProjectName }}/src/domain/ports"
)

type GrpcWebControllers interface {
	TodoController() pb.TodoRpcServer
}

type GrpcWebControllersContainer struct {
	repos ports.Repositories
}

func NewGrpcWebControllers(repos ports.Repositories) GrpcWebControllers {
	return &GrpcWebControllersContainer{repos}
}

func (c *GrpcWebControllersContainer) TodoController() pb.TodoRpcServer {
	return NewTodoGrpcWebController(c.repos)
}
