package middlewares

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func AdminMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		role := c.Get("role").(string)

		if role != "ADMIN" {
			return c.JSON(http.StatusForbidden, echo.Map{
				"message": "Forbidden: Only admin can access this resources",
			})
		}

		return next(c)
	}
}
