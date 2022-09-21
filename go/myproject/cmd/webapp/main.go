package main

import (
	"flag"
	"fmt"
	"os"
	"runtime/debug"

	"{{ .ProjectName }}/src/domain/gateways"
	"{{ .ProjectName }}/src/domain/services"
	"{{ .ProjectName }}/src/infra/conf"
	"{{ .ProjectName }}/src/infra/datastores"
	"{{ .ProjectName }}/src/infra/frameworks"
	"{{ .ProjectName }}/src/infra/logger"
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
	conf.Parse(configFile)
	config := conf.Get()
	fmt.Println("Using configuration file:", configFile)

	logger.Configure()

	stores := datastores.With(config).New()
	gates := gateways.With(config).Inject(stores)
	servicez := services.With(config).Inject(gates)
	servers := frameworks.New(config).Inject(servicez)
	servers.Start()
}
