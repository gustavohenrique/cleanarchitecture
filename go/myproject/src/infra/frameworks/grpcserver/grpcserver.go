package grpcserver

import (
	"context"
	"crypto/tls"
	"fmt"
	"log"
	"math"
	"net"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"{{ .ProjectName }}/src/infra/conf"
	"{{ .ProjectName }}/src/infra/frameworks/grpcserver/controllers"
	"{{ .ProjectName }}/src/interfaces"
	pb "{{ .ProjectName }}/src/proto"
)

type GrpcServer struct {
	config    *conf.Config
	services  interfaces.IService
	rawServer *grpc.Server
}

func With(config *conf.Config) interfaces.IServer {
	return &GrpcServer{config: config}
}

func (s *GrpcServer) New(services interfaces.IService) interfaces.IServer {
	var opts []grpc.ServerOption
	opts = append(opts, serverInterceptor(services))
	config := s.config
	if config.Grpc.TLS.Enabled {
		key := config.Grpc.TLS.Key
		cert := config.Grpc.TLS.Cert
		credentials := getCredentials(key, cert)
		opts = append(opts, grpc.Creds(credentials))
	}

	maxSendMsgSize := config.Grpc.MaxSendMsgSize
	if maxSendMsgSize < 1 {
		maxSendMsgSize = math.MaxInt32
	}
	maxRecvMsgSize := config.Grpc.MaxReceiveMsgSize
	if maxRecvMsgSize < 1 {
		maxRecvMsgSize = math.MaxInt32
	}
	opts = append(opts, grpc.MaxSendMsgSize(maxSendMsgSize))
	opts = append(opts, grpc.MaxRecvMsgSize(maxRecvMsgSize))
	s.rawServer = grpc.NewServer(opts...)
	s.services = services
	return s
}

func (g *GrpcServer) GetRawServer() interface{} {
	return g.rawServer
}

func (g *GrpcServer) Configure(params interface{}) {
	pb.RegisterTodoRpcServer(g.rawServer, controllers.NewTodoRpcController(g.services))
}

func (g *GrpcServer) Start(address string, port int) error {
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", address, port))
	if err != nil {
		return err
	}
	go func() {
		fmt.Printf("⇨ gRPC server started on %s%s:%d%s\n", string("\033[32m"), address, port, string("\033[0m"))
		log.Fatalln(g.rawServer.Serve(lis))
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT)
	<-quit
	_, cancel := context.WithCancel(context.Background())
	cancel()
	g.rawServer.GracefulStop()
	return nil
}

func shouldSkip(method string) bool {
	skipRouters := conf.Get().Grpc.SkipRouters
	for _, i := range skipRouters {
		if strings.HasSuffix(method, i) {
			return true
		}
	}
	return false
}

func serverInterceptor(services interfaces.IService) grpc.ServerOption {
	return grpc.UnaryInterceptor(
		func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
			if req != nil && shouldSkip(info.FullMethod) {
				return handler(ctx, req)
			}
			/*
				md, ok := metadata.FromIncomingContext(ctx)
				if !ok {
					return nil, customerror.Invalid("No credentials in gRPC metadata")
				}
				userID := md.Get("X-User-Id")
				token := md.Get("X-CSRF-Token")
				if len(userID) == 0 || len(token) == 0 {
					return nil, customerror.Invalid("X-User-Id or X-CSRF-Token are missing")
				}
				header := metadata.Pairs("X-CSRF-Token", token[0])
				grpc.SetHeader(ctx, header)
				withUserID := metadata.Pairs("user_id", userID[0])
				ctx = metadata.NewIncomingContext(ctx, metadata.Join(md, withUserID))
			*/
			return handler(ctx, req)
		},
	)
}

func getCredentials(key, cert string) credentials.TransportCredentials {
	certificate, err := tls.LoadX509KeyPair(cert, key)
	if err != nil {
		log.Fatalln("Could not load the certificates files:", key, cert)
	}
	return credentials.NewTLS(&tls.Config{
		Certificates:             []tls.Certificate{certificate},
		InsecureSkipVerify:       false,
		PreferServerCipherSuites: true,
		MinVersion:               tls.VersionTLS12,
		ClientAuth:               tls.RequireAndVerifyClientCert,
		CipherSuites: []uint16{
			tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
			tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
			tls.TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305,
			tls.TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305,
			tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
		},
	})
}
