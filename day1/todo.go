package main

import (
	"encoding/json"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

type Todo struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Category    string `json:"category"`
	IsCompleted bool   `json:"isCompleted"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

// global variable
var todos []Todo

// load json file
func LoadData() error {
	file, err := os.ReadFile("dummy.json")
	if err != nil {
		return err
	}

	err = json.Unmarshal(file, &todos)
	return err
}

// save data
func SaveData() error {
	jsonData, err := json.MarshalIndent(todos, "", "  ")

	if err != nil {
		return err
	}

	err = os.WriteFile("dummy.json", jsonData, 0644)
	return err
}

// HANDLER : GET /todos
func GetAllTodoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method not allowed"))
		return
	}

	// filter : hanya todo yang belum dihapus
	var activeTodos []Todo
	for _, t := range todos {
		activeTodos = append(activeTodos, t)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(activeTodos)
}

// HANDLER : GET /todos/:id
func GetTodoByIdHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method not allowed"))
		return
	}

	w.Header().Set("Content-Type", "application/json")

	// extract id dari URL
	// example : /todos/2
	path := r.URL.Path
	parts := strings.Split(path, "/")

	// validasi format URL
	if len(parts) < 3 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error" : "invalid request format"}`))
		return
	}

	idStr := parts[2]
	id, err := strconv.Atoi(idStr)

	// validasi format ID
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error" : "id must be a number"}`))
		return
	}

	// search todo by id
	for _, t := range todos {
		// skip deleted items
		if !t.DeletedAt.IsZero() {
			continue
		}

		if t.ID == id {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(t)
			return
		}
	}

	// jika id tidak ditemukan
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte(`{"error" : "todo not found"}`))
}

func CreateTodoHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var newTodo Todo

	// decode json body
	err := json.NewDecoder(r.Body).Decode(&newTodo)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error" : "invalid request body"}`))
		return
	}

	// validasi required fields
	if newTodo.Title == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error" : "title is required"}`))
		return
	}

	// generate new ID
	newID := 1
	if len(todos) > 0 {
		newID = todos[len(todos)-1].ID + 1
	}
	newTodo.ID = newID

	// set timestamps
	now := time.Now().Local().UTC()
	newTodo.CreatedAt = now
	newTodo.UpdatedAt = now
	newTodo.DeletedAt = time.Time{}

	// simpan ke memory
	todos = append(todos, newTodo)

	// simpan ke json
	SaveData()

	// response
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newTodo)
}

func UpdateTodoHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	path := r.URL.Path
	parts := strings.Split(path, "/")

	if len(parts) < 3 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error" : "invalid request format"}`))
	}

	idStr := parts[2]
	id, err := strconv.Atoi(idStr)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error" : "id must be a number"}`))
		return
	}

	// decode payload
	var updatedData Todo
	err = json.NewDecoder(r.Body).Decode(&updatedData)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error" : "invalid request body"}`))
		return
	}

	// find todo & update
	for i, t := range todos {
		if t.ID == id && t.DeletedAt.IsZero() {

			// hanya update jika user kirim payload
			if updatedData.Title != "" {
				todos[i].Title = updatedData.Title
			}
			if updatedData.Description != "" {
				todos[i].Description = updatedData.Description
			}
			if updatedData.Category != "" {
				todos[i].Category = updatedData.Category
			}

			// isCompleted boleh true/false -> selalu update
			todos[i].IsCompleted = updatedData.IsCompleted

			// update timestamp
			now := time.Now().Local().UTC()
			todos[i].UpdatedAt = now

			SaveData()

			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(todos[i])
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte(`{"error" : "todo not found"}`))
}

func DeleteTodoHandler(w http.ResponseWriter, r *http.Request) {

}

// -- COMBINED ROUTE --
func TodosHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		GetAllTodoHandler(w, r)
	case http.MethodPost:
		CreateTodoHandler(w, r)
	default:

		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte(`{"error" : "method not allowed"}`))
	}
}

func TodosByIdHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		GetTodoByIdHandler(w, r)
	case http.MethodPut:
		UpdateTodoHandler(w, r)
	case http.MethodDelete:
		DeleteTodoHandler(w, r)
	default:

		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte(`{"error" : "method not allowed"}`))
	}
}
