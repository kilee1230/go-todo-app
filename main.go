package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"

	router "github.com/kilee1230/go-todo-app/routes"
)

func main() {
	app := fiber.New()

	var port int = 3000

	log.Info(fmt.Sprintf("🚀 Server running at http://localhost:%d", port))

	router.SetupRoutes(app)

	log.Fatal(app.Listen(fmt.Sprintf(":%d", port)))
}
