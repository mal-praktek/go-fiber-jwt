package router

import (
	"example/api/controller"
	"example/api/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRoutes(app *fiber.App) {
	// /api
	api := app.Group("/api", logger.New())
	api.Get("/", hallo)

	// /auth
	auth := app.Group("/auth")
	auth.Post("/login", controller.Login)

	// /user
	user := app.Group("/user")
	user.Get("/list/:id", controller.GetUser)
	user.Post("/", controller.CreateUser)
	user.Put("/:id", hallo)
	user.Delete("/:id", hallo)

	// /product
	product := app.Group("/product", middleware.Protected())
	product.Get("/", controller.GetAllProducts)
	product.Get("/:id", controller.GetProduct)
	product.Post("/", hallo)
	product.Delete("/:id", hallo)
}

func hallo(c *fiber.Ctx) error {
	return c.SendString("Hallo World")
}
