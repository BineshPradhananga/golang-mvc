package main

import (
	"fmt"

	"github.com/binesh/gomvc/controllers"
	"github.com/binesh/gomvc/initalizer"
	"github.com/gofiber/fiber/v2"
)

func init() {
	initalizer.LoadEnvVars()
	initalizer.ConnectToDatabase()
	initalizer.SyncDb()
}

func main() {
	fmt.Println("hello world")

	app := fiber.New()

	app.Get("/", controllers.PostIndex)

	app.Get("/test", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Listen(":3000")
	// app.Listen(":" + os.Getenv(port))

}
