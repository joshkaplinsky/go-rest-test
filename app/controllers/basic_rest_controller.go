package controllers

import (
	"github.com/gofiber/fiber/v2"
	"log"
)

func Delete(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusOK)
}

func Get(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).SendString("Hello, world.\n")
}

func Patch(c *fiber.Ctx) error {
	log.Printf("%q", c.Body())
	return c.SendStatus(fiber.StatusOK)
}

func Post(c *fiber.Ctx) error {
	log.Printf("%q", c.Body())
	return c.SendStatus(fiber.StatusOK)
}

func Put(c *fiber.Ctx) error {
	log.Printf("%q", c.Body())
	return c.SendStatus(fiber.StatusOK)
}
