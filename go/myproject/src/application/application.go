package application

import (
	"flag"
	"fmt"
	"log"
	"os"

	"{{ .ProjectName }}/src/application/grpcserver"
	"{{ .ProjectName }}/src/application/grpcwebserver"
	"{{ .ProjectName }}/src/application/httpserver"
	"{{ .ProjectName }}/src/application/server"
	"{{ .ProjectName }}/src/application/websocketserver"
	"{{ .ProjectName }}/src/infra"
	"{{ .ProjectName }}/src/repositories"
	"{{ .ProjectName }}/src/services"
	"{{ .ProjectName }}/src/shared/conf"
	"{{ .ProjectName }}/src/shared/logger"
	"{{ .ProjectName }}/src/valueobjects"
)

type Application struct {
	config           *conf.Config
	flags            valueobjects.Flags
	serviceContainer services.ServiceContainer
	HttpServer       server.Server
	GrpcWebServer    server.Server
	GrpcServer       server.Server
	WebsocketServer  server.Server
}

func New() *Application {
	return &Application{}
}

func (a *Application) ParseCommandLineArgs() *Application {
	var flags valueobjects.Flags
	flag.StringVar(&flags.ConfigFile, "config", "config.yaml", "Configuration file")
	flag.Parse()
	a.flags = flags
	return a
}

func (a *Application) LoadConfigurationFile() *Application {
	configFile := os.Getenv("CONFIG_FILE")
	if configFile != "" {
		a.flags.ConfigFile = configFile
	}
	conf.Parse(a.flags.ConfigFile)
	a.config = conf.Get()
	fmt.Println("Using configuration file:", a.flags.ConfigFile)
	return a
}

func (a *Application) CreateServers() *Application {
	logger.Configure()
	infraContainer := infra.New()
	repositoryContainer := repositories.New(infraContainer)
	serviceContainer := services.New(repositoryContainer)
	a.serviceContainer = serviceContainer

	grpcServer := grpcserver.New(serviceContainer)
	grpcServer.Configure(nil)
	a.GrpcServer = grpcServer

	httpServer := httpserver.New(serviceContainer)
	httpServer.Configure(nil)
	a.HttpServer = httpServer

	grpcwebServer := grpcwebserver.New(serviceContainer)
	grpcwebServer.Configure(httpServer)
	a.GrpcWebServer = grpcwebServer

	websocketServer := websocketserver.New(serviceContainer)
	websocketServer.Configure(httpServer)
	a.WebsocketServer = websocketServer
	return a
}

func (a *Application) Start() {
	go func() {
		address := a.config.Grpc.Address
		port := a.config.Grpc.Port
		if err := a.GrpcServer.Start(address, port); err != nil {
			log.Fatal("gRPC server:", err)
		}
	}()
	address := a.config.Http.Address
	port := a.config.Http.Port
	go a.WebsocketServer.Start(address, port)
	if err := a.HttpServer.Start(address, port); err != nil {
		log.Fatal("HTTP server:", err)
	}
}
