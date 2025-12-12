package routes

import (
	"echo-todo-api/controllers"
	"echo-todo-api/services"

	"github.com/labstack/echo/v4"
)

func TodoRoutes(e *echo.Echo) {
	service := services.NewTodoService()
	controller := controllers.NewTodoController(service)

	e.GET("/api/todos", controller.GetAllTodos)
	e.GET("/api/todos/:id", controller.GetById)
	e.POST("/api/todos", controller.CreateTodo)
}
