package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestGetTodos(t *testing.T) {
	app := fiber.New()
	app.Get("/todos", GetTodos)

	resp, err := app.Test(httptest.NewRequest("GET", "/todos", nil))
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestAddTodo(t *testing.T) {
	app := fiber.New()
	app.Post("/todos", AddTodo)

	mockUUID := "test-uuid"

	todo := &Todo{
		ID:    mockUUID,
		Title: "Test Todo",
	}

	todoBytes, _ := json.Marshal(todo)
	req := httptest.NewRequest("POST", "/todos", bytes.NewBuffer(todoBytes))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusCreated, resp.StatusCode)
}
