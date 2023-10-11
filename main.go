package main

import (
	"fmt"
	"os"

	"github.com/binesh/gomvc/initalizer"
	"github.com/binesh/gomvc/routes"
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
	// app.Static("/", "./public")

	routes.Routes(app)
	// app.Listen(":3000")
	app.Listen(":" + os.Getenv("PORT"))

}
