package main

import (
	"os"

	"github.com/binesh/gomvc/initalizer"
	"github.com/binesh/gomvc/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func init() {
	initalizer.LoadEnvVars()
	initalizer.ConnectToDatabase()
	initalizer.SyncDb()
}

func main() {

	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowCredentials: true, //Very important while using a HTTPonly Cookie, frontend can easily get and return back the cookie.
	}))
	routes.Routes(app)
	app.Listen(":" + os.Getenv("PORT"))

}
