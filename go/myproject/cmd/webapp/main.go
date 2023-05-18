package main

import (
	"flag"
	"fmt"
	"os"
	"runtime/debug"

	"{{ .ProjectName }}/src/adapters/controllers"
	"{{ .ProjectName }}/src/adapters/repositories"
	"{{ .ProjectName }}/src/components/configurator"
	"{{ .ProjectName }}/src/infrastructure"
)

func init() {
	debug.SetGCPercent(500)
}

func main() {
	var configFile string
	flag.StringVar(&configFile, "config", "config.yaml", "Configuration file")
	flag.Parse()

	if configFile == "" {
		configFile = os.Getenv("CONFIG_FILE")
	}
	fmt.Println("Using configuration file:", configFile)

	config := configurator.Parse(configFile)
	infra := infrastructure.New(config)
	repos := repositories.New(infra.DataStores())

	httpServer := infra.HttpServer(controllers.NewHttpControllers(repos))
	grpcServer := infra.GrpcServer(controllers.NewGrpcControllers(repos))

	grpcWebServer := infra.GrpcWebServer(controllers.NewGrpcWebControllers(repos))
	grpcWebServer.Configure(httpServer)

	natsServer := infra.NatsServer(controllers.NewNatsControllers(repos))
	natsServer.Configure()

	go grpcServer.Start()
	go natsServer.Start()
	httpServer.Start()
}
