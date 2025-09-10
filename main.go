package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/lalatina11/go-fiber-crud-dasar/db"
	"github.com/lalatina11/go-fiber-crud-dasar/db/migrations"
	"github.com/lalatina11/go-fiber-crud-dasar/lib"
	"github.com/lalatina11/go-fiber-crud-dasar/routers"
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
