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
	app := fiber.New()

	app.Get("/health", Health)

	resp, err := app.Test(httptest.NewRequest("GET", "/health", nil))
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	body, err := io.ReadAll(resp.Body)
	assert.Nil(t, err)
	assert.Equal(t, "OK", string(body))
}
