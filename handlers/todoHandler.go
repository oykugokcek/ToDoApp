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
