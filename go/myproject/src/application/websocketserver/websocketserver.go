package websocketserver

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"

	"{{ .ProjectName }}/src/application/server"
	"{{ .ProjectName }}/src/services"
	"{{ .ProjectName }}/src/shared/conf"
	"{{ .ProjectName }}/src/shared/logger"
	"{{ .ProjectName }}/src/valueobjects"
)

type WebsocketServer struct {
	config           *conf.Config
	serviceContainer services.ServiceContainer
}

func New(serviceContainer services.ServiceContainer) server.Server {
	return &WebsocketServer{
		config:           conf.Get(),
		serviceContainer: serviceContainer,
	}
}

func (h *WebsocketServer) GetRawServer() interface{} {
	return nil
}

func (h *WebsocketServer) Configure(params interface{}) {
	httpServer := params.(server.Server)
	e := httpServer.GetRawServer().(*echo.Echo)
	prefix := h.config.Websocket.RouterPrefix
	e.GET(prefix+"/:room", func(c echo.Context) error {
		id := c.Param("room")
		logger.Info("Room:", id)
		h.serveWs(c.Response(), c.Request(), id)
		return nil
	})
	e.POST(prefix+"/:room", func(c echo.Context) error {
		type M struct {
			Message string `json:"message"`
		}
		res := valueobjects.NewHttpResponse()
		var req M
		if err := c.Bind(&req); err != nil {
			logger.Error("Cannot marshal JSON from request:", err)
			res.SetError(err, "Cannot marshal JSON")
			return c.JSON(http.StatusBadRequest, res)
		}
		id := c.Param("room")
		message := req.Message
		hub.SendTo(id, message)
		res.SetData("message sent")
		return c.JSON(http.StatusOK, res)
	})
}

func (h *WebsocketServer) Start(address string, port int) error {
	fmt.Printf("⇨ Websocket server started on %s%s:%d%s\n", string("\033[32m"), address, port, string("\033[0m"))
	hub.run()
	return nil
}
