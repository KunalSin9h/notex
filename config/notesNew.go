package config

import "github.com/gofiber/fiber/v2"

// New godoc
//
//	@summary	Create a new note for the authenticated user
//	@tags		notes
//	@Security	ApiKeyAuth
//	@success	200
//	@router		/notes [post]
func (app *Config) New(c *fiber.Ctx) error {
	return nil
}
