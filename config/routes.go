package config

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/encryptcookie"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/swagger"
)

// Running application
func (app *Config) Run() {
	if err := app.routes().Listen(fmt.Sprintf(":%d", app.Port)); err != nil {
		panic(err)
	}
}

// Routes defile all the API Endpoints
func (app *Config) routes() *fiber.App {
	routes := fiber.New()

	// CORS settings
	routes.Use(cors.New())

	// Logger for logging every API Request
	routes.Use(logger.New())

	// Swagger 2.0 API Documentation
	routes.Get("/swagger/*", swagger.HandlerDefault)

	// Request Throttling (Rate limiting API Access)
	// Default is 15 request per 30 Second window by single IP address
	routes.Use(limiter.New(*rateLimiterConfig()))

	// Encrypted Cookies for secure Authentication from the Browser
	routes.Use(encryptcookie.New(encryptcookie.Config{
		Key: encryptcookie.GenerateKey(),
		// A new Key is generated every-time which will make current cookies invalid
		// to have same key on every run, use a persisted key (from env vars maybe)
	}))

	// Check if application is up and running
	// at GET /
	routes.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(routes.Stack())
	})

	//------------------
	// APIs
	//-----------------
	api := routes.Group("/api")

	// GET /api
	api.Get("/", func(c *fiber.Ctx) error { return c.Status(http.StatusOK).SendString("Notex is up and running\n") })

	//------------------
	// Auth APIs
	//------------------
	authAPI := api.Group("/auth")
	authAPI.Post("/signup", app.Login)
	authAPI.Post("/login", app.SignUp)

	//------------------
	// Notes APIs
	//------------------
	notesAPI := api.Group("/notes")
	notesAPI.Get("/", app.Get)
	notesAPI.Get("/:id", app.GetByID)
	notesAPI.Post("/", app.New)
	notesAPI.Put("/:id", app.Update)
	notesAPI.Delete("/:id", app.Delete)
	notesAPI.Post("/:id/share", app.Share)

	//------------------
	// Search APIs
	//------------------
	searchAPI := api.Group("/search")
	searchAPI.Get("/", app.ByQuery)

	return routes
}
