package todo_test

import (
	"testing"

	"{{ .ProjectName }}/src/infra/frameworks/httpserver"
	"{{ .ProjectName }}/src/infra/test"
	"{{ .ProjectName }}/src/infra/test/assert"
	"{{ .ProjectName }}/src/infra/test/httpclient"
)

func TestCreateWithEmptyBody(t *testing.T) {
	req := httpclient.
		WithOauth2Mock().
		DoPOST("/v1/todo", "")
	services := test.NewTestHelper().WithOauth2Mock(t).GetServices()
	resp, _ := httpserver.WithStub(services).ServeHTTP(req)
	assert.HttpStatusCode(t, resp, 400)
}
