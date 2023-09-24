package handlers

import (
	// "fmt"
	// "log"

	"github.com/gofiber/fiber/v2"
	// "github.com/gofiber/fiber/v2/middleware/session"
)

// NotFound returns custom 404 page
func NotFound(c *fiber.Ctx) error {
	return c.Status(404).SendFile("./static/private/404.html")
}

func Index(c *fiber.Ctx) error {
	// return func(c *fiber.Ctx) error {
		// session, err := sessionStore.Get(c)
		// if err != nil {
		// 	log.Printf("Filed to get session store %v\n", err)
		// }
		// fmt.Println(session)
		// fmt.Println(session.Get("pdfDocuments"))
		
		return c.Render("index", fiber.Map{
			"Title": "Hello, World!",
		}, "layouts/main")
	// }
}

// func Doc(sessionStore *session.Store) fiber.Handler {
// 	return func(c *fiber.Ctx) error {
// 		session, err := sessionStore.Get(c)
// 		if err != nil {
// 			log.Printf("Filed to get session store %v\n", err)
// 		}
// 		fmt.Println(session)
// 		fmt.Println(session.Get("pdfDocuments"))
// 		return c.Render("doc", fiber.Map{
// 			"Title": "Hello, Doc!",
// 			"numDocs": 22,
// 			"pdfDocuments": session.Get("pdfDocuments"),
// 		}, "layouts/main")
// 	}
	
// }

func Faq(c *fiber.Ctx) error {
	return c.Render("faq", fiber.Map{
		"Title": "Hello, FAQ!",
	}, "layouts/main")
}

func Summarize(c *fiber.Ctx) error {
	return c.Render("summarize", fiber.Map{
		"Title": "Hello, Summarize!",
	}, "layouts/main")
}

func Listen(c *fiber.Ctx) error {
	return c.Render("listen", fiber.Map{
		"Title": "Hello, Listen!",
	}, "layouts/main")
}

