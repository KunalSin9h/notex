package config

import "github.com/gofiber/fiber/v2"

// Update godoc
//
//	@summary	Update an existing note by ID for the authenticated user
//	@tags		notes
//	@Security	ApiKeyAuth
//	@success	200
//	@router		/notes/{id} [put]
func (app *Config) Update(c *fiber.Ctx) error {
	return nil
}
