package gateways

import (
	"{{ .ProjectName }}/src/domain/gateways/todo"
	"{{ .ProjectName }}/src/infra/conf"
	"{{ .ProjectName }}/src/interfaces"
)

type Gateway struct {
	config      *conf.Config
	todoGateway interfaces.ITodoGateway
}

func With(config *conf.Config) interfaces.IGateway {
	return &Gateway{
		config: config,
	}
}

func (g *Gateway) Inject(store interfaces.IDataStore) interfaces.IGateway {
	g.todoGateway = todo.New(store)
	return g
}

func (g *Gateway) TodoGateway() interfaces.ITodoGateway {
	return g.todoGateway
}
