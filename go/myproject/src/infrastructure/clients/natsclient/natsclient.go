package natsclient

import (
	"fmt"
	"time"

	"github.com/nats-io/nats.go"

	"{{ .ProjectName }}/src/components/configurator"
)

type NatsClient interface {
	Publish(subject, message string) error
	Request(subject, message string) (string, error)
	Subscribe(subject string, fn func(string)) error
}

type natsClient struct {
	config   *configurator.Config
	conn     *nats.Conn
}

func New(config *configurator.Config) NatsClient {
	return &natsClient{
		config: config,
	}
}

func (n *natsClient) Publish(subject, message string) error {
	nc, err := n.getConnection()
	if err != nil {
		return err
	}
	// defer nc.Close()
	nc.Publish(subject, []byte(message))
	nc.Flush()
	return nc.LastError()
}

func (n *natsClient) Request(subject, message string) (string, error) {
	nc, err := n.getConnection()
	if err != nil {
		return "", err
	}
	// defer nc.Close()
	msg, err := nc.Request(subject, []byte(message), 2*time.Second)
	if err != nil {
		return "", err
	}
	return string(msg.Data), err
}

func (n *natsClient) Subscribe(subject string, fn func(string)) error {
	nc, err := n.getConnection()
	if err != nil {
		return err
	}
	nc.Subscribe(subject, func(msg *nats.Msg) {
		message := string(msg.Data)
		go fn(message)
	})
	nc.Flush()
	if err := nc.LastError(); err != nil {
		return err
	}
	return nil
}

func (n *natsClient) getConnection() (*nats.Conn, error) {
	if n.conn != nil && n.conn.IsConnected() {
		return n.conn, nil
	}
	config := n.config
	url := fmt.Sprintf("nats://%s:%d", config.Nats.Address, config.Nats.Port)
	opts := []nats.Option{nats.Name(config.Nats.ClientName)}
	nc, err := nats.Connect(url, opts...)
	if err != nil {
		return nil, err
	}
	n.conn = nc
	return n.conn, nil
}
