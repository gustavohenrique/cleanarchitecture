package todo

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"myproject/src/adapters"
	"myproject/src/interfaces"
	"myproject/src/services"
	"myproject/src/shared/conf"
	"myproject/src/shared/customerror"
	log "myproject/src/shared/logger"
	"myproject/src/valueobjects"
)

type TodoController struct {
	config          *conf.Config
	todoItemAdapter adapters.TodoItemAdapter
	todoService     interfaces.ITodoService
}

func NewTodoController(serviceContainer services.ServiceContainer) *TodoController {
	return &TodoController{
		todoService:     serviceContainer.GetTodoService(),
		todoItemAdapter: adapters.NewTodoItemAdapter(),
		config:          conf.Get(),
	}
}

func (h *TodoController) AddRoutesTo(group *echo.Group) {
	group.GET("", h.ReadAll)
	group.POST("", h.Create)
	group.PUT("/:id", h.Update)
}

func (h *TodoController) ReadAll(c echo.Context) error {
	res := valueobjects.NewHttpResponse()
	ctx := c.Request().Context()
	items, err := h.todoService.ReadAll(ctx)
	if err != nil {
		res.SetError(err)
		return c.JSON(http.StatusInternalServerError, res)
	}
	res.SetData(h.todoItemAdapter.FromEntityListToResponseList(items))
	return c.JSON(http.StatusOK, res)
}

func (h *TodoController) Create(c echo.Context) error {
	var req valueobjects.TodoItemRequest
	res := valueobjects.NewHttpResponse()
	if err := c.Bind(&req); err != nil {
		log.Error("Cannot marshal JSON from request:", err)
		res.SetError(err, "Cannot marshal JSON")
		return c.JSON(http.StatusBadRequest, res)
	}
	if err := req.Validate(); err != nil {
		res.SetError(err, "Invalid request")
		return c.JSON(customerror.StatusCodeFrom(err), res)
	}
	entity := h.todoItemAdapter.FromRequestToEntity(req)
	ctx := c.Request().Context()
	saved, err := h.todoService.Create(ctx, entity)
	if err != nil {
		log.Error("Failed to create a TodoItem:", err)
		res.SetError(err)
		return c.JSON(http.StatusInternalServerError, res)
	}
	res.SetData(h.todoItemAdapter.FromEntityToResponse(saved))
	return c.JSON(http.StatusCreated, res)
}

func (h *TodoController) Update(c echo.Context) error {
	var req valueobjects.TodoItemRequest
	res := valueobjects.NewHttpResponse()
	if err := c.Bind(&req); err != nil {
		log.Error("Cannot marshal JSON from request:", err)
		res.SetError(err, "Cannot marshal JSON")
		return c.JSON(http.StatusBadRequest, res)
	}
	req.ID = c.Param("id")
	entity := h.todoItemAdapter.FromRequestToEntity(req)
	ctx := c.Request().Context()
	saved, err := h.todoService.Update(ctx, entity)
	if err != nil {
		log.Error("Failed to update the TodoItem:", err)
		res.SetError(err)
		return c.JSON(http.StatusInternalServerError, res)
	}
	res.SetData(h.todoItemAdapter.FromEntityToResponse(saved))
	return c.JSON(http.StatusCreated, res)
}
