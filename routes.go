package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/encryptcookie"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/swagger"
)

// Running application
func (app *Config) run() {
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
		// to prevent from this, use a persisted key (from env vars maybe)
	}))

	routes.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(routes.Stack())
	})

	//------------------
	// APIs
	//-----------------
	api := routes.Group("/api")

	//------------------
	// Auth APIs
	//------------------
	auth := api.Group("/auth")
	auth.Post("/signup")
	auth.Post("/login")

	//------------------
	// Notes APIs
	//------------------
	notes := api.Group("/notes")
	notes.Get("/")
	notes.Get("/:id")
	notes.Post("/")
	notes.Put("/:id")
	notes.Delete("/:id")
	notes.Post("/:id/share")

	//------------------
	// Search APIs
	//------------------
	search := api.Group("/search")
	search.Get("/")

	return routes
}
