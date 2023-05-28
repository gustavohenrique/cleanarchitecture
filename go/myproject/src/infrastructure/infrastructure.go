package infrastructure

import (
	"{{ .ProjectName }}/src/adapters/controllers"
	"{{ .ProjectName }}/src/components/configurator"
	"{{ .ProjectName }}/src/infrastructure/datastores"
	"{{ .ProjectName }}/src/infrastructure/servers"
{{ if .HasGrpcServer }}
	"{{ .ProjectName }}/src/infrastructure/servers/grpcserver"
{{ end }}
{{ if .HasGrpcWebServer }}
	"{{ .ProjectName }}/src/infrastructure/servers/grpcwebserver"
{{ end }}
{{ if .HasHttpServer }}
	"{{ .ProjectName }}/src/infrastructure/servers/httpserver"
{{ end }}
{{ if .HasNatsServer }}
	"{{ .ProjectName }}/src/infrastructure/servers/natsserver"
{{ end }}
)

type Infra interface {
	DataStores() datastores.Stores
{{ if .HasHttpServer }}
	HttpServer(handlers controllers.HttpControllers) servers.Server
{{ end }}
{{ if .HasGrpcServer }}
	GrpcServer(handlers controllers.GrpcControllers) servers.Server
{{ end }}
{{ if .HasGrpcWebServer }}
	GrpcWebServer(handlers controllers.GrpcWebControllers) servers.Server
{{ end }}
{{ if .HasNatsServer }}
	NatsServer(handlers controllers.NatsControllers) servers.Server
{{ end }}
}

type InfraContainer struct {
	config *configurator.Config
}

func New(config *configurator.Config) Infra {
	return InfraContainer{config: config}
}
{{ if .HasHttpServer }}
func (i InfraContainer) HttpServer(controllers controllers.HttpControllers) servers.Server {
	return httpserver.New(i.config, controllers)
}
{{ end }}
{{ if .HasGrpcServer }}
func (i InfraContainer) GrpcServer(controllers controllers.GrpcControllers) servers.Server {
	return grpcserver.New(i.config, controllers)
}
{{ end }}
{{ if .HasGrpcWebServer }}
func (i InfraContainer) GrpcWebServer(controllers controllers.GrpcWebControllers) servers.Server {
	return grpcwebserver.New(i.config, controllers)
}
{{ end }}
{{ if .HasNatsServer }}
func (i InfraContainer) NatsServer(controllers controllers.NatsControllers) servers.Server {
	return natsserver.New(i.config, controllers)
}
{{ end }}
func (i InfraContainer) DataStores() datastores.Stores {
	return datastores.New(i.config)
}
