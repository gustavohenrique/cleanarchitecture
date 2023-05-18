package httpserver

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"{{ .ProjectName }}/assets"
	"{{ .ProjectName }}/src/adapters/controllers"
	"{{ .ProjectName }}/src/components/configurator"
	"{{ .ProjectName }}/src/infrastructure/servers"
	_ "{{ .ProjectName }}/src/infrastructure/servers/httpserver/docs"
)

// @title My TODO application
// @version 1.0
// @contact.name Ravoni Team
// @contact.url https://ravoni.com
// @contact.email contact@ravoni.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8001
// @schemes http https
// @BasePath /v1

type HttpServer struct {
	rawServer   *echo.Echo
	config      *configurator.Config
	controllers controllers.HttpControllers
}

func New(config *configurator.Config, controllers controllers.HttpControllers) servers.Server {
	rawServer := echo.New()
	rawServer.Debug = config.Debug
	rawServer.HideBanner = true
	return &HttpServer{
		config:      config,
		rawServer:   rawServer,
		controllers: controllers,
	}
}

func (s *HttpServer) RawServer() interface{} {
	return s.rawServer
}

func (s *HttpServer) Configure(params ...interface{}) {
	s.addMiddlewares()

	e := s.rawServer
	// e.GET("/healthcheck", func(c echo.Context) error {
	// if metrics.CpuUsagePercentage() < 95.0 {
	// return c.String(http.StatusOK, "Ok")
	// }
	// return c.String(http.StatusServiceUnavailable, "")
	// })
	// e.GET("/metrics", func(c echo.Context) error {
	// c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	// return c.String(http.StatusOK, metrics.Collect())
	// })
	e.GET("/docs/*", s.addSwaggerDocs())

	s.serveEmbedWebPage(e, assets.NewWebPage())
	s.serveEmbedStaticFiles(e, assets.NewStaticFile())

	v1 := e.Group("/v1")
	// s.addJwtMiddlewareTo(v1)

	todo := v1.Group("/todo")
	todoController := s.controllers.TodoController()
	todo.GET("", todoController.ReadOne)
}

func (s *HttpServer) Start() error {
	addr := fmt.Sprintf("%s:%d", s.config.Http.Address, s.config.Http.Port)
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

// func (r *Router) addJwtMiddlewareTo(e *echo.Group) {
// e.Use(r.tokenController.Verify)
// }

func (s *HttpServer) serveEmbedWebPage(e *echo.Echo, webPage assets.WebPage) {
	group := e.Group("web")
	group.GET("", func(c echo.Context) error {
		return c.Redirect(307, "/web/index.html")
	})
	group.GET("/:filename", func(c echo.Context) error {
		filename := c.Param("filename")
		if !strings.HasSuffix(filename, ".html") {
			filename = filename + ".html"
		}
		content, _ := webPage.Lookup(filename)
		if content == "" {
			return c.String(http.StatusNotFound, fmt.Sprintf("Not found: %s\n", filename))
		}
		tpl := Parse(content, map[string]interface{}{
			"message": "Hello World!",
		})
		return c.HTML(http.StatusOK, tpl)
	})
}

func (s *HttpServer) serveEmbedStaticFiles(e *echo.Echo, staticFile assets.StaticFile) {
	files := staticFile.GetFS()
	e.GET("/static/*", echo.WrapHandler(http.StripPrefix("/static/", files)))
}

func (s *HttpServer) addMiddlewares() {
	e := s.rawServer
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     s.config.Http.Origins,
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
	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: 5,
		Skipper: func(c echo.Context) bool {
			return strings.Contains(c.Request().URL.Path, "docs")
		},
	}))
}
