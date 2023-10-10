package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/oykugokcek/ToDoApp/controllers"
	"github.com/oykugokcek/ToDoApp/handler"
)

// SetupRoutes func
func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	auth := api.Group("/auth")

	auth.Post("/register", controllers.Register)
	auth.Post("/login", controllers.Login)

	v1 := api.Group("/user")
	v1.Get("/", handler.GetAllUsers)
	v1.Post("/", handler.CreateUser)
	v1.Get("/:id", handler.GetSingleUser)
	v1.Put("/:id", handler.UpdateUser)
	v1.Delete("/:id", handler.DeleteUserByID)

}
