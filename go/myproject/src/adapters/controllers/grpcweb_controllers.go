package controllers

import (
	"{{ .ProjectName }}/pb"
	"{{ .ProjectName }}/src/domain/ports"
)

type GrpcWebControllers interface {
{{ range .Models }}
	{{ .CamelCaseName }}Controller() pb.{{ .CamelCaseName }}RpcServer
{{ end }}
}

type GrpcWebControllersContainer struct {
	repos ports.Repositories
}

func NewGrpcWebControllers(repos ports.Repositories) GrpcWebControllers {
	return &GrpcWebControllersContainer{repos}
}
{{ range .Models }}
func (c *GrpcWebControllersContainer) {{ .CamelCaseName }}Controller() pb.{{ .CamelCaseName }}RpcServer {
	return New{{ .CamelCaseName }}GrpcWebController(c.repos)
}
{{ end }}
