package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tegarsubkhan236/go-fiber-project/database"
	"github.com/tegarsubkhan236/go-fiber-project/router"
	"log"
)

func main() {
	app := fiber.New()
	database.ConnectDB()

	router.SetupRoutes(app)

	log.Fatal(app.Listen(":4000"))
}
