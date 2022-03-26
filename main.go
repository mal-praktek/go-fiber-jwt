package main

import (
	"example/api/router"
	"example/pkg/config"
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {
	app := fiber.New()
	config.ConnectDB()

	router.SetupRoutes(app)

	log.Fatal(app.Listen(":4000"))
}
