package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/oykugokcek/ToDoApp/controllers"
	"github.com/oykugokcek/ToDoApp/handlers"
)

// SetupRoutes func
func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	auth := api.Group("/auth")

	auth.Post("/register", controllers.Register)
	auth.Post("/login", controllers.Login)
	auth.Post("/logout", controllers.Logout)

	user := api.Group("/user")

	user.Get("/", handlers.GetAllUsers)
	user.Post("/", handlers.CreateUser)
	user.Get("/:id", handlers.GetSingleUser)
	user.Put("/:id", handlers.UpdateUser)
	user.Delete("/:id", handlers.DeleteUserByID)

}
