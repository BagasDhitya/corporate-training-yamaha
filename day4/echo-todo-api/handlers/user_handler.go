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

	// simpan token di HTTP Only Cookie
	cookie := new(http.Cookie)
	cookie.Name = "jwt"
	cookie.Value = token
	cookie.Path = "/"
	cookie.HttpOnly = true // tidak bisa diakses JS
	cookie.Secure = false  // pakai https -> true
	cookie.SameSite = http.SameSiteStrictMode
	cookie.MaxAge = 86400 // 1 hari

	c.SetCookie(cookie)

	return c.JSON(http.StatusOK, echo.Map{
		"message": "Login success",
		"role":    user.Role,
	})
}

// admin get all users
func AdminGetAllUsers(c echo.Context) error {
	role := c.Get("role").(string)

	if role != "ADMIN" {
		return c.JSON(http.StatusForbidden, echo.Map{
			"message": "You are not allowed to access this resource.",
		})
	}

	query := `SELECT id, email, role, created_at, updated_at FROM users WHERE deleted_at IS NULL`
	rows, err := QueryHelper(query)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Failed to fetch users",
		})
	}

	defer rows.Close()
	var users []models.User

	for rows.Next() {
		var u models.User
		rows.Scan(&u.ID, &u.Email, &u.Role, &u.CreatedAt, &u.UpdatedAt)
		users = append(users, u)
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "success",
		"data":    users,
	})
}
