package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/oykugokcek/ToDoApp/model"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *fiber.Ctx) error {
	var data map[string]string

	err := c.BodyParser(&data)

	if err != nil {
		return err
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)

	user := model.User{
		Username: data["username"],
		Email:    data["email"],
		Password: password,
	}

	return c.JSON(user)
}
