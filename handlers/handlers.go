package handlers

import (
	"github.com/gofiber/fiber/v2"
)

// NotFound returns custom 404 page
func NotFound(c *fiber.Ctx) error {
	return c.Status(404).SendFile("./static/private/404.html")
}

func Index(c *fiber.Ctx) error {
	// Render index within layouts/main
	return c.Render("index", fiber.Map{
		"Title": "Hello, World!",
	}, "layouts/main")
}

func Servers(c *fiber.Ctx) error {
	return c.Render("servers", fiber.Map{
		"Title": "Hello, Servers!",
	}, "layouts/main")
}

func Faq(c *fiber.Ctx) error {
	return c.Render("faq", fiber.Map{
		"Title": "Hello, FAQ!",
	}, "layouts/main")
}


