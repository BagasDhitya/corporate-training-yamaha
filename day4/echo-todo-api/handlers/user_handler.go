package handlers

import (
	"database/sql"
	"net/http"

	"echo-todo-api/config"
	"echo-todo-api/models"
	"echo-todo-api/utils"

	"github.com/labstack/echo/v4"
)

// register
func Register(c echo.Context) error {
	var req models.RegisterRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Invalid request",
		})
	}

	if req.Email == "" || req.Password == "" {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Email & password required",
		})
	}

	// hashing password sebelum masuk ke db
	hashed, err := utils.HashPassword(req.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Failed to hash password",
		})
	}

	query := `INSERT INTO users (email, password, role, created_at, updated_at)
			VALUES ($1, $2, 'GUEST', NOW(), NOW()) RETURNING id`

	// pengecekan apakah user sudah terdaftar/belum
	var userID int
	err = config.DB.QueryRow(query, req.Email, hashed).Scan(&userID)

	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Failed to register (email may already exist)",
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, echo.Map{
		"message": "Registered successfully",
		"user_id": userID,
	})
}

// login
func Login(c echo.Context) error {
	var req models.LoginRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Invalid request",
		})
	}

	var user models.User

	query := `SELECT id, email, password, role
			  FROM users
			  WHERE email = $1 AND deleted_at IS NULL`

	err := config.DB.QueryRow(query, req.Email).Scan(&user.ID, &user.Email, &user.Password, &user.Role)

	if err == sql.ErrNoRows {
		return c.JSON(http.StatusUnauthorized, echo.Map{
			"message": "Invalid credentials",
		})
	}

	if !utils.CheckPassword(user.Password, req.Password) {
		return c.JSON(http.StatusUnauthorized, echo.Map{
			"message": "Invalid password",
		})
	}

	// generate JWT
	token, err := utils.GenerateToken(user.ID, user.Role)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Failed to generate token",
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "Login success",
		"email":   user.Email,
		"token":   token,
		"role":    user.Role,
	})
}
