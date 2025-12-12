package routes

import (
	"echo-todo-api/controllers"
	"echo-todo-api/services"

	"echo-todo-api/middlewares"

	"github.com/labstack/echo/v4"
)

func UserRoutes(e *echo.Echo) {
	service := services.NewUserService()
	controller := controllers.NewUserController(service)

	e.POST("/api/auth/register", controller.Register)
	e.POST("/api/auth/login", controller.Login)
	e.GET("/api/auth/users", middlewares.AuthMiddleware(middlewares.AdminMiddleware(controller.AdminGetAllUsers)))
}
