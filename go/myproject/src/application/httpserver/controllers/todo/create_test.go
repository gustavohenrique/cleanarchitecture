package todo_test

import (
	"testing"

	"{{ .ProjectName }}/src/application/httpserver"
	"{{ .ProjectName }}/src/shared/test"
	"{{ .ProjectName }}/src/shared/test/assert"
	"{{ .ProjectName }}/src/shared/test/httpclient"
)

func TestCreateWithEmptyBody(t *testing.T) {
	t.Skip()
	req := httpclient.DoPOST("/v1/todo", "")
	serviceContainer := test.GetServiceContainer()
	resp, _ := httpserver.With(&serviceContainer).ServeHTTP(req)
	assert.HttpStatusCode(t, resp, 400)
}
