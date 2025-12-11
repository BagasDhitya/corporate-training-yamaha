package middlewares

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		// ambil header authorization
		authHeader := c.Request().Header.Get("Authorization")

		if authHeader == "" {
			return c.JSON(http.StatusUnauthorized, echo.Map{
				"message": "Missing Authorization header",
			})
		}

		// cek format, harus : Bearer <token>
		var tokenString string
		_, err := fmt.Sscanf(authHeader, "Bearer %s", &tokenString)

		if err != nil {
			c.JSON(http.StatusUnauthorized, echo.Map{
				"message": "Invalid Authorization format",
			})
		}

		// parse token
		token, parseErr := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		log.Println("SECRET MIDDLEWARE:", os.Getenv("JWT_SECRET"))

		log.Println("token : ", token.Valid)
		log.Println("err : ", parseErr)

		if parseErr != nil || !token.Valid {
			return c.JSON(http.StatusUnauthorized, echo.Map{
				"message": "Invalid or expired token",
			})
		}

		// baca claims
		claims := token.Claims.(jwt.MapClaims)

		c.Set("user_id", int(claims["user_id"].(float64)))
		c.Set("role", claims["role"].(string))

		return next(c)
	}
}
