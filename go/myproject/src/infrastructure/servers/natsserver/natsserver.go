package natsserver

import (
	"fmt"
	"log"

	"github.com/nats-io/nats-server/v2/server"

	"{{ .ProjectName }}/src/adapters/controllers"
	"{{ .ProjectName }}/src/components/configurator"
	"{{ .ProjectName }}/src/infrastructure/clients/natsclient"
	"{{ .ProjectName }}/src/infrastructure/servers"
)

type NatsServer struct {
	rawServer   *server.Server
	config      *configurator.Config
	controllers controllers.NatsControllers
}

func New(config *configurator.Config, controllers controllers.NatsControllers) servers.Server {
	opts := &server.Options{
		ServerName: config.Nats.ServerName,
		Host:       config.Nats.Address,
		Port:       config.Nats.Port,
		NoSigs:     config.Nats.NoSigs,
		Debug:      config.Nats.Debug,
		Trace:      config.Nats.Trace,
	}
	rawServer, err := server.NewServer(opts)
	if err != nil {
		log.Fatal("Nats server error:", err)
	}
	// rawServer.ConfigureLogger()
	return &NatsServer{
		config:      config,
		rawServer:   rawServer,
		controllers: controllers,
	}
}

func (n *NatsServer) RawServer() interface{} {
	return n.rawServer
}

func (n *NatsServer) Configure(params ...interface{}) {
	client := natsclient.New(n.config)
	todoController := n.controllers.With(client).TodoController()
	todoController.SubscribeToNewTodo()
}

func (n *NatsServer) Start() error {
	// if err := server.Run(n.rawServer); err != nil {
	// return err
	// }
	go n.rawServer.Start()
	var tls string
	if n.config.Nats.TLS.Enabled {
		tls = "(TLS enabled)"
	}
	address := n.config.Nats.Address
	port := n.config.Nats.Port
	fmt.Printf("⇨ nats server started on %s%s:%d%s %s\n", string("\033[32m"), address, port, string("\033[0m"), tls)
	n.rawServer.Shutdown()
	return nil
}
