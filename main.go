package main

import (
	"fiber-go/db"
	"fiber-go/db/migrations"
	"fiber-go/lib"
	"fiber-go/routers"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	PORT := lib.EnvPort()

	db.DB()                    // =>  Check DB Connection
	migrations.TodoMigration() // => Todo Migration
	app := fiber.New()
	app.Use(logger.New(logger.Config{
		Format: "Request: ${method} ${path} ${status}\n",
	}))

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{"error": false, "data": nil, "message": "Hello World!"})
	})

	routers.TodoRoute(app)

	log.Fatal(app.Listen(":" + PORT))
}
