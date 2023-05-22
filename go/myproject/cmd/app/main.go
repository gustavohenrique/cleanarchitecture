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

{{ if and .HasNatsServer .HasHttpServer }}
	go natsServer.Start()
	httpServer.Start()
{{ end }}
{{ if and .HasGrpcServer .HasHttpServer}}
    go grpcServer.Start()
	httpServer.Start()
{{ else }}
	{{ if .HasGrpcServer }}
	grpcServer.Start()
	{{ else }}
		{{ if .HasNatsServer }}
		natsServer.Start()
		{{ else }}
		httpServer.Start()
		{{ end }}
	{{ end }}
{{ end }}
}
