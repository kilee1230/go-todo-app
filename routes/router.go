package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"

	handler "github.com/kilee1230/go-todo-app/handlers"
)

// SetupRoutes setup router api
func SetupRoutes(app *fiber.App) {
	api := app.Group("/api", logger.New())

	// health check
	api.Get("/health", handler.Health)

	// todo routes
	todoRouter := api.Group("/todos")
	todoRouter.Get("/", handler.GetTodos)
	todoRouter.Post("/", handler.AddTodo)
	todoRouter.Get("/:id", handler.GetTodoByID)
	todoRouter.Patch("/:id", handler.UpdateTodo)
	todoRouter.Delete("/:id", handler.DeleteTodo)

}
