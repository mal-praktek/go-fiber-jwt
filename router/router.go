package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	handler2 "github.com/tegarsubkhan236/go-fiber-project/src/handler"
	"github.com/tegarsubkhan236/go-fiber-project/src/middleware"
)

func SetupRoutes(app *fiber.App) {
	// middleware
	api := app.Group("/api", logger.New())
	api.Get("/", handler2.Hello)

	// auth
	auth := app.Group("/auth")
	auth.Post("/login", handler2.Login)

	// user
	user := app.Group("/user")
	user.Get("/list/:id", handler2.GetUser)
	user.Post("/", handler2.CreateUser)
	user.Put("/:id", hallo)
	user.Delete("/:id", hallo)

	// product
	product := app.Group("/product", middleware.Protected())
	product.Get("/", handler2.GetAllProducts)
	product.Get("/:id", handler2.GetProduct)
	product.Post("/", handler2.CreateProduct)
	product.Delete("/:id", handler2.DeleteProduct)
}

func hallo(c *fiber.Ctx) error {
	return c.SendString("Hallo World")
}
