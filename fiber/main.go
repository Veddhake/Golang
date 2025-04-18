package main

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type Todo struct {
	Id        int    `json:"id:"`
	Name      string `json:"name"`
	Completed bool   `json:"completed"`
}

var todos = []*Todo{
	{Id: 1, Name: "Walk the dog", Completed: false},
	{Id: 2, Name: "Walk the cat", Completed: false},
}

func GetTodos(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).JSON(todos)
}

func CreateTodo(ctx *fiber.Ctx) error {
	type request struct {
		Name string `json:"name"`
	}

	var body request

	if err := ctx.BodyParser(&body); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse JSON",
		})
	}

	if body.Name == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "todo name is required",
		})
	}

	// Add the new todo to the list
	newTodo := &Todo{
		Id:        len(todos) + 1,
		Name:      body.Name,
		Completed: false,
	}
	todos = append(todos, newTodo)

	return ctx.Status(fiber.StatusCreated).JSON(newTodo)
}

func GetTodo(ctx *fiber.Ctx) error {
	// Get the 'id' parameter from the URL path, e.g., /todos/2 -> "2"
	paramsId := ctx.Params("id")

	// Convert the string ID to an integer
	id, err := strconv.Atoi(paramsId)
	if err != nil {
		// If conversion fails (e.g., input is "abc"), return a 400 Bad Request
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse id",
		})
	}

	// Loop through all todos to find a match with the given ID
	for _, todo := range todos {
		if todo.Id == id {
			// If a matching todo is found, return it with a 200 OK status
			return ctx.Status(fiber.StatusOK).JSON(todo)
		}
	}

	// If no matching todo is found, return a 404 Not Found error
	return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
		"error": "todo not found",
	})
}

func DeleteTodo(ctx *fiber.Ctx) error {
	// Get the 'id' parameter from the URL path, e.g., /todos/2 -> "2"
	paramsId := ctx.Params("id")

	// Convert the string ID to an integer
	id, err := strconv.Atoi(paramsId)
	if err != nil {
		// If conversion fails (e.g., input is "abc"), return a 400 Bad Request
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse id",
		})
	}

	for i, todo := range todos {
		if todo.Id == id {
			todos = append(todos[0:i], todos[i+1:]...)
			return ctx.SendStatus(fiber.StatusNoContent)
		}
	}
	// If no matching todo is found, return a 404 Not Found error
	return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
		"error": "todo not found",
	})
}

func UpdateTodo(ctx *fiber.Ctx) error {

	type request struct {
		Name      *string `json:"name"`
		Completed *bool   `json:"completed"`
	}

	// Get the 'id' parameter from the URL path, e.g., /todos/2 -> "2"
	paramsId := ctx.Params("id")

	// Convert the string ID to an integer
	id, err := strconv.Atoi(paramsId)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse id",
		})
	}

	var body request
	err = ctx.BodyParser(&body)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse JSON body",
		})
	}

	// Find the todo by ID
	var todo *Todo
	for _, t := range todos {
		if t.Id == id {
			todo = t
			break
		}
	}

	// If todo not found
	if todo == nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "todo not found",
		})
	}

	// Update fields if provided
	if body.Name != nil {
		todo.Name = *body.Name
	}
	if body.Completed != nil {
		todo.Completed = *body.Completed
	}

	// Return the updated todo
	return ctx.Status(fiber.StatusOK).JSON(todo)
}

func main() {

	app := fiber.New() // Used to create the application.

	app.Use(logger.New())

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("hello world")
	})

	app.Get("/todos", GetTodos)
	app.Post("/todos", CreateTodo)
	app.Get("/todos/:id", GetTodo)
	app.Delete("/todos/:id", DeleteTodo)
	app.Patch("/todos/:id", UpdateTodo)

	app.Listen(":3000") // app runnning on local port

}
