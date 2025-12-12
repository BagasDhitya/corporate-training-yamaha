package services

import (
	"database/sql"
	"errors"

	"echo-todo-api/models"
	"echo-todo-api/utils"
	"echo-todo-api/config"
)

type UserService struct{}

func NewUserService() *UserService {
	return &UserService{}
}

// REGISTER
func (s *UserService) Register(req models.RegisterRequest) (int, error) {

	if req.Email == "" || req.Password == "" {
		return 0, errors.New("email & password required")
	}

	// hash password
	hashed, err := utils.HashPassword(req.Password)
	if err != nil {
		return 0, errors.New("failed to hash password")
	}

	query := `
		INSERT INTO users (email, password, role, created_at, updated_at)
		VALUES ($1, $2, 'GUEST', NOW(), NOW()) 
		RETURNING id
	`

	var userID int
	err = config.DB.QueryRow(query, req.Email, hashed).Scan(&userID)

	if err != nil {
		return 0, errors.New("email already exists")
	}

	return userID, nil
}

// LOGIN
func (s *UserService) Login(req models.LoginRequest) (models.User, string, error) {

	var user models.User

	query := `
		SELECT id, email, password, role
		FROM users
		WHERE email = $1 AND deleted_at IS NULL
	`

	err := config.DB.QueryRow(query, req.Email).Scan(&user.ID, &user.Email, &user.Password, &user.Role)

	if err == sql.ErrNoRows {
		return user, "", errors.New("invalid credentials")
	}

	// cek password
	if !utils.CheckPassword(user.Password, req.Password) {
		return user, "", errors.New("invalid password")
	}

	// buat token JWT
	token, err := utils.GenerateToken(user.ID, user.Role)
	if err != nil {
		return user, "", errors.New("failed to generate token")
	}

	return user, token, nil
}

// ADMIN - GET ALL USERS
func (s *UserService) AdminGetAllUsers() ([]models.User, error) {
	query := `SELECT id, email, role, created_at, updated_at FROM users WHERE deleted_at IS NULL`

	rows, err := config.DB.Query(query)
	if err != nil {
		return nil, errors.New("failed to fetch users")
	}

	defer rows.Close()

	var users []models.User

	for rows.Next() {
		var u models.User
		rows.Scan(&u.ID, &u.Email, &u.Role, &u.CreatedAt, &u.UpdatedAt)
		users = append(users, u)
	}

	return users, nil
}
