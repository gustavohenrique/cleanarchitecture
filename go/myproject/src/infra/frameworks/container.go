package frameworks

import (
	"log"

	"{{ .ProjectName }}/src/infra/conf"
	"{{ .ProjectName }}/src/infra/frameworks/grpcserver"
	"{{ .ProjectName }}/src/infra/frameworks/grpcwebserver"
	"{{ .ProjectName }}/src/infra/frameworks/httpserver"
	"{{ .ProjectName }}/src/infra/frameworks/websocketserver"
	"{{ .ProjectName }}/src/interfaces"
)

type ServerContainer struct {
	config *conf.Config

	httpServer      interfaces.IServer
	grpcWebServer   interfaces.IServer
	grpcServer      interfaces.IServer
	websocketServer interfaces.IServer
}

func New(config *conf.Config) *ServerContainer {
	return &ServerContainer{
		config: config,
	}
}

func (c *ServerContainer) Inject(services interfaces.IService) *ServerContainer {
	gs := grpcserver.With(c.config).New(services)
	gs.Configure(nil)
	c.grpcServer = gs

	hs := httpserver.With(c.config).New(services)
	hs.Configure(nil)
	c.httpServer = hs

	gws := grpcwebserver.With(c.config).New(services)
	gws.Configure(hs)
	c.grpcWebServer = gws

	wss := websocketserver.With(c.config).New(services)
	wss.Configure(hs)
	c.websocketServer = wss

	return c
}

func (c *ServerContainer) Start() {
	go func() {
		address := c.config.Grpc.Address
		port := c.config.Grpc.Port
		if err := c.grpcServer.Start(address, port); err != nil {
			log.Fatal("gRPC server:", err)
		}
	}()
	address := c.config.Http.Address
	port := c.config.Http.Port

	go c.websocketServer.Start(address, port)

	if err := c.httpServer.Start(address, port); err != nil {
		log.Fatal("HTTP server:", err)
	}
}
