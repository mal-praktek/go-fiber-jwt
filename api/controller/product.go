package controller

import (
	"example/pkg/config"
	"example/pkg/model"
	"github.com/gofiber/fiber/v2"
)

func GetAllProducts(c *fiber.Ctx) error {
	db := config.DB
	var product model.Product
	db.Find(&product)
	return c.JSON(fiber.Map{"status": "success", "message": "All Products", "data": product})
}

func GetProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	db := config.DB
	var product model.Product
	db.Find(&product, id)
	if product.Title == "" {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "error", "message": "No product find with ID", "data": nil})
	}
	return c.JSON(fiber.Map{"status": "success", "message": "Product Found", "data": product})
}

func CreateProduct(c *fiber.Ctx) error {
	return c.SendStatus(404)
}

func DeleteProduct(c *fiber.Ctx) error {
	return c.SendStatus(404)
}
