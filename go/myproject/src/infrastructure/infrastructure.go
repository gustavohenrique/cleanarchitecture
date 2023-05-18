package infrastructure

import (
	"{{ .ProjectName }}/src/adapters/controllers"
	"{{ .ProjectName }}/src/components/configurator"
	"{{ .ProjectName }}/src/infrastructure/datastores"
	"{{ .ProjectName }}/src/infrastructure/servers"
	"{{ .ProjectName }}/src/infrastructure/servers/grpcserver"
	"{{ .ProjectName }}/src/infrastructure/servers/grpcwebserver"
	"{{ .ProjectName }}/src/infrastructure/servers/httpserver"
	"{{ .ProjectName }}/src/infrastructure/servers/natsserver"
)

type Infra interface {
	DataStores() datastores.Stores
	HttpServer(handlers controllers.HttpControllers) servers.Server
	GrpcServer(handlers controllers.GrpcControllers) servers.Server
	GrpcWebServer(handlers controllers.GrpcWebControllers) servers.Server
	NatsServer(handlers controllers.NatsControllers) servers.Server
}

type InfraContainer struct {
	config *configurator.Config
}

func New(config *configurator.Config) Infra {
	return InfraContainer{config: config}
}

func (i InfraContainer) HttpServer(controllers controllers.HttpControllers) servers.Server {
	return httpserver.New(i.config, controllers)
}

func (i InfraContainer) GrpcServer(controllers controllers.GrpcControllers) servers.Server {
	return grpcserver.New(i.config, controllers)
}

func (i InfraContainer) GrpcWebServer(controllers controllers.GrpcWebControllers) servers.Server {
	return grpcwebserver.New(i.config, controllers)
}

func (i InfraContainer) NatsServer(controllers controllers.NatsControllers) servers.Server {
	return natsserver.New(i.config, controllers)
}

func (i InfraContainer) DataStores() datastores.Stores {
	return datastores.New(i.config)
}
