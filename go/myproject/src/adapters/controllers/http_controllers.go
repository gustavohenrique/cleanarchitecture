package controllers

import "{{ .ProjectName }}/src/domain/ports"

type HttpControllers interface {
{{ range .Models }}
	{{ .CamelCaseName }}Controller() *{{ .CamelCaseName }}HttpController
{{ end }}
}

type HttpControllersContainer struct {
	repos ports.Repositories
}

func NewHttpControllers(repos ports.Repositories) HttpControllers {
	return &HttpControllersContainer{repos}
}
{{ range .Models }}
func (c *HttpControllersContainer) {{ .CamelCaseName }}Controller() *{{ .CamelCaseName }}HttpController {
	return New{{ .CamelCaseName }}HttpController(c.repos)
}
{{ end }}
