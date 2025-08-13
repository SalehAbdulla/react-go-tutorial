package main

import (
	"fmt"
	"log"
	"github.com/gofiber/fiber/v2"
)

type ToDo struct {
	Id        int    `json:"id"`
	Completed bool   `json:"completed"`
	Body      string `json:"body"`
}

var todos = []ToDo{
	{
		Id:        1,
		Completed: false,
		Body:      "Never Stop Pushing",
	},
}

func main() {
	app := fiber.New()

	// Simple get Request
	app.Get("/api/todos", func(c *fiber.Ctx) error {
		return c.JSON(todos)
	})

	// Post request to a new todo
	app.Post("/api/todos", func(c *fiber.Ctx) error {
		todo := &ToDo{}
		if err := c.BodyParser(todo); err != nil {
			return err
		}

		if todo.Body == "" {
			return c.Status(400).JSON(fiber.Map{"error": "Todo body is required"})
		}

		todo.Id = len(todos) + 1
		todos = append(todos, *todo)

		return c.Status(201).JSON(todo)
	})

	// Update a new todo
	app.Patch("/api/todos/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")

		for i, todo := range todos {
			if fmt.Sprint(todo.Id) == id {
				todos[i].Completed = true
				return c.Status(200).JSON(todos[i])
			}
		}
		return c.Status(404).JSON(fiber.Map{"Msg": "todo not found"})
	})
	

	// Delete todo
	app.Delete("/api/todo/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		for i, todo := range todos {
			if fmt.Sprint(todo.Id) ==  id {
				todos = append(todos[:i], todos[i+1:]...)
				return c.Status(200).JSON(fiber.Map{"success": true})
			}
		}
		return c.Status(404).JSON(fiber.Map{"Msg": "todo not found"})
	})

	log.Fatal(app.Listen(":3000"))
}
