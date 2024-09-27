package http

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
	"toDoListRestApi/src/internal/domain"
	"toDoListRestApi/src/internal/service"
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

// GetAll @Summary Get all users
// @Description Get all users with pagination
func (h *TodoHandler) GetAll(c echo.Context) error {
	offsetParam := c.QueryParam("offset")
	limitParam := c.QueryParam("limit")

	offset, err := strconv.Atoi(offsetParam)
	if err != nil {
		offset = 0
	}

	limit, err := strconv.Atoi(limitParam)
	if err != nil {
		limit = 10
	}

	todos, err := h.service.GetAllWithPagination(offset, limit)
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
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID"})
	}

	var existingTodo domain.Todo

	if _, err := h.service.GetByID(uint(id)); err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Todo not found"})
	}

	var updates map[string]interface{}
	if err := c.Bind(&updates); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	if title, ok := updates["title"]; ok {
		if titleStr, ok := title.(string); ok {
			existingTodo.Title = titleStr
		} else {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid type for title"})
		}
	}

	if completed, ok := updates["completed"]; ok {
		if completedBool, ok := completed.(bool); ok {
			existingTodo.Completed = completedBool
		} else {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid type for completed"})
		}
	}
	existingTodo.ID = uint(id)

	if err := h.service.Update(&existingTodo); err != nil {
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
