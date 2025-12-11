package main

import (
	"database/sql"
	"net/http"
	"strconv"
	"strings"

	"echo-todo-api/config"
	"echo-todo-api/models"

	"github.com/labstack/echo/v4"
)

func GetAllTodos(c echo.Context) error {
	search := c.QueryParam("search")
	category := c.QueryParam("category")
	sort := c.QueryParam("sort")

	// default sorting : DESC
	if sort != "asc" {
		sort = "desc"
	}

	query := `SELECT id, title, description, category, isCompleted, created_at, updated_at, deleted_at
		FROM todos
		WHERE deleted_at is NULL`

	// filter by category
	if category != "" {
		query = query + " AND category = $CATEGORY"
	}

	// search: title or description
	if search != "" {
		query = query + " AND (LOWER(title) LIKE $SEARCH OR LOWER(description) LIKE $SEARCH) "
	}

	// sorting
	query = query + " ORDER BY created_at " + sort

	// --- dynamic parameter mapping
	var params []interface{}
	paramIndex := 1

	if category != "" {
		query = strings.Replace(query, "$CATEGORY", "$"+strconv.Itoa(paramIndex), 1)
		params = append(params, category)
		paramIndex++
	}

	if search != "" {
		query = strings.Replace(query, "$SEARCH", "$"+strconv.Itoa(paramIndex), -1)
		params = append(params, "%"+strings.ToLower(search)+"%")
		paramIndex++
	}

	rows, err := config.DB.Query(query, params...)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Failed to fetch todos",
			"error":   err.Error(),
		})
	}

	defer rows.Close()

	var todos []models.Todo

	for rows.Next() {
		var t models.Todo
		err := rows.Scan(
			&t.ID,
			&t.Title,
			&t.Description,
			&t.Category,
			&t.IsCompleted,
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

func GetById(c echo.Context) error {

	id := c.Param("id")

	var t models.Todo

	query := `SELECT id, title, description, category, isCompleted, created_at, updated_at, deleted_at
			  FROM todos
			  WHERE id = $1 AND deleted_at IS NULL`

	err := config.DB.QueryRow(query, id).Scan(
		&t.ID,
		&t.Title,
		&t.Description,
		&t.Category,
		&t.IsCompleted,
		&t.CreatedAt,
		&t.UpdatedAt,
		&t.DeletedAt,
	)

	if err == sql.ErrNoRows {
		return c.JSON(http.StatusNotFound, echo.Map{
			"message": "Todo not found",
		})
	}

	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Failed to fetch todo",
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "success",
		"data":    t,
	})
}

func CreateTodo(c echo.Context) error {
	var input models.Todo

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

	err := config.DB.QueryRow(
		query,
		input.Title,
		input.Description,
		input.Category,
		input.IsCompleted,
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
