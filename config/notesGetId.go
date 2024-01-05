package config

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

// GetByID godoc
//
//	@summary	Get a note by ID for the authenticated user
//	@tags		notes
//	@param		id	path	string	true	"Notes ID"
//	@Security	ApiKeyAuth
//	@success	200	{object}	APIResponse
//	@success	204	{object}	APIResponse
//	@router		/notes/{id} [get]
func (app *Config) GetByID(c *fiber.Ctx) error {
	userID := c.Locals("userID").(string)

	// NotesID is given by query parameter of the request
	notesID := c.Params("id")

	data, err := app.Repo.GetNotesByID(notesID, userID)
	if err != nil {
		return SendErrorWithMessage(c, http.StatusNoContent, err, "No notes with that id")
	}

	return c.Status(http.StatusOK).JSON(APIResponse{
		Message: "Successfully get notes by id for authenticated user",
		Data:    data,
	})
}
