package main

import (
	"log"
	"time"

	"github.com/labstack/echo/v4"
)

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
