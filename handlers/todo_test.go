package handler

import (
	"encoding/json"
	"io"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestGetTodos(t *testing.T) {
	todos = []Todo{
		{ID: "1", Title: "Todo 1", Status: "pending"},
		{ID: "2", Title: "Todo 2", Status: "pending"},
	}

	app := fiber.New()
	app.Get("/todos", GetTodos)

	resp, err := app.Test(httptest.NewRequest("GET", "/todos", nil))
	assert.Nil(t, err)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)

	// Read the response body
	body, _ := io.ReadAll(resp.Body)

	// Unmarshal the response body into a slice of Todo structs
	var returnedTodos []Todo
	json.Unmarshal(body, &returnedTodos)

	// Assert that the returned todos are equal to the expected todos
	assert.Equal(t, todos, returnedTodos)
}

func TestAddTodo(t *testing.T) {
	t.Run("return 400 when body cannot be parsed", func(t *testing.T) {
		app := fiber.New()
		app.Post("/todos", AddTodo)

		body := strings.NewReader("not a valid json")
		req := httptest.NewRequest("POST", "/todos", body)
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)
		assert.Nil(t, err)
		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	})

	t.Run("return 400 when title is empty", func(t *testing.T) {
		app := fiber.New()
		app.Post("/todos", AddTodo)

		body := strings.NewReader(`{"title":""}`)
		req := httptest.NewRequest("POST", "/todos", body)
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)
		assert.Nil(t, err)
		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
	})

	t.Run("return 201 when todo is successfully created", func(t *testing.T) {
		app := fiber.New()
		app.Post("/todos", AddTodo)

		body := strings.NewReader(`{"title":"Test Todo"}`)
		req := httptest.NewRequest("POST", "/todos", body)
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)
		assert.Nil(t, err)
		assert.Equal(t, fiber.StatusCreated, resp.StatusCode)

		// Read the response body
		respBody, _ := io.ReadAll(resp.Body)

		// Unmarshal the response body into a Todo struct
		var returnedTodo Todo
		json.Unmarshal(respBody, &returnedTodo)

		// Assert that the returned todo has the expected title
		assert.Equal(t, "Test Todo", returnedTodo.Title)
	})
}

func TestFindTodoByID(t *testing.T) {
	todos = []Todo{
		{ID: "1", Title: "Todo 1", Status: "pending"},
		{ID: "2", Title: "Todo 2", Status: "pending"},
	}

	t.Run("return nil when no todo found with the given ID", func(t *testing.T) {
		todo, index := FindTodoByID("3")
		assert.Nil(t, todo)
		assert.Equal(t, -1, index)
	})

	t.Run("return todo when todo found with the given ID", func(t *testing.T) {
		expectedTodo := &todos[0]
		todo, index := FindTodoByID("1")
		assert.Equal(t, expectedTodo, todo)
		assert.Equal(t, 0, index)
	})
}
