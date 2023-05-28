package controllers

import (
	"{{ .ProjectName }}/src/domain/ports"
	"{{ .ProjectName }}/src/infrastructure/clients/natsclient"
)

type NatsControllers interface {
	With(client natsclient.NatsClient) NatsControllers
{{ range .Models }}
	{{ .CamelCaseName }}Controller() *{{ .CamelCaseName }}NatsController
{{ end }}
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
{{ range .Models }}
func (c *NatsControllersContainer) {{ .CamelCaseName }}Controller() *{{ .CamelCaseName }}NatsController {
	return New{{ .CamelCaseName }}NatsController(c.repos, c.client)
}
{{ end }}
