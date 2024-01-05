package config

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

// Delete godoc
//
//	@summary	Delete a note by ID for the authenticated user
//	@tags		notes
//	@param		id	path	string	true	"Notes ID"
//	@Security	ApiKeyAuth
//	@success	200 {object} APIResponse
//	@success	400 {object} APIResponse
//	@router		/notes/{id} [delete]
func (app *Config) Delete(c *fiber.Ctx) error {
	userID := c.Locals("userID").(string)

	notesID := c.Params("id")

	if err := app.Repo.DeleteNotes(notesID, userID); err != nil {
		return SendError(c, http.StatusBadRequest, err)
	}

	return c.Status(http.StatusOK).JSON(APIResponse{
		Message: "Successfully deleted notes by id",
	})
}
