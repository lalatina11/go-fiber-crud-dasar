package controllers

import (
	db2 "fiber-go/db"
	"fiber-go/db/models"
	"fiber-go/lib"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetAllTodos(ctx *fiber.Ctx) error {
	var todos []models.Todo
	var db = db2.DB()
	err := db.Find(&todos).Error
	if err != nil {
		return ctx.Status(400).JSON(lib.JsonResponse{Error: true, Data: nil, Message: err.Error()})
	}
	return ctx.Status(200).JSON(lib.JsonResponse{Error: false, Data: todos, Message: "Success to Get All Todos!"})
}

func GetTodoById(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return ctx.Status(400).JSON(lib.JsonResponse{Error: true, Data: nil, Message: "Invalid id!"})
	}
	var db = db2.DB()
	var todo *models.Todo
	err = db.Find(&todo, id).Error
	if err != nil {
		return ctx.Status(400).JSON(lib.JsonResponse{Error: true, Data: nil, Message: err.Error()})
	}
	if todo.Title == "" || todo.Description == "" {
		return ctx.Status(404).JSON(lib.JsonResponse{Error: true, Data: nil, Message: "Cannot Found Todo with id " + strconv.Itoa(id)})
	}
	return ctx.Status(200).JSON(lib.JsonResponse{Error: false, Data: todo, Message: "Success getting todo with ID " + strconv.Itoa(id)})
}

func DeleteTodo(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return ctx.Status(400).JSON(lib.JsonResponse{Error: true, Data: nil, Message: "Invalid id!"})
	}
	var db = db2.DB()
	var todo *models.Todo
	err = db.Find(&todo, id).Error
	if err != nil {
		return ctx.Status(400).JSON(lib.JsonResponse{Error: true, Data: nil, Message: err.Error()})
	}

	if todo.Title == "" || todo.Description == "" {
		return ctx.Status(404).JSON(lib.JsonResponse{Error: true, Data: nil, Message: "Cannot Found Todo with id " + strconv.Itoa(id)})
	}
	err = db.Delete(&todo, todo.ID).Error
	if err != nil {
		return ctx.Status(400).JSON(lib.JsonResponse{Error: true, Data: nil, Message: err.Error()})
	}
	return ctx.Status(200).JSON(lib.JsonResponse{Error: true, Data: nil, Message: "Success to delete todo with ID " + strconv.Itoa(id)})
}

func CreateTodo(ctx *fiber.Ctx) error {
	var todo models.Todo

	if err := ctx.BodyParser(&todo); err != nil {
		return ctx.Status(400).JSON(lib.JsonResponse{Error: true, Data: nil, Message: err.Error()})
	}
	if todo.Title == "" || todo.Description == "" {
		return ctx.Status(400).JSON(lib.JsonResponse{Error: true, Data: nil, Message: "Please insert title and description!"})
	}
	var db = db2.DB()
	err := db.Where("title = ?", todo.Title).First(&todo).Error
	if err == nil {
		return ctx.Status(400).JSON(lib.JsonResponse{Error: true, Data: nil, Message: "Title must be unique!"})
	}

	err = db.Create(&todo).Error
	if err != nil {
		return ctx.Status(400).JSON(lib.JsonResponse{Error: true, Data: nil, Message: err.Error()})
	}
	return ctx.Status(201).JSON(lib.JsonResponse{Error: false, Data: todo, Message: "Success to create new Todo!"})
}

func UpdateTodo(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return ctx.Status(400).JSON(lib.JsonResponse{Error: true, Data: nil, Message: "Invalid id!"})
	}
	var updatedTodo models.Todo
	var db = db2.DB()
	if err = ctx.BodyParser(&updatedTodo); err != nil {
		return ctx.Status(400).JSON(lib.JsonResponse{Error: true, Data: nil, Message: err.Error()})
	}
	if updatedTodo.Title == "" && updatedTodo.Description == "" {
		return ctx.Status(400).JSON(lib.JsonResponse{Error: true, Data: nil, Message: "Please insert title and description!"})
	}
	var todo models.Todo
	err = db.Find(&todo, id).Error
	if err != nil {
		return ctx.Status(400).JSON(lib.JsonResponse{Error: true, Data: nil, Message: err.Error()})
	}
	if updatedTodo.Title != "" && updatedTodo.Description == "" {
		todo.Title = updatedTodo.Title
	} else if updatedTodo.Title == "" && updatedTodo.Description != "" {
		todo.Description = updatedTodo.Description
	} else {
		todo.Title = updatedTodo.Title
		todo.Description = updatedTodo.Description
	}
	db.Save(&todo)
	return ctx.Status(200).JSON(lib.JsonResponse{Error: false, Data: todo, Message: "Success to update todo with ID " + strconv.Itoa(id)})
}
