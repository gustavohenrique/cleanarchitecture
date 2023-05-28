package controllers

import (
	"context"

	"{{ .ProjectName }}/pb"
	"{{ .ProjectName }}/src/domain/ports"
)

{{ range .Models }}
type {{ .CamelCaseName }}GrpcWebController struct {
	pb.Unimplemented{{ .CamelCaseName }}RpcServer
	{{ .LowerCaseName }}Repository ports.{{ .CamelCaseName }}Repository
}

func New{{ .CamelCaseName }}GrpcWebController(repos ports.Repositories) pb.{{ .CamelCaseName }}RpcServer {
	return &{{ .CamelCaseName }}GrpcWebController{
		{{ .LowerCaseName }}Repository: repos.{{ .CamelCaseName }}Repository(),
	}
}

func (s *{{ .CamelCaseName }}GrpcWebController) Search(ctx context.Context, req *pb.SearchRequest) (*pb.SearchResponse, error) {
	resp := &pb.SearchResponse{}
	return resp, nil
}

func (s *{{ .CamelCaseName }}GrpcWebController) Create(ctx context.Context, req *pb.{{ .CamelCaseName }}) (*pb.{{ .CamelCaseName }}, error) {
	return req, nil
}

func (s *{{ .CamelCaseName }}GrpcWebController) Update(ctx context.Context, req *pb.{{ .CamelCaseName }}) (*pb.{{ .CamelCaseName }}, error) {
	return req, nil
}

func (s *{{ .CamelCaseName }}GrpcWebController) Remove(ctx context.Context, req *pb.{{ .CamelCaseName }}) (*pb.Nothing, error) {
	return nil, nil
}
{{ end }}
