package routes

import (
	"github.com/binesh/gomvc/controllers"
	"github.com/binesh/gomvc/middleware"
	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App) {
	//function name should be start with  capital letter
	post := app.Group("/post", middleware.JWTValidation())
	post.Get("/list", controllers.PostIndex)
	post.Get("/create", controllers.PostCreate)
	post.Post("/show", controllers.PostShow)

	app.Get("/test", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	api := app.Group("/user")
	api.Get("/register_page", controllers.RegisterPage)
	api.Post("/register", controllers.Register)
	api.Post("/login", controllers.Login)

	api.Get("/get_user", controllers.User, middleware.JWTValidation())
	api.Post("/logout", controllers.Logout, middleware.JWTValidation())
}
