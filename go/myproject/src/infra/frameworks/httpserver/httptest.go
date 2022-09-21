package httpserver

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"

	"github.com/labstack/echo/v4"

	"{{ .ProjectName }}/src/infra/conf"
	"{{ .ProjectName }}/src/infra/frameworks/httpserver/models"
	"{{ .ProjectName }}/src/interfaces"
)

type HttpTest struct {
	services interfaces.IService
}

func WithStub(services interfaces.IService) HttpTest {
	return HttpTest{services: services}
}

func (h HttpTest) ServeHTTP(req *http.Request) (*http.Response, models.HttpResponse) {
	server := With(conf.Get()).New(h.services)
	server.Configure(nil)
	resp := httptest.NewRecorder()

	rawServer := server.GetRawServer().(*echo.Echo)
	rawServer.ServeHTTP(resp, req)
	result := resp.Result()
	body, err := io.ReadAll(result.Body)
	if err != nil {
		panic(err)
	}
	var response models.HttpResponse
	json.Unmarshal(body, &response)
	return result, response
}
