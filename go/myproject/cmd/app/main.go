package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"runtime/debug"
	"syscall"

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
{{ if .HasHttpServer }}
	httpServer := infra.HttpServer(controllers.NewHttpControllers(repos))
{{ end }}
{{ if .HasGrpcServer }}
	grpcServer := infra.GrpcServer(controllers.NewGrpcControllers(repos))
{{ end }}
{{ if .HasGrpcWebServer }}
	grpcWebServer := infra.GrpcWebServer(controllers.NewGrpcWebControllers(repos))
	grpcWebServer.Configure(httpServer)
{{ end }}
{{ if .HasNatsServer }}
	natsServer := infra.NatsServer(controllers.NewNatsControllers(repos))
	natsServer.Configure()
{{ end }}

{{ if .HasNatsServer }}go natsServer.Start(){{ end }}
{{ if .HasGrpcServer }}go grpcServer.Start(){{ end }}
{{ if .HasHttpServer }}go httpServer.Start(){{ end }}
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT)
	<-quit
}
