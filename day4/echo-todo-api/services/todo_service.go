package services

import (
	"database/sql"
	"strconv"
	"strings"

	"echo-todo-api/config"
	"echo-todo-api/models"
)

type TodoService struct{}

func NewTodoService() *TodoService {
	return &TodoService{}
}

func (s *TodoService) GetAll(search, category, sort string) ([]models.Todo, error) {
	// default sorting : DESC
	if sort != "asc" {
		sort = "desc"
	}

	query := `SELECT id, title, description, category, isCompleted, created_at, updated_at, deleted_at
		FROM todos
		WHERE deleted_at is NULL`

	if category != "" {
		query = query + " AND category = $CATEGORY"
	}

	if search != "" {
		query = query + " AND (LOWER(title) LIKE $SEARCH OR LOWER(description) LIKE $SEARCH) "
	}

	// sorting
	query = query + " ORDER BY created_at " + sort

	// dynamic param mapping
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
		return nil, err
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
			return nil, err
		}
		todos = append(todos, t)
	}

	return todos, nil
}

func (s *TodoService) GetById(id int) (*models.Todo, error) {
	var t models.Todo

	query := `SELECT id, title, description, category, isCompleted, created_at, updated_at, deleted_at
			  FROM todos WHERE id = $1 AND deleted_at IS NULL`

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
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	return &t, nil
}

func (s *TodoService) CreateTodo(todo *models.Todo) error {
	query := `INSERT INTO todos (title, description, category, isCompleted, created_at, updated_at)
			  VALUES ($1, $2, $3, $4, NOW(), NOW())
			  RETURNING id, created_at, updated_at`

	return config.DB.QueryRow(
		query,
		todo.Title,
		todo.Description,
		todo.Category,
		todo.IsCompleted,
	).Scan(&todo.ID, &todo.CreatedAt, &todo.UpdatedAt)
}
