package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()

	app.Use(
		logger.New(), // add Logger middleware
	)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("go-todo-api is up and running!")
	})

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.Status(200).SendString("OK")
	})

	var port int = 3000

	log.Info(fmt.Sprintf("🚀 Server running at http://localhost:%d", port))

	log.Fatal(app.Listen(fmt.Sprintf(":%d", port)))

}
