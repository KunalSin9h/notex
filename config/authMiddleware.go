package config

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (app *Config) VerifyUser(c *fiber.Ctx) error {
	// See if Access Token is in Cookies (for browser based client only)
	accessToken := c.Cookies("accessToken", "")

	// see if token is in X-API-Key: {} header for non-browser based clients, like CLI, curl, Mobile apps etc
	authHeader := c.Request().Header.Peek("X-API-Key")

	apiToken := string(authHeader)

	var token string
	if accessToken != "" {
		token = accessToken
	}
	if apiToken != "" {
		token = apiToken
	}

	// userID is the of current logged in user
	userID, err := app.Repo.VerifySession(token)

	if token == "" || err != nil {
		return SendError(c, http.StatusUnauthorized, fmt.Errorf("missing api token"))
	}

	// Setting the logged in user's userID in the context,
	// co can be used ty handlers
	c.Locals("userID", userID)
	return c.Next()
}
