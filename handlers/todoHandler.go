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

func GetTodoById(c *fiber.Ctx) error {
	// Get the TODO ID from the URL parameters
	todoID := c.Params("id")

	// Retrieve the TODO item from the database by its ID
	db := database.DB.Db
	var todo models.Todo
	result := db.First(&todo, todoID)

	if result.Error != nil {
		// Handle the case where the TODO item is not found
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "TODO not found", "data": result.Error})
	}

	// Return the TODO item as a JSON response
	return c.JSON(fiber.Map{"status": "success", "message": "TODO found", "data": todo})
}

func UpdateTodo(c *fiber.Ctx) error {
	// Get the TODO ID from the URL parameters
	todoID := c.Params("id")

	// Parse the request body into the 'updatedTodo' struct
	updatedTodo := new(models.Todo)
	if err := c.BodyParser(updatedTodo); err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Check your input", "data": err})
	}

	// Retrieve the TODO item from the database by its ID
	db := database.DB.Db
	var todo models.Todo
	result := db.First(&todo, todoID)

	if result.Error != nil {
		// Handle the case where the TODO item is not found
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "TODO not found", "data": result.Error})
	}

	// Update the TODO item with the new data
	todo.Title = updatedTodo.Title
	todo.Completed = updatedTodo.Completed

	// Save the updated TODO item in the database
	db.Save(&todo)

	// Return the updated TODO item as a JSON response
	return c.JSON(fiber.Map{"status": "success", "message": "TODO updated", "data": todo})
}

func DeleteTodo(c *fiber.Ctx) error {
	// Get the TODO ID from the URL parameters
	todoID := c.Params("id")

	// Retrieve the TODO item from the database by its ID
	db := database.DB.Db
	var todo models.Todo
	result := db.First(&todo, todoID)

	if result.Error != nil {
		// Handle the case where the TODO item is not found
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "TODO not found", "data": result.Error})
	}

	// Delete the TODO item from the database
	db.Delete(&todo)

	// Return a success response
	return c.JSON(fiber.Map{"status": "success", "message": "TODO deleted", "data": nil})
}
