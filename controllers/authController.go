package controllers

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/oykugokcek/ToDoApp/database"
	"github.com/oykugokcek/ToDoApp/model"
	"golang.org/x/crypto/bcrypt"
)

const SecretKey = "somethingsecret"

// register handler
func Register(c *fiber.Ctx) error {
	var data map[string]string

	err := c.BodyParser(&data)

	if err != nil {
		return err
	}

	// Generate a new UUID
	userID := uuid.New()

	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)

	user := model.User{
		ID:       userID, // Set the generated UUID
		Username: data["username"],
		Email:    data["email"],
		Password: password,
	}

	// Save the user record in the database
	database.DB.Db.Create(&user)

	return c.JSON(user)
}

// Login handler
func Login(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	var user model.User

	// Find the user by email
	database.DB.Db.Where("email = ?", data["email"]).First(&user)

	// Check if the user was found
	if user.ID == uuid.Nil {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "user not found",
		})
	}

	// Compare the provided password with the stored password hash
	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(data["password"])); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "incorrect password",
		})
	}

	// Create a JWT token for the user
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    user.ID.String(),                      // Convert UUID to string
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), // 1 day
	})

	// Sign the JWT token
	token, err := claims.SignedString([]byte(SecretKey))

	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "could not login",
		})
	}

	// Set the JWT token as a cookie
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "success",
	})
}

// Logout
func Logout(c *fiber.Ctx) error {
	// Clears the user's session cookie
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "success",
	})
}
