package grpcclient

import "{{ .ProjectName }}/pb"

type GrpcClient interface {
	{{ range .Models }}
	{{ .CamelCaseName }}Client() pb.{{ .CamelCaseName }}RpcClient
	{{ end }}
}
