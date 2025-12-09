package main

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type Todo struct {
	ID          int        `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Category    string     `json:"category"`
	isCompleted bool       `json:"isCompleted"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at"`
}

func GetAllTodos(c echo.Context) error {
	rows, err := DB.Query(`
		SELECT id, title, description, category, isCompleted, created_at, updated_at, deleted_at
		FROM todos
		WHERE deleted_at is NULL
		ORDER BY created_at DESC
	`)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Failed to fetch todos",
			"error":   err.Error(),
		})
	}

	defer rows.Close()

	var todos []Todo

	for rows.Next() {
		var t Todo
		err := rows.Scan(
			&t.ID,
			&t.Title,
			&t.Description,
			&t.Category,
			&t.isCompleted,
			&t.CreatedAt,
			&t.UpdatedAt,
			&t.DeletedAt,
		)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{
				"message": "Failed to parse todo data",
				"error":   err.Error(),
			})
		}

		todos = append(todos, t)
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "success",
		"data":    todos,
	})
}

func CreateTodo(c echo.Context) error {
	var input Todo

	// bind json input ke struct
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Invalid request body",
			"error":   err.Error(),
		})
	}

	// validasi : minimal title
	if input.Title == "" {
		return c.JSON(http.StatusBadGateway, echo.Map{
			"message": "Title is required",
		})
	}

	query := `INSERT INTO todos (title, description, category, isCompleted, created_at, updated_at)
			  VALUES ($1, $2, $3, $4, NOW(), NOW())
			  RETURNING id, created_at, updated_at`

	err := DB.QueryRow(
		query,
		input.Title,
		input.Description,
		input.Category,
		input.isCompleted,
	).Scan(&input.ID, &input.CreatedAt, &input.UpdatedAt)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Failed to create todo",
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, echo.Map{
		"message": "Todo created successfully",
		"data":    input,
	})
}
