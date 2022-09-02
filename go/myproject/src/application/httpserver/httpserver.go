package httpserver

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"{{ .ProjectName }}/assets"
	r "{{ .ProjectName }}/src/application/httpserver/router"
	"{{ .ProjectName }}/src/application/server"
	"{{ .ProjectName }}/src/services"
	"{{ .ProjectName }}/src/shared/conf"
	log "{{ .ProjectName }}/src/shared/logger"
	"{{ .ProjectName }}/src/shared/metrics"
)

type HttpServer struct {
	rawServer        *echo.Echo
	config           *conf.Config
	serviceContainer services.ServiceContainer
	router           *r.Router
}

func New(serviceContainer services.ServiceContainer) server.Server {
	config := conf.Get()
	rawServer := echo.New()
	rawServer.Debug = config.Debug
	rawServer.HideBanner = true

	return &HttpServer{
		rawServer:        rawServer,
		config:           config,
		serviceContainer: serviceContainer,
		router:           r.NewRouter(serviceContainer),
	}
}

func (s *HttpServer) GetRawServer() interface{} {
	return s.rawServer
}

func (s *HttpServer) Configure(params interface{}) {
	s.addMiddlewares()

	e := s.rawServer
	e.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "ok")
	})
	e.GET("/metrics", func(c echo.Context) error {
		return c.String(http.StatusOK, metrics.Collect())
	})
	s.router.ServeEmbedWebPage(e, assets.NewWebPage())
	s.router.ServeEmbedStaticFiles(e, assets.NewStaticFile())

	s.router.AddRestEndpoints(e)
}

func (s *HttpServer) Start(address string, port int) error {
	addr := fmt.Sprintf("%s:%d", address, port)
	e := s.rawServer
	go func() {
		if s.config.Http.TLS.Enabled {
			key := s.config.Http.TLS.Key
			cert := s.config.Http.TLS.Cert
			log.Fatal(e.StartTLS(addr, cert, key))
		}
		log.Fatal(e.Start(addr))
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT)
	<-quit
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		return err
	}
	return nil
}

func (s *HttpServer) addMiddlewares() {
	e := s.rawServer
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowCredentials: true,
		AllowMethods: []string{
			http.MethodOptions,
			http.MethodGet,
			http.MethodPut,
			http.MethodPost,
			http.MethodDelete,
			http.MethodHead,
		},
		ExposeHeaders: []string{
			"grpc-status",
			"grpc-message",
			"grpc-timeout",
			"content-length",
			"X-Auth-Token",
		},
		AllowHeaders: []string{
			"Accept",
			"Accept-Encoding",
			"Authorization",
			"XMLHttpRequest",
			"X-Requested-With",
			"X-Request-ID",
			"X-Auth-Token",
			"X-User-Id",
			"X-user-agent",
			"X-grpc-web",
			"grpc-status",
			"grpc-message",
			"grpc-timeout",
			"Content-Type",
			"Content-Length",
			"User-Agent",
			"X-Amzn-Trace-Id",
			"X-Forwarded-For",
			"X-Forwarded-Port",
			"X-Real-Ip",
			"X-SDK-Version",
			"X-SDK-Agent",
		},
	}))
	e.Use(middleware.BodyLimit("10M"))
	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{Level: 5}))
}
