package config

import "github.com/gofiber/fiber/v2"

// Share godoc
//
//	@summary	Share a note with another user for the authenticated user
//	@tags		notes
//	@Security	ApiKeyAuth
//	@success	200
//	@router		/notes/{id}/share [post]
func (app *Config) Share(c *fiber.Ctx) error {
	return nil
}
