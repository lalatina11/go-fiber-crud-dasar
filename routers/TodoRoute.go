package routers

import (
	"fiber-go/controllers"

	"github.com/gofiber/fiber/v2"
)

func TodoRoute(route *fiber.App) {
	route.Get("/api/todos", controllers.GetAllTodos)
	route.Get("/api/todos/:id", controllers.GetTodoById)
	route.Delete("/api/todos/:id", controllers.DeleteTodo)
	route.Post("/api/todos", controllers.CreateTodo)
	route.Patch("/api/todos/:id", controllers.UpdateTodo)
}
