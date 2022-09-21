package grpcwebserver

import (
	"context"
	"strings"

	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"github.com/labstack/echo/v4"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	"{{ .ProjectName }}/src/infra/conf"
	"{{ .ProjectName }}/src/infra/frameworks/grpcwebserver/controllers"
	"{{ .ProjectName }}/src/interfaces"
	pb "{{ .ProjectName }}/src/proto"
	"{{ .ProjectName }}/src/shared/customerror"
)

type GrpcWebServer struct {
	config      *conf.Config
	services    interfaces.IService
	rawServer   *grpc.Server
	wrappedGrpc *grpcweb.WrappedGrpcServer
}

func With(config *conf.Config) interfaces.IServer {
	return &GrpcWebServer{config: config}
}

func (s *GrpcWebServer) New(services interfaces.IService) interfaces.IServer {
	var opts []grpc.ServerOption
	opts = append(opts, serverInterceptor(services))

	// It's increase to 5MB the maximum size allowed for requests and responses
	opts = append(opts, grpc.MaxSendMsgSize(5*1024*1024*1024))
	opts = append(opts, grpc.MaxRecvMsgSize(5*1024*1024*1024))
	rawServer := grpc.NewServer(opts...)
	s.rawServer = rawServer
	s.wrappedGrpc = grpcweb.WrapServer(rawServer)
	s.services = services
	return s
}

func (g *GrpcWebServer) GetRawServer() interface{} {
	return g.rawServer
}

func (g *GrpcWebServer) Configure(params interface{}) {
	wrapped := g.wrappedGrpc

	httpServer := params.(interfaces.IServer)
	e := httpServer.GetRawServer().(*echo.Echo)
	// {{ .ProjectName }} is the package name in .proto file
	e.Any("{{ .ProjectName }}.*", func(c echo.Context) error {
		req := c.Request()
		res := c.Response()
		if wrapped.IsGrpcWebRequest(req) || wrapped.IsAcceptableGrpcCorsRequest(req) {
			wrapped.ServeHTTP(res, req)
		}
		return nil
	})

	pb.RegisterTodoRpcServer(g.rawServer, controllers.NewTodoWebController(g.services))
}

func (g *GrpcWebServer) Start(address string, port int) error {
	return nil
}

// The skipped router are defined in a variable (for non authenticated users)
func shouldSkip(method string) bool {
	skipRouters := conf.Get().Auth.SkipRouters
	for _, i := range skipRouters {
		if strings.HasSuffix(method, i) {
			return true
		}
	}
	return false
}

// It's like a middleware for gRPC.
func serverInterceptor(services interfaces.IService) grpc.ServerOption {
	return grpc.UnaryInterceptor(
		func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
			if req != nil && shouldSkip(info.FullMethod) {
				return handler(ctx, req)
			}
			md, ok := metadata.FromIncomingContext(ctx)
			if !ok {
				return nil, customerror.Invalid("No credentials in gRPC metadata")
			}
			token := md.Get("X-CSRF-Token")
			if len(token) == 0 {
				return nil, customerror.Invalid("X-CSRF-Token is missing")
			}
			header := metadata.Pairs("X-CSRF-Token", token[0])
			grpc.SetHeader(ctx, header)
			withUserID := metadata.Pairs("user_id", "123456")
			ctx = metadata.NewIncomingContext(ctx, metadata.Join(md, withUserID))
			return handler(ctx, req)
		},
	)
}
