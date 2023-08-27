package pages

import (
	"codeberg.org/aryak/mozhi/utils"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"runtime"
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
