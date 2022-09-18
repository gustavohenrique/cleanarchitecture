package todo

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"{{ .ProjectName }}/src/infra/conf"
	"{{ .ProjectName }}/src/infra/frameworks/httpserver/models"
	"{{ .ProjectName }}/src/infra/logger"
	"{{ .ProjectName }}/src/interfaces"
	"{{ .ProjectName }}/src/shared/customerror"
)

type TodoController struct {
	config      *conf.Config
	todoService interfaces.ITodoService
	adapter     *Adapter
}

func With(config *conf.Config) *TodoController {
	return &TodoController{config: config}
}

func (h *TodoController) New(services interfaces.IService) *TodoController {
	h.todoService = services.GetTodoService()
	return h
}

func (h *TodoController) AddRoutesTo(group *echo.Group) {
	group.GET("", h.ReadAll)
	group.POST("", h.Create)
	group.PUT("/:id", h.Update)
}

// @Description Get all TODO items
// @Accept json
// @Produce json
// @Success 200 {object} models.HttpResponse{data=TodoJsonResponse}
// @Router /todo [get]
func (h *TodoController) ReadAll(c echo.Context) error {
	res := models.NewHttpResponse()
	ctx := c.Request().Context()
	items, err := h.todoService.ReadAll(ctx)
	if err != nil {
		res.SetError(err)
		return c.JSON(http.StatusInternalServerError, res)
	}
	res.SetData(h.adapter.ToJsonResponses(items))
	return c.JSON(http.StatusOK, res)
}

// @Description Create TODO item
// @Accept json
// @Produce json
// @Param TODO body TodoJsonRequest true "Payload"
// @Success 201 {object} models.HttpResponse{data=TodoJsonResponse}
// @Router /todo [post]
func (h *TodoController) Create(c echo.Context) error {
	var req TodoJsonRequest
	res := models.NewHttpResponse()
	if err := c.Bind(&req); err != nil {
		logger.Error("Cannot marshal JSON from request:", err)
		res.SetError(err, "Cannot marshal JSON")
		return c.JSON(http.StatusBadRequest, res)
	}
	if err := req.Validate(); err != nil {
		res.SetError(err, "Invalid request")
		return c.JSON(customerror.StatusCodeFrom(err), res)
	}
	entity := req.ToEntity()
	ctx := c.Request().Context()
	saved, err := h.todoService.Create(ctx, entity)
	if err != nil {
		logger.Error("Failed to create a Todo:", err)
		res.SetError(err)
		return c.JSON(http.StatusInternalServerError, res)
	}
	res.SetData(h.adapter.ToJsonResponse(saved))
	return c.JSON(http.StatusCreated, res)
}

func (h *TodoController) Update(c echo.Context) error {
	var req TodoJsonRequest
	res := models.NewHttpResponse()
	if err := c.Bind(&req); err != nil {
		logger.Error("Cannot marshal JSON from request:", err)
		res.SetError(err, "Cannot marshal JSON")
		return c.JSON(http.StatusBadRequest, res)
	}
	req.ID = c.Param("id")
	entity := req.ToEntity()
	ctx := c.Request().Context()
	saved, err := h.todoService.Update(ctx, entity)
	if err != nil {
		logger.Error("Failed to update the Todo:", err)
		res.SetError(err)
		return c.JSON(http.StatusInternalServerError, res)
	}
	res.SetData(h.adapter.ToJsonResponse(saved))
	return c.JSON(http.StatusCreated, res)
}
