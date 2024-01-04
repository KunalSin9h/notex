package notes

import "github.com/gofiber/fiber/v2"

// Get godoc
//
//	@summary	Get a list of all notes for the authenticated user
//	@tags		notes
//	@Security	ApiKeyAuth
//	@success	200
//	@router		/notes [get]
func Get(c *fiber.Ctx) error {
	return nil
}
