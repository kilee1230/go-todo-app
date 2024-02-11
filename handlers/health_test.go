package handler

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestHealth(t *testing.T) {

	t.Run("return 200 with OK", func(t *testing.T) {
		app := fiber.New()
		app.Get("/health", Health)

		resp, err := app.Test(httptest.NewRequest("GET", "/health", nil))
		assert.Nil(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)

		body, err := io.ReadAll(resp.Body)
		assert.Nil(t, err)
		assert.Equal(t, "OK", string(body))
	})

	t.Run("return 404 when path does not exist", func(t *testing.T) {
		app := fiber.New()

		mockTodos := func(c *fiber.Ctx) error {
			return c.Status(404).SendString("Not Found")
		}

		app.Get("/health", mockTodos)

		resp, err := app.Test(httptest.NewRequest("GET", "/nonexistent", nil))
		assert.Nil(t, err)
		assert.Equal(t, http.StatusNotFound, resp.StatusCode)
	})

	t.Run("return 500 with Internal Server Error", func(t *testing.T) {
		app := fiber.New()

		mockHealth := func(c *fiber.Ctx) error {
			return c.Status(500).SendString("Internal Server error")
		}

		app.Get("/health", mockHealth)

		resp, err := app.Test(httptest.NewRequest("GET", "/health", nil))
		assert.Nil(t, err)
		assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)

		body, err := io.ReadAll(resp.Body)
		assert.Nil(t, err)
		assert.Equal(t, "Internal Server error", string(body))
	})
}
