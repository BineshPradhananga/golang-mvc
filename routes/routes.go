package routes

import (
	"github.com/binesh/gomvc/controllers"
	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App) {
	//function name should be start with  capital letter
	app.Get("/list", controllers.PostIndex)
	app.Get("/create", controllers.PostCreate)
	app.Post("/show", controllers.PostShow)

	app.Get("/test", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	api := app.Group("/user")

	api.Get("/get_user", controllers.User)

	api.Post("/register", controllers.Register)

	api.Post("/login", controllers.Login)

	api.Post("/logout", controllers.Logout)
}
