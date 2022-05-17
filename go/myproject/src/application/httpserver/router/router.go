package router

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"

	"myproject/assets"
	"myproject/src/application/httpserver/controllers/todo"
	"myproject/src/application/httpserver/templaterender"
	"myproject/src/services"
	"myproject/src/shared/conf"
)

type Router struct {
	serviceContainer services.ServiceContainer
	config           *conf.Config
}

func NewRouter(serviceContainer services.ServiceContainer) *Router {
	return &Router{
		serviceContainer: serviceContainer,
		config:           conf.Get(),
	}
}

func (r *Router) ServeEmbedWebPage(e *echo.Echo, webPage assets.WebPage) {
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
		tpl := templaterender.Parse(content, map[string]interface{}{
			"message": "Hello World!",
		})
		return c.String(http.StatusOK, tpl)
	})
}

func (r *Router) ServeEmbedStaticFiles(e *echo.Echo, staticFile assets.StaticFile) {
	files := staticFile.GetFS()
	e.GET("/static/*", echo.WrapHandler(http.StripPrefix("/static/", files)))
}

func (r *Router) AddRestEndpoints(e *echo.Echo) {
	todo.NewTodoController(r.serviceContainer).AddRoutesTo(e.Group("todo"))
}
