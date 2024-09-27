package app

import (
	"github.com/labstack/echo/v4"
	"toDoListRestApi/src/internal/configs"
	"toDoListRestApi/src/internal/delivery/http"
	"toDoListRestApi/src/internal/repository"
	"toDoListRestApi/src/internal/service"
	"toDoListRestApi/src/pkg/database"
)

func Run() {
	dbModel := configs.GetDbParams()
	db := database.NewClient(dbModel)

	todoRepo := repository.NewTodoRepository(db)
	todoService := service.NewTodoService(todoRepo)
	todoHandler := http.NewTodoHandler(todoService)

	e := echo.New()

	e.POST("/todos", todoHandler.Create)
	e.GET("/todos", todoHandler.GetAll)
	e.GET("/todos/:id", todoHandler.GetByID)
	e.PUT("/todos/:id", todoHandler.Update)
	e.DELETE("/todos/:id", todoHandler.Delete)

	e.Logger.Fatal(e.Start(":8080"))
}
