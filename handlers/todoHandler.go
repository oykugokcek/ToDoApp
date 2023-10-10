package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/oykugokcek/ToDoApp/database"
	"github.com/oykugokcek/ToDoApp/models"
)

func GetTodos(c *fiber.Ctx) error {
	db := database.DB.Db
	var todos []models.Todo
	db.Find(&todos)
	return c.JSON(&todos)
}

func CreateTodo(c *fiber.Ctx) error {
	db := database.DB.Db
	todo := new(models.Todo)

	// Parse the request body into the 'todo' struct
	if err := c.BodyParser(todo); err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Check your input", "data": err})
	}

	// Create the todo in the database
	if err := db.Create(todo).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not create todo", "data": err})
	}

	// Handle success case (e.g., return a 201 status code)
	return c.Status(201).JSON(fiber.Map{"status": "success", "message": "Todo created", "data": todo})
}
