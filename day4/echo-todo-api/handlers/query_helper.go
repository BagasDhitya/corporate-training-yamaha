package handlers

import (
	"database/sql"
	"echo-todo-api/config"
)

func QueryHelper(query string) (*sql.Rows, error) {
	return config.DB.Query(query)
}
