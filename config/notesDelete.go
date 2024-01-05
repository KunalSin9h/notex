package config

import "github.com/gofiber/fiber/v2"

// Get godoc
//
//	@summary	Delete a note by ID for the authenticated user
//	@tags		notes
//	@Security	ApiKeyAuth
//	@success	200
//	@router		/notes/{id} [delete]
func (app *Config) Delete(c *fiber.Ctx) error {
	return nil
}