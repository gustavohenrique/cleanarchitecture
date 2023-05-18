package controllers

import (
	"fmt"

	"{{ .ProjectName }}/src/domain/ports"
	"{{ .ProjectName }}/src/infrastructure/clients/natsclient"
)

type TodoNatsController struct {
	todoRepository ports.TodoRepository
	client         natsclient.NatsClient
}

func NewTodoNatsController(repos ports.Repositories, client natsclient.NatsClient) *TodoNatsController {
	return &TodoNatsController{
		todoRepository: repos.TodoRepository(),
		client:         client,
	}
}

func (c *TodoNatsController) SubscribeToNewTodo() error {
	return c.client.Subscribe("new-todo", func(message string) {
		fmt.Println("new-todo event:", message)
	})
}
