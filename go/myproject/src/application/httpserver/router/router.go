package router

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"

	"{{ .ProjectName }}/assets"
	"{{ .ProjectName }}/src/application/httpserver/controllers/auth"
	"{{ .ProjectName }}/src/application/httpserver/controllers/todo"
	"{{ .ProjectName }}/src/application/httpserver/templaterender"
	"{{ .ProjectName }}/src/services"
	"{{ .ProjectName }}/src/shared/conf"
)

type Router struct {
	serviceContainer services.ServiceContainer
	config           *conf.Config
	todoController   *todo.TodoController
	tokenController  *auth.TokenController
}

func NewRouter(serviceContainer services.ServiceContainer) *Router {
	return &Router{
		serviceContainer: serviceContainer,
		config:           conf.Get(),
		todoController:   todo.NewTodoController(serviceContainer),
		tokenController:  auth.NewTokenController(serviceContainer),
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
		return c.HTML(http.StatusOK, tpl)
	})
}

func (r *Router) ServeEmbedStaticFiles(e *echo.Echo, staticFile assets.StaticFile) {
	files := staticFile.GetFS()
	e.GET("/static/*", echo.WrapHandler(http.StripPrefix("/static/", files)))
}

func (r *Router) AddRestEndpoints(e *echo.Echo) {
	r.tokenController.AddRoutesTo(e.Group("token"))

	v1 := e.Group("/v1")
	r.addJwtMiddlewareTo(v1)
	r.todoController.AddRoutesTo(v1.Group("/todo"))
}

func (r *Router) addJwtMiddlewareTo(e *echo.Group) {
	e.Use(r.tokenController.Verify)
}
