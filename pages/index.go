package pages

import (
	"os"
	"runtime"

	"codeberg.org/aryak/simplytranslate/utils"
	"github.com/gofiber/fiber/v2"
)

func HandleIndex(c *fiber.Ctx) error {
	return c.Render("index", fiber.Map{
		"host":         c.Hostname(),
		"branch":       utils.Branch,
		"version":      utils.Version(),
		"fiberversion": fiber.Version,
		"goversion":    runtime.Version(),
		"setupstatus":  os.Getenv("SIMPLYTRANSLATE_SETUP_COMPLETE"),
	})
}
