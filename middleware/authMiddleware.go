package middleware

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var Validator = validator.New()

type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func ValidateEmail(email string) bool {
	if err := Validator.Var(email, "required,email"); err != nil {
		return false
	}
	return true
}

func EmailValidationMiddleware() func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		email := c.FormValue("email")

		if !ValidateEmail(email) {
			validationError := ValidationError{
				Field:   "email",
				Message: "Invalid email address.",
			}
			return c.Status(fiber.StatusBadRequest).JSON(validationError)
		}

		return c.Next()
	}
}

func ValidateUsername(username string) bool {
	if err := Validator.Var(username, "required,alphanum,min=3,max=20"); err != nil {
		return false
	}
	return true
}

func UsernameValidationMiddleware() func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		username := c.FormValue("username")

		if !ValidateUsername(username) {
			validationError := ValidationError{
				Field:   "username",
				Message: "Invalid username. Username must be alphanumeric and have a length between 3 and 20 characters.",
			}
			return c.Status(fiber.StatusBadRequest).JSON(validationError)
		}

		return c.Next()
	}
}

func ValidatePassword(password string) bool {
	if len(password) < 8 {
		return false
	}
	return true
}

func PasswordValidationMiddleware() func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		password := c.FormValue("password")

		if !ValidatePassword(password) {
			validationError := ValidationError{
				Field:   "password",
				Message: "Invalid password. Password must be at least 8 characters long.",
			}
			return c.Status(fiber.StatusBadRequest).JSON(validationError)
		}

		return c.Next()
	}
}

// func UniqueUsernameEmail() func(*fiber.Ctx) error {
// 	return func(c *fiber.Ctx) error {
// 		// Parse request body to get username and email
// 		var data map[string]string
// 		if err := c.BodyParser(&data); err != nil {
// 			return fiber.ErrBadRequest
// 		}

// 		// Check if username is unique
// 		existingUser := database.DB.Db.Where("username = ?", data["username"]).First(&model.User{})
// 		if existingUser.RowsAffected != 0 {
// 			return c.Status(fiber.StatusConflict).JSON(fiber.Map{
// 				"message": "Username already exists",
// 			})
// 		}

// 		// Check if email is unique
// 		existingUser = database.DB.Db.Where("email = ?", data["email"]).First(&model.User{})
// 		if existingUser.RowsAffected != 0 {
// 			return c.Status(fiber.StatusConflict).JSON(fiber.Map{
// 				"message": "Email already exists",
// 			})
// 		}

// 		return c.Next()
// 	}
// }
