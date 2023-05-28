package controllers

import (
	"{{ .ProjectName }}/pb"
	"{{ .ProjectName }}/src/domain/ports"
)

type GrpcControllers interface {
{{ range .Models }}
	{{ .CamelCaseName }}Controller() pb.{{ .CamelCaseName }}RpcServer
{{ end }}
}

type GrpcControllersContainer struct {
	repos ports.Repositories
}

func NewGrpcControllers(repos ports.Repositories) GrpcControllers {
	return &GrpcControllersContainer{repos}
}
{{ range .Models }}
func (c *GrpcControllersContainer) {{ .CamelCaseName }}Controller() pb.{{ .CamelCaseName }}RpcServer {
	return New{{ .CamelCaseName }}GrpcController(c.repos)
}
{{ end }}
