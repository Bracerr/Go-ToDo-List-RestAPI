package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
	"toDoListRestApi/src/domain"
	"toDoListRestApi/src/service"
)

type TodoHandler struct {
	service service.TodoService
}

func NewTodoHandler(service service.TodoService) *TodoHandler {
	return &TodoHandler{service: service}
}

func (h *TodoHandler) Create(c echo.Context) error {
	var todo domain.Todo
	if err := c.Bind(&todo); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	if err := h.service.Create(&todo); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusCreated, todo)
}

func (h *TodoHandler) GetAll(c echo.Context) error {
	todos, err := h.service.GetAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, todos)
}

func (h *TodoHandler) GetByID(c echo.Context) error {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	todo, err := h.service.GetByID(uint(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, err)
	}
	return c.JSON(http.StatusOK, todo)
}

func (h *TodoHandler) Update(c echo.Context) error {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	var todo domain.Todo
	if err := c.Bind(&todo); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	todo.ID = uint(id)
	if err := h.service.Update(&todo); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.NoContent(http.StatusNoContent)
}

func (h *TodoHandler) Delete(c echo.Context) error {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	if err := h.service.Delete(uint(id)); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.NoContent(http.StatusNoContent)
}
