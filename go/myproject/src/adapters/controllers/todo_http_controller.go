package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"{{ .ProjectName }}/src/domain/ports"
	"{{ .ProjectName }}/src/wire"
)

type TodoHttpController struct {
	todoRepository ports.TodoRepository
}

func NewTodoHttpController(repos ports.Repositories) *TodoHttpController {
	return &TodoHttpController{
		todoRepository: repos.TodoRepository(),
	}
}

// @Description Get TODO by ID
// @Accept json
// @Produce json
// @Param id path string true "TODO ID"
// @Router /v1/todo/{id} [get]
// @Success 200 {object} wire.HttpResponse{data=wire.TodoHttpResponse}
// @Failure 404 {object} wire.HttpResponse{}
func (h *TodoHttpController) ReadOne(c echo.Context) error {
	ctx := c.Request().Context()
	req := &wire.TodoHttpRequest{ID: c.Param("id")}
	res := wire.NewHttpResponse()
	found, err := h.todoRepository.ReadOne(ctx, req.ToModel())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, res.Error(err))
	}
	data := wire.TodoHttpResponse{}.Of(found)
	return c.JSON(http.StatusOK, res.Success(data))
}

// @Description Create TODO item
// @Accept json
// @Produce json
// @Param TODO body wire.TodoHttpRequest true "Payload"
// @Router /v1/todo [post]
// @Success 201 {object} wire.HttpResponse{data=wire.TodoHttpResponse}
// @Failure 400 {object} wire.HttpResponse{}
// @Failure 500 {object} wire.HttpResponse{}
func (h *TodoHttpController) Create(c echo.Context) error {
	var req wire.TodoHttpRequest
	res := wire.NewHttpResponse()
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, res.Error(err))
	}
	if err := req.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, res.Error(err))
	}
	ctx := c.Request().Context()
	saved, err := h.todoRepository.Create(ctx, req.ToModel())
	if err != nil {
		// logger.Error("Failed to create a Todo:", err)
		return c.JSON(http.StatusInternalServerError, res.Error(err))
	}
	created := wire.TodoHttpResponse{}.Of(saved)
	return c.JSON(http.StatusCreated, res.Success(created))
}
