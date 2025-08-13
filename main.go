package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
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

	// GET all todos
	app.Get("/api/todos", func(c *fiber.Ctx) error {
		return c.JSON(todos)
	})

	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error Fetching dotenv variable")
	}

	PORT := os.Getenv("PORT")

	// POST a new todo
	app.Post("/api/todos", func(c *fiber.Ctx) error {
		todo := &ToDo{}
		if err := c.BodyParser(todo); err != nil {
			fmt.Println("Parse error:", err)
			return c.Status(400).JSON(fiber.Map{"error": "Invalid JSON format"})
		}

		if todo.Body == "" {
			return c.Status(400).JSON(fiber.Map{"error": "Todo body is required"})
		}

		todo.Id = len(todos) + 1
		todos = append(todos, *todo)

		return c.Status(201).JSON(todo)
	})

	// PATCH a todo to mark completed
	app.Patch("/api/todos/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")

		for i, todo := range todos {
			if fmt.Sprint(todo.Id) == id {
				todos[i].Completed = true
				return c.Status(200).JSON(todos[i])
			}
		}
		return c.Status(404).JSON(fiber.Map{"Msg": "Todo not found"})
	})

	// DELETE a todo
	app.Delete("/api/todos/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		for i, todo := range todos {
			if fmt.Sprint(todo.Id) == id {
				todos = append(todos[:i], todos[i+1:]...)
				return c.Status(200).JSON(fiber.Map{"success": true})
			}
		}
		return c.Status(404).JSON(fiber.Map{"Msg": "Todo not found"})
	})

	log.Fatal(app.Listen(":" + PORT))
}
