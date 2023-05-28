package controllers

import (
	"fmt"

	"{{ .ProjectName }}/src/domain/ports"
	"{{ .ProjectName }}/src/infrastructure/clients/natsclient"
)

{{ range .Models }}
type {{ .CamelCaseName }}NatsController struct {
	{{ .LowerCaseName }}Repository ports.{{ .CamelCaseName }}Repository
	client         natsclient.NatsClient
}

func New{{ .CamelCaseName }}NatsController(repos ports.Repositories, client natsclient.NatsClient) *{{ .CamelCaseName }}NatsController {
	return &{{ .CamelCaseName }}NatsController{
		{{ .LowerCaseName }}Repository: repos.{{ .CamelCaseName }}Repository(),
		client:         client,
	}
}

func (c *{{ .CamelCaseName }}NatsController) SubscribeToNew{{ .CamelCaseName }}() error {
	return c.client.Subscribe("new-{{ .LowerCaseName }}", func(message string) {
		fmt.Println("new-{{ .LowerCaseName }} event:", message)
	})
}
{{ end }}
