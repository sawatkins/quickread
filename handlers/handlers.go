package handlers

import (
	// "fmt"
	// "log"
	// "time"

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

func Summarize() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// var nextSummary string
		// if summaryTime.Unused {
		// 	nextSummary = "Summary has not been used yet!"
		// } else {
		// 	nextSummary = summaryTime.NextTime.UTC().Format("2006-01-02 15:04:05")
		// }
		return c.Render("summarize", fiber.Map{}, "layouts/main")
	}
}

func Listen(c *fiber.Ctx) error {
	return c.Render("listen", fiber.Map{
		"LastSummary": "Hello, Listen!",
	}, "layouts/main")
}
