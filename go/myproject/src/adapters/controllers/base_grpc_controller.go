package controllers

import (
	"context"

	"{{ .ProjectName }}/pb"
	"{{ .ProjectName }}/src/domain/ports"
)

{{ range .Models }}
type {{ .CamelCaseName }}GrpcController struct {
	pb.Unimplemented{{ .CamelCaseName }}RpcServer
	{{ .LowerCaseName }}Repository ports.{{ .CamelCaseName }}Repository
}

func New{{ .CamelCaseName }}GrpcController(repos ports.Repositories) pb.{{ .CamelCaseName }}RpcServer {
	return &{{ .Name }}GrpcController{
		{{ .LowerCaseName }}Repository: repos.{{ .CamelCaseName }}Repository(),
	}
}

func (s *{{ .CamelCaseName }}GrpcController) Create(ctx context.Context, req *pb.{{ .CamelCaseName }}) (*pb.{{ .CamelCaseName }}, error) {
	return req, nil
}

func (s *{{ .CamelCaseName }}GrpcController) Update(ctx context.Context, req *pb.{{ .CamelCaseName }}) (*pb.{{ .CamelCaseName }}, error) {
	return req, nil
}

func (s *{{ .CamelCaseName }}GrpcController) Remove(ctx context.Context, req *pb.{{ .CamelCaseName }}) (*pb.Nothing, error) {
	return nil, nil
}
{{ end }}
