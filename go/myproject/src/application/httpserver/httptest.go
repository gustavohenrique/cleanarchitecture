package httpserver

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"

	"github.com/labstack/echo/v4"

	"myproject/src/services"
	"myproject/src/valueobjects"
)

type HttpTest struct {
	serviceContainer *services.ServiceContainer
}

func With(serviceContainer *services.ServiceContainer) HttpTest {
	return HttpTest{serviceContainer}
}

func (h HttpTest) ServeHTTP(req *http.Request) (*http.Response, valueobjects.HttpResponse) {
	server := New(*h.serviceContainer)
	server.Configure(nil)
	resp := httptest.NewRecorder()

	rawServer := server.GetRawServer().(*echo.Echo)
	rawServer.ServeHTTP(resp, req)
	result := resp.Result()
	body, err := ioutil.ReadAll(result.Body)
	if err != nil {
		panic(err)
	}
	var response valueobjects.HttpResponse
	json.Unmarshal(body, &response)
	return result, response
}
