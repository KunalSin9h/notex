package notes

import "github.com/gofiber/fiber/v2"

// GetByID godoc
//
//	@summary	Get a note by ID for the authenticated user
//	@tags		notes
//	@Security	ApiKeyAuth
//	@success	200
//	@router		/notes/{id} [get]
func GetByID(c *fiber.Ctx) error {
	return nil
}
