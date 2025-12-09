package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {

	// load env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// connect DB
	ConnectDB()

	// init echo
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello world! DB Connected. Echo API is running")
	})

	// todo router
	e.GET("/api/todos", GetAllTodos)

	port := os.Getenv("PORT")
	e.Logger.Fatal(e.Start(":" + port))
}
