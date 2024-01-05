package config

import (
	"time"

	"github.com/gofiber/fiber/v2/middleware/limiter"
)

// These are defaults, the values can be changes using env vars
// to make it simple i have hard corded it
var (
	D_MAX        = 15               // nolint
	D_EXPIRATION = 30 * time.Second // nolint
)

func rateLimiterConfig() *limiter.Config {
	// Rate limiter can be ignored for local development
	/*
		To ignore rate limiter use this property in the config

		Next: func(c *fiber.Ctx) bool {
			return c.IP() == "127.0.0.1"
		},
	*/

	config := limiter.Config{
		Max:        D_MAX, // {Max} request in per {Expiration} interval
		Expiration: D_EXPIRATION,
	}

	return &config
}
