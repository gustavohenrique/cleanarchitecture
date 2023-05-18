package httpserver

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"

	"github.com/labstack/echo/v4"

	"{{ .ProjectName }}/src/adapters/controllers"
	"{{ .ProjectName }}/src/components/configurator"
	"{{ .ProjectName }}/src/wire"
)

type HttpTest struct {
	controllers controllers.HttpControllers
}

func WithStub(controllers controllers.HttpControllers) HttpTest {
	return HttpTest{controllers}
}

func (h HttpTest) ServeHTTP(req *http.Request) (*http.Response, wire.HttpResponse) {
	config := configurator.Get()
	server := New(config, h.controllers)
	server.Configure(nil)
	resp := httptest.NewRecorder()

	rawServer := server.RawServer().(*echo.Echo)
	rawServer.ServeHTTP(resp, req)
	result := resp.Result()
	body, err := io.ReadAll(result.Body)
	if err != nil {
		panic(err)
	}
	var response wire.HttpResponse
	json.Unmarshal(body, &response)
	return result, response
}
