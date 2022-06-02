package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	app.Post("/", func(c *fiber.Ctx) error {
		fmt.Printf("%q", c.Body())
		return c.Status(fiber.StatusOK).SendString("ok\n")
	})

	app.Listen(":3000")
}
