package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("go-todo-api is up and running!")
	})

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.Status(200).SendString("OK")
	})

	var port int = 3000

	log.Println(fmt.Sprintf("🚀 Server running at http://localhost:%d", port))

	log.Fatal(app.Listen(fmt.Sprintf(":%d", port)))

}
