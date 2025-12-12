package controllers

import (
	"echo-todo-api/models"
	"echo-todo-api/services"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	Service *services.UserService
}

func NewUserController(service *services.UserService) *UserController {
	return &UserController{Service: service}
}

// REGISTER
func (ctl *UserController) Register(c echo.Context) error {
	var req models.RegisterRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "invalid request body",
		})
	}

	userID, err := ctl.Service.Register(req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, echo.Map{
		"message": "Registered successfully",
		"user_id": userID,
	})
}

// LOGIN
func (ctl *UserController) Login(c echo.Context) error {
	var req models.LoginRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "invalid request",
		})
	}

	user, token, err := ctl.Service.Login(req)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{
			"message": err.Error(),
		})
	}

	// set JWT cookie
	cookie := new(http.Cookie)
	cookie.Name = "jwt"
	cookie.Value = token
	cookie.Path = "/"
	cookie.HttpOnly = true
	cookie.Secure = false
	cookie.SameSite = http.SameSiteStrictMode
	cookie.MaxAge = 86400

	c.SetCookie(cookie)

	return c.JSON(http.StatusOK, echo.Map{
		"message": "Login success",
		"role":    user.Role,
	})
}

// ADMIN GET ALL USERS
func (ctl *UserController) AdminGetAllUsers(c echo.Context) error {

	role := c.Get("role").(string)
	if role != "ADMIN" {
		return c.JSON(http.StatusForbidden, echo.Map{
			"message": "You are not allowed to access this resource.",
		})
	}

	users, err := ctl.Service.AdminGetAllUsers()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "success",
		"data":    users,
	})
}
