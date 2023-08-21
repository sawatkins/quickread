package handlers

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

// NotFound returns custom 404 page
func NotFound(c *fiber.Ctx) error {
	return c.Status(404).SendFile("./static/private/404.html")
}

func Index(sessionStore *session.Store) fiber.Handler {
	return func(c *fiber.Ctx) error {
		session, _ := sessionStore.Get(c)
		session.Set("name", "john")
		log.Println(session)
		log.Println(session.Get("name"))
		session.Save()
		return c.Render("index", fiber.Map{
			"Title": "Hello, World!",
		}, "layouts/main")
	}
}

func Doc(sessionStore *session.Store) fiber.Handler {
	return func(c *fiber.Ctx) error {
		session, _ := sessionStore.Get(c)
		log.Println(session)
		log.Println(session.Get("name"))
		session.Save()
		return c.Render("doc", fiber.Map{
			"Title": "Hello, Doc!",
		}, "layouts/main")
	}
	
}

func Faq(c *fiber.Ctx) error {
	return c.Render("faq", fiber.Map{
		"Title": "Hello, FAQ!",
	}, "layouts/main")
}

func Import(c *fiber.Ctx) error {
	return c.Render("import", fiber.Map{
		"Title": "Hello, Import!",
	}, "layouts/main")
}
