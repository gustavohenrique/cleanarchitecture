package gogrpc

import (
	"fmt"
	"math"
	"os"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/keepalive"
)

const (
	DEFAULT_HOST = "server.myproject.com"
	DEFAULT_PORT = 80
)

type Config struct {
	Host         string
	Port         int
	Token        string
	PingInterval int
	Timeout      int
	TLS          struct {
		Key  string
		Cert string
	}
}

type Connection interface {
	grpc.ClientConnInterface
}

func Connect(config *Config) (Connection, error) {
	params := keepalive.ClientParameters{
		PermitWithoutStream: true, // send pings even without active streams
	}
	if config.PingInterval > 0 {
		params.Time = time.Duration(config.PingInterval) * time.Second // send pings every X seconds if there is no activity
	}
	if config.Timeout > 0 {
		params.Timeout = time.Duration(config.Timeout) * time.Second // wait X seconds for ping ack before considering the connection dead
	}
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithKeepaliveParams(params))
	call := grpc.WaitForReady(false)
	opts = append(opts, grpc.WithDefaultCallOptions(call))
	opts = append(opts, grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(math.MaxInt32)))
	opts, err := addTls(opts, config)
	if err != nil {
		return nil, err
	}
	address := getAddressOrDefaultFrom(config)
	conn, err := grpc.Dial(address, opts...)
	return conn, err
}

func getAddressOrDefaultFrom(config *Config) string {
	host := config.Host
	if host == "" {
		host = DEFAULT_HOST
	}
	port := config.Port
	if port < 1 {
		port = DEFAULT_PORT
	}
	return fmt.Sprintf("%s:%d", host, port)
}

func addTls(opts []grpc.DialOption, config *Config) ([]grpc.DialOption, error) {
	key := config.TLS.Key
	cert := config.TLS.Cert
	if cert == "" || key == "" {
		return append(opts, grpc.WithInsecure()), nil
	}
	_, keyStat := os.Stat(key)
	_, certStat := os.Stat(cert)
	if os.IsNotExist(keyStat) || os.IsNotExist(certStat) {
		return append(opts, grpc.WithInsecure()), fmt.Errorf("Certificate %s or Key %s not found", cert, key)
	}
	creds, err := credentials.NewClientTLSFromFile(cert, "")
	if err != nil {
		return append(opts, grpc.WithInsecure()), err
	}
	opts = append(opts, grpc.WithTransportCredentials(creds))
	return opts, nil
}
