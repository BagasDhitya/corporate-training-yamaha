package controllers

import (
	"net/http"
	"strconv"

	"echo-todo-api/models"
	"echo-todo-api/services"

	"github.com/labstack/echo/v4"
)

type TodoController struct {
	Service *services.TodoService
}

func NewTodoController(service *services.TodoService) *TodoController {
	return &TodoController{Service: service}
}

func (ctr *TodoController) GetAllTodos(c echo.Context) error {
	search := c.QueryParam("search")
	category := c.QueryParam("category")
	sort := c.QueryParam("sort")

	todos, err := ctr.Service.GetAll(search, category, sort)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "Failed to fetch todos",
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "success",
		"data":    todos,
	})
}

func (ctr *TodoController) GetById(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)

	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Invalid ID format",
		})
	}

	todo, err := ctr.Service.GetById(id)

	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{
			"message": "Todo not found",
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "success",
		"data":    todo,
	})
}

func (ctr *TodoController) CreateTodo(c echo.Context) error {
	var input models.Todo

	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Invalid request body",
		})
	}

	if input.Title == "" {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Title is required",
		})
	}

	// call service
	err := ctr.Service.CreateTodo(&input)
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
