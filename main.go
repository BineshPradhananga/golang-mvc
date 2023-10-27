package main

import (
	"os"

	"github.com/binesh/gomvc/initalizer"
	"github.com/binesh/gomvc/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/template/html/v2"
)

func init() {
	initalizer.LoadEnvVars()
	initalizer.ConnectToDatabase()
	// Perform database schema migrations
	initalizer.SyncDb()
}

func main() {
	//load template
	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})
	app.Use(cors.New(cors.Config{
		AllowCredentials: true, //Very important while using a HTTPonly Cookie, frontend can easily get and return back the cookie.
	}))
	routes.Routes(app)
	app.Listen(":" + os.Getenv("PORT"))

}
