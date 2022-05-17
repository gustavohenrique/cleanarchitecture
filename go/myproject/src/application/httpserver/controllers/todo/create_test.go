package todo_test

import (
	"testing"

	"myproject/src/application/httpserver"
	"myproject/src/shared/test"
	"myproject/src/shared/test/assert"
	"myproject/src/shared/test/httpclient"
)

func TestCreateWithEmptyBody(t *testing.T) {
	req := httpclient.DoPOST("/todo", "")
	serviceContainer := test.GetServiceContainer()
	resp, _ := httpserver.With(&serviceContainer).ServeHTTP(req)
	assert.HttpStatusCode(t, resp, 400)
}
