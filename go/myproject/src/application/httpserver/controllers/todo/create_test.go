package todo_test

import (
	"testing"

	"{{ .ProjectName }}/src/application/httpserver"
	"{{ .ProjectName }}/src/shared/test"
	"{{ .ProjectName }}/src/shared/test/assert"
	"{{ .ProjectName }}/src/shared/test/httpclient"
)

func TestCreateWithEmptyBody(t *testing.T) {
	req := httpclient.
		WithOauth2Mock().
		DoPOST("/v1/todo", "")
	serviceContainer := test.NewTestHelper().WithOauth2Mock(t).GetServiceContainer()
	resp, _ := httpserver.With(&serviceContainer).ServeHTTP(req)
	assert.HttpStatusCode(t, resp, 400)
}
