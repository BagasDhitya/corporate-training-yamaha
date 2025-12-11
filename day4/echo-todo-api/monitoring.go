package main

import (
	"io"
	"log"
	"os"
	"time"

	"github.com/labstack/echo/v4"
)

func ExportLogger() {

	// pastikan folder logs
	if _, err := os.Stat("logs"); os.IsNotExist(err) {
		os.Mkdir("logs", 0755)
	}

	// buka file log
	file, err := os.OpenFile("logs/app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)

	if err != nil {
		log.Fatal("Failed to open log file : ", err)
	}

	// multi writer : sekaligus ke file dan console
	multiWriter := io.MultiWriter(os.Stdout, file)

	// set log output
	log.SetOutput(multiWriter)
}

func MonitoringMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		start := time.Now()
		err := next(c)
		stop := time.Now()

		latency := stop.Sub(start)

		log.Printf(`
		===== Request Log =====
		METHOD		: %s
		PATH		: %s
		STATUS CODE	: %s
		LATENCY		: %s
		IP ADDRESS	: %s
		USER AGENT	: %s
		`,
			c.Request().Method,
			c.Request().URL.Path,
			c.Response().Status,
			latency,
			c.RealIP(),
			c.Request().UserAgent())

		return err
	}
}