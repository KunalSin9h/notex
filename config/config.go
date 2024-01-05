package config

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kunalsin9h/notex/data"
)

/*
data.Repository is an interface which we will use to put a real MongoDB database when deploying in production
and put a mock database for testing
*/
type Config struct {
	Repo data.Repository
	Port uint16
}

type APIResponse struct {
	Message string `json:"message"`
	Error   string `json:"error"`
	Data    any    `json:"data"`
}

func SendError(c *fiber.Ctx, code int, err error) error {
	return c.Status(code).JSON(APIResponse{
		Message: "Failed to execute request",
		Error:   err.Error(),
		Data:    nil,
	})
}
