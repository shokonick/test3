package pages

import (
	"fmt"
	"runtime"

	"codeberg.org/aryak/mozhi/utils"

	"github.com/gofiber/fiber/v2"
)

func HandleIndex(c *fiber.Ctx) error {
	// Check if all required queries are present
	m := c.Queries()
	sl, _ := m["sl"]
	tl, _ := m["tl"]
	engine, _ := m["engine"]
	text, _ := m["text"]
	if sl != "" && tl != "" && engine != "" && text != "" {
		fmt.Println("Work")
	}

	return c.Render("index", fiber.Map{
		"host":         c.Hostname(),
		"version":      utils.Version(),
		"fiberversion": fiber.Version,
		"goversion":    runtime.Version(),
	})
}
