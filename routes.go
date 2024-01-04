package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
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
	return routes
}
