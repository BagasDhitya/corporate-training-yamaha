package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	// load env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// connect DB
	ConnectDB()
	ExportLogger()

	// init echo
	e := echo.New()

	e.Use(MonitoringMiddleware)

	// === CORS ===
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:5173"},

		AllowMethods:     []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete},
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, "Authorization"},
		AllowCredentials: true,
	}))

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello world! DB Connected. Echo API is running")
	})

	// todo router
	e.GET("/api/todos", GetAllTodos)
	e.GET("/api/todos/:id", GetById)
	e.POST("/api/todos", CreateTodo)

	port := os.Getenv("PORT")
	e.Logger.Fatal(e.Start(":" + port))
}
