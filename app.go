package main

import (
	"context"
	"flag"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/template/html/v2"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	kagi "github.com/httpjamesm/kagigo"
	"github.com/joho/godotenv"
	"github.com/sawatkins/quickread/handlers"
	"github.com/sawatkins/quickread/models"
)

var (
	port    = flag.String("port", ":8080", "Port to listen on")
	prefork = flag.Bool("prefork", false, "Enable prefork in Production")
)

func main() {
	flag.Parse()

	err := godotenv.Load()
	if err != nil {
	   log.Printf("Error loading .env file")
	}

	// Connected with database
	//database.Connect()

	// Load AWS credentials
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-west-1"))
	if err != nil {
		log.Printf("Failed to load AWS configuration, %v", err)
	}
	s3Client := s3.NewFromConfig(cfg)
	s3PresignClient := s3.NewPresignClient(s3Client)

	// Load Kagi credentials
	kagiClient := kagi.NewClient(&kagi.ClientConfig{
		APIKey:     os.Getenv("KAGI_API_KEY"),
		APIVersion: "v0",
	})

	// Create a new engine
	engine := html.New("./views", ".html")
	engine.Reload(true) // disable in prod
	engine.Debug(true) // disable in prod

	// Create fiber app
	app := fiber.New(fiber.Config{
		Prefork: *prefork, // go run app.go -prefork
		Views:   engine,
	})

	// Create sessions
	sessionStore := session.New()
	sessionStore.RegisterType([]models.PDFDocument{})
	// app.Use(sessionStore) // what does this do? is is necessary?

	// Middleware
	app.Use(recover.New())
	app.Use(logger.New())

	// Setup static files
	app.Static("/", "./static/public")

	// Auth
	auth := setAuth()

	// Create a /api/v1 endpoint
	// v1 := app.Group("/api/v1")
	// userApis := v1.Group("/user")
	// userApis.Post("/createUser", handlers.CreateUser)

	// Routes
	app.Get("/", auth, handlers.Index(sessionStore))
	app.Get("/doc", auth, handlers.Doc(sessionStore))
	app.Get("/summarize", auth, handlers.Summarize)
	app.Get("/listen", auth, handlers.Listen)
	app.Get("/faq", auth, handlers.Faq)
	app.Get("/import", auth, handlers.Import)
	// Non-user routes
	app.Post("/upload-doc", handlers.UploadDoc(s3Client, s3PresignClient))
	app.Get("/summarize-doc", handlers.SummarizeDoc(s3PresignClient, kagiClient))

	// Handle not founds
	app.Use(handlers.NotFound)

	log.Println("Server starting on port", *port)

	// Listen on port 8080
	log.Fatal(app.Listen(*port)) // go run app.go -port=:8080
}

func setAuth() func(*fiber.Ctx) error {
	config := basicauth.Config{
		Users: map[string]string{
			"john": "doee",
		},
		Realm: "super sercret realm",
	}
	auth := basicauth.New(config)
	return auth
}
