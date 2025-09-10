package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lalatina11/go-fiber-crud-dasar/controllers"
)

func TodoRoute(route *fiber.App) {
	route.Get("/api/todos", controllers.GetAllTodos)
	route.Get("/api/todos/:id", controllers.GetTodoByID)
	route.Delete("/api/todos/:id", controllers.DeleteTodo)
	route.Post("/api/todos", controllers.CreateTodo)
	route.Patch("/api/todos/:id", controllers.UpdateTodo)
}
