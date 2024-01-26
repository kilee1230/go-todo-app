package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type Todo struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Status string `json:"status"`
}

var todos = []Todo{}

func GetTodos(c *fiber.Ctx) error {
	return c.JSON(todos)
}

func AddTodo(c *fiber.Ctx) error {
	todo := new(Todo)

	if err := c.BodyParser(todo); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	// Check if the Title field is empty
	if todo.Title == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "title is required",
		})
	}

	todo.ID = uuid.New().String() // Generate a new ID
	todo.Status = "pending"       // Set the initial status
	todos = append(todos, *todo)  // Add the todo to the list

	return c.Status(fiber.StatusCreated).JSON(todo)
}

// FindTodoByID finds a todo by ID
func FindTodoByID(id string) (*Todo, int) {
	for i, t := range todos {
		if t.ID == id {
			return &t, i
		}
	}

	return nil, -1
}

func GetTodoByID(c *fiber.Ctx) error {
	id := c.Params("id") // Get the ID from the URL parameters

	// Use the FindTodoByID function to find the todo
	todo, _ := FindTodoByID(id)
	if todo == nil {
		// If no todo with the given ID was found, return a 404 Not Found error
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "No todo found with the given ID",
		})
	}

	return c.JSON(todo) // Return the todo as a JSON response
}

func UpdateTodo(c *fiber.Ctx) error {
	id := c.Params("id") // Get the ID from the URL parameters
	todo := new(Todo)

	if err := c.BodyParser(todo); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	// Find the todo with the given ID
	t, index := FindTodoByID(id)
	if t == nil {
		// If no todo with the given ID was found, return a 404 Not Found error
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "No todo found with the given ID",
		})
	}

	// Only update the Title field if it's not empty
	if todo.Title != "" {
		todos[index].Title = todo.Title
	}

	// Only update the Status field if it's not empty
	if todo.Status != "" {
		todos[index].Status = todo.Status
	}

	return c.JSON(todos[index]) // Return the updated todo as a JSON response
}

func DeleteTodo(c *fiber.Ctx) error {
	id := c.Params("id") // Get the ID from the URL parameters

	// Find the todo with the given ID
	_, index := FindTodoByID(id)
	if index == -1 {
		// If no todo with the given ID was found, return a 404 Not Found error
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "No todo found with the given ID",
		})
	}

	// Remove the todo from the todos slice
	todos = append(todos[:index], todos[index+1:]...)

	return c.SendStatus(fiber.StatusNoContent) // Return a 204 No Content response
}
