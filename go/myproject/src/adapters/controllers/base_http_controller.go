package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"{{ .ProjectName }}/src/domain/ports"
	"{{ .ProjectName }}/src/wire"
)

{{ range .Models }}
type {{ .CamelCaseName }}HttpController struct {
	{{ .LowerCaseName }}Repository ports.{{ .CamelCaseName }}Repository
}

func New{{ .CamelCaseName }}HttpController(repos ports.Repositories) *{{ .CamelCaseName }}HttpController {
	return &{{ .CamelCaseName }}HttpController{
		{{ .LowerCaseName }}Repository: repos.{{ .CamelCaseName }}Repository(),
	}
}

// @Description Get {{ .LowerCaseName }} by ID
// @Accept json
// @Produce json
// @Param id path string true "{{ .LowerCaseName }} ID"
// @Router /v1/{{ .LowerCaseName }}/{id} [get]
// @Success 200 {object} wire.HttpResponse{data=wire.{{ .CamelCaseName }}HttpResponse}
// @Failure 404 {object} wire.HttpResponse{}
func (h *{{ .CamelCaseName }}HttpController) ReadOne(c echo.Context) error {
	ctx := c.Request().Context()
	req := &wire.{{ .CamelCaseName }}HttpRequest{ID: c.Param("id")}
	res := wire.NewHttpResponse()
	found, err := h.{{ .LowerCaseName }}Repository.ReadOne(ctx, req.ToModel())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, res.Error(err))
	}
	data := wire.{{ .CamelCaseName }}HttpResponse{}.Of(found)
	return c.JSON(http.StatusOK, res.Success(data))
}

// @Description Create {{ .LowerCaseName }} item
// @Accept json
// @Produce json
// @Param {{ .LowerCaseName }} body wire.{{ .CamelCaseName }}HttpRequest true "Payload"
// @Router /v1/{{ .LowerCaseName }} [post]
// @Success 201 {object} wire.HttpResponse{data=wire.{{ .CamelCaseName }}HttpResponse}
// @Failure 400 {object} wire.HttpResponse{}
// @Failure 500 {object} wire.HttpResponse{}
func (h *{{ .CamelCaseName }}HttpController) Create(c echo.Context) error {
	var req wire.{{ .CamelCaseName }}HttpRequest
	res := wire.NewHttpResponse()
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, res.Error(err))
	}
	if err := req.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, res.Error(err))
	}
	ctx := c.Request().Context()
	saved, err := h.{{ .LowerCaseName }}Repository.Create(ctx, req.ToModel())
	if err != nil {
		// logger.Error("Failed to create a {{ .CamelCaseName }}:", err)
		return c.JSON(http.StatusInternalServerError, res.Error(err))
	}
	created := wire.{{ .CamelCaseName }}HttpResponse{}.Of(saved)
	return c.JSON(http.StatusCreated, res.Success(created))
}
{{ end }}
