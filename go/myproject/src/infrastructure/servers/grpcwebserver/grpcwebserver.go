package grpcwebserver

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"github.com/labstack/echo/v4"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	"{{ .ProjectName }}/pb"
	"{{ .ProjectName }}/src/adapters/controllers"
	"{{ .ProjectName }}/src/components/configurator"
	"{{ .ProjectName }}/src/infrastructure/servers"
)

const FIVE_MB = 5 * 1024 * 1024 * 1024

type GrpcWebServer struct {
	rawServer   *grpc.Server
	wrappedGrpc *grpcweb.WrappedGrpcServer
	config      *configurator.Config
	controllers controllers.GrpcWebControllers
}

func New(config *configurator.Config, controllers controllers.GrpcWebControllers) servers.Server {
	var opts []grpc.ServerOption
	opts = append(opts, serverInterceptor(controllers))

	maxSendMsgSize := config.Grpc.MaxSendMsgSize
	if maxSendMsgSize < 1 {
		maxSendMsgSize = FIVE_MB // math.MaxInt32
	}
	maxRecvMsgSize := config.Grpc.MaxReceiveMsgSize
	if maxRecvMsgSize < 1 {
		maxRecvMsgSize = FIVE_MB
	}
	opts = append(opts, grpc.MaxSendMsgSize(maxSendMsgSize))
	opts = append(opts, grpc.MaxRecvMsgSize(maxRecvMsgSize))

	rawServer := grpc.NewServer(opts...)
	return &GrpcWebServer{
		config:      config,
		rawServer:   rawServer,
		wrappedGrpc: grpcweb.WrapServer(rawServer),
		controllers: controllers,
	}
}

func (g *GrpcWebServer) RawServer() interface{} {
	return g.rawServer
}

func (g *GrpcWebServer) Configure(params ...interface{}) {
	if len(params) != 1 {
		log.Fatalln("Please configure the gRPC-Web server to use the HTTP server")
	}
	wrapped := g.wrappedGrpc
	httpServer := params[0].(servers.Server)
	e := httpServer.RawServer().(*echo.Echo)
	// {{ .ProjectName }} is the package name in .proto file
	e.Any("{{ .ProjectName }}.*", func(c echo.Context) error {
		req := c.Request()
		res := c.Response()
		if wrapped.IsGrpcWebRequest(req) || wrapped.IsAcceptableGrpcCorsRequest(req) {
			wrapped.ServeHTTP(res, req)
		}
		return nil
	})

	{{ range .Models }}
	pb.Register{{ .CamelCaseName }}RpcServer(g.rawServer, g.controllers.{{ .CamelCaseName }}Controller())
	{{ end }}
}

// gRPC-Web uses the HTTPServer to run
func (g *GrpcWebServer) Start() error {
	return nil
}

// The skipped router are defined in a variable (for non authenticated users)
func shouldSkip(method string) bool {
	config := configurator.Get()
	skipRouters := config.Grpc.SkipRouters
	for _, i := range skipRouters {
		if strings.HasSuffix(method, i) {
			return true
		}
	}
	return false
}

// It's like a middleware for gRPC.
func serverInterceptor(controllers controllers.GrpcWebControllers) grpc.ServerOption {
	return grpc.UnaryInterceptor(
		func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
			if req != nil && shouldSkip(info.FullMethod) {
				return handler(ctx, req)
			}
			md, ok := metadata.FromIncomingContext(ctx)
			if !ok {
				return nil, fmt.Errorf("no credentials in gRPC metadata")
			}
			token := md.Get("X-CSRF-Token")
			if len(token) == 0 {
				return nil, fmt.Errorf("the X-CSRF-Token is missing")
			}
			header := metadata.Pairs("X-CSRF-Token", token[0])
			grpc.SetHeader(ctx, header)
			withUserID := metadata.Pairs("user_id", "123456")
			ctx = metadata.NewIncomingContext(ctx, metadata.Join(md, withUserID))
			return handler(ctx, req)
		},
	)
}
