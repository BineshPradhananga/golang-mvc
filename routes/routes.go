package routes

import (
	"github.com/binesh/gomvc/controllers"
	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App) {
	//function name should be start with  capital letter
	app.Get("/list", controllers.PostIndex)
	app.Get("/create", controllers.PostCreate)

	app.Get("/test", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
}
