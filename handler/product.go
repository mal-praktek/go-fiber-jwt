package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tegarsubkhan236/go-fiber-project/database"
	"github.com/tegarsubkhan236/go-fiber-project/model"
)

func GetAllProducts(c *fiber.Ctx) error {
	db := database.DB
	var product model.Product
	db.Find(&product)
	return c.JSON(fiber.Map{"status": "success", "message": "All Products", "data": product})
}

func GetProduct(c *fiber.Ctx) error {
	return c.SendStatus(404)
}

func CreateProduct(c *fiber.Ctx) error {
	return c.SendStatus(404)
}

func DeleteProduct(c *fiber.Ctx) error {
	return c.SendStatus(404)
}
