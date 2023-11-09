package controllers
import (
	// "github.com/binesh/gomvc/helpers"
	"github.com/gofiber/fiber/v2"
)

func PermissionPage(c *fiber.Ctx) error {
	return c.Render("permission", fiber.Map{
		"Title": "Permission",
	})
}