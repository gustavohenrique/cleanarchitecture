package controllers

import (
	"{{ .ProjectName }}/src/domain/ports"
	"{{ .ProjectName }}/src/infrastructure/clients/natsclient"
)

type NatsControllers interface {
	With(client natsclient.NatsClient) NatsControllers
	TodoController() *TodoNatsController
}

type NatsControllersContainer struct {
	repos  ports.Repositories
	client natsclient.NatsClient
}

func NewNatsControllers(repos ports.Repositories) NatsControllers {
	return &NatsControllersContainer{repos: repos}
}

func (c *NatsControllersContainer) With(client natsclient.NatsClient) NatsControllers {
	c.client = client
	return c
}

func (c *NatsControllersContainer) TodoController() *TodoNatsController {
	return NewTodoNatsController(c.repos, c.client)
}
