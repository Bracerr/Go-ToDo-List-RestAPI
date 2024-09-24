package main

import (
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	"toDoListRestApi/src/domain"
	"toDoListRestApi/src/handler"
	"toDoListRestApi/src/repository"
	"toDoListRestApi/src/service"
)

func main() {
	envErr := godotenv.Load(".env")

	if envErr != nil {
		log.Fatal("Error loading .env file")
	}
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")

	dsn := "host=" + dbHost + " user=" + dbUser + " password=" + dbPassword + " dbname=" + dbName + " port=" + dbPort + " sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	errDb := db.AutoMigrate(&domain.Todo{})
	if errDb != nil {
		return
	}

	todoRepo := repository.NewTodoRepository(db)
	todoService := service.NewTodoService(todoRepo)
	todoHandler := handler.NewTodoHandler(todoService)

	e := echo.New()

	e.POST("/todos", todoHandler.Create)
	e.GET("/todos", todoHandler.GetAll)
	e.GET("/todos/:id", todoHandler.GetByID)
	e.PUT("/todos/:id", todoHandler.Update)
	e.DELETE("/todos/:id", todoHandler.Delete)

	e.Logger.Fatal(e.Start(":8080"))
}
