package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/tegarsubkhan236/go-fiber-project/handler"
	"github.com/tegarsubkhan236/go-fiber-project/middleware"
)

func SetupRoutes(app *fiber.App) {
	// middleware
	api := app.Group("/api", logger.New())
	api.Get("/", handler.Hello)

	// auth
	auth := app.Group("/auth")
	auth.Post("/login", handler.Login)

	// user
	user := app.Group("/user")
	user.Get("/list/:id", handler.GetUser)
	user.Post("/", handler.CreateUser)
	user.Put("/:id", hallo)
	user.Delete("/:id", hallo)

	// product
	product := app.Group("/product")
	product.Get("/", handler.GetAllProducts)
	product.Get("/:id", hallo)
	product.Post("/", middleware.Protected(), hallo)
	product.Delete("/:id", middleware.Protected(), hallo)
}

func hallo(c *fiber.Ctx) error {
	return c.SendString("Hallo World")
}
