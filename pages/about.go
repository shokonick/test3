package pages

import (
	"github.com/gofiber/fiber/v2"
)

func HandleAbout(c *fiber.Ctx) error {
	return c.Render("about", fiber.Map{})
}
