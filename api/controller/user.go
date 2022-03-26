package controller

import (
	"example/api/service"
	"example/pkg/config"
	"example/pkg/model"
	"github.com/gofiber/fiber/v2"
)

func GetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	user, err := service.GetUserById(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  "error",
			"message": "No User find with ID",
			"data":    nil,
		})
	}
	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "User Found",
		"data":    user,
	})
}

//func CreateUserTest(c *fiber.Ctx) error {
//	bodyRequest := new(model.User)
//	if err := c.BodyParser(bodyRequest); err != nil {
//		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
//			"status": "error",
//			"message": "Review your input",
//			"data": err,
//		})
//	}
//	hash, err := service.HashPassword(bodyRequest.Password)
//	if err != nil {
//		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
//			"status": "error",
//			"message": "Couldn't hash password",
//			"data": err,
//		})
//	}
//	bodyRequest.Password = hash
//	user, err := service.CreateUser(bodyRequest)
//
//}

func CreateUser(c *fiber.Ctx) error {
	type NewUser struct {
		Username string `json:"username"`
		Email    string `json:"email"`
	}
	db := config.DB
	user := new(model.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}
	hash, err := service.HashPassword(user.Password)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": "Couldn't hash password", "data": err})
	}
	user.Password = hash

	if err := db.Create(&user).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": "Couldn't create user", "data": err})
	}
	newUser := NewUser{
		Email:    user.Email,
		Username: user.Username,
	}
	return c.JSON(fiber.Map{"status": "success", "message": "Created user", "data": newUser})
}
