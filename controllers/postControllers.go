package controllers

import "github.com/gofiber/fiber/v2"

// function first letter must be capital
func PostIndex(c *fiber.Ctx) error {

	return c.SendString("Hello, World! this is testing")
}
