package controllers

import "{{ .ProjectName }}/src/domain/ports"

type HttpControllers interface {
	TodoController() *TodoHttpController
}

type HttpControllersContainer struct {
	repos ports.Repositories
}

func NewHttpControllers(repos ports.Repositories) HttpControllers {
	return &HttpControllersContainer{repos}
}

func (c *HttpControllersContainer) TodoController() *TodoHttpController {
	return NewTodoHttpController(c.repos)
}
