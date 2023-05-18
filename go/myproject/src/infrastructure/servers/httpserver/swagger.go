package httpserver

import (
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func (s *HttpServer) addSwaggerDocs() echo.HandlerFunc {
	return echoSwagger.WrapHandler
}
