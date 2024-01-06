package config

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/kunalsin9h/notex/user"
)

// Update godoc
//
//	@summary	Update an existing note by ID for the authenticated user
//	@tags		notes
//	@accept		json
//	@param		id		path	string					true	"Notes ID"
//	@param		request	body	NewNotesRequestPayload	true	"New notes request payload"
//	@Security	ApiKeyAuth
//	@success	200	{object}	APIResponse
//	@success	401 {object} APIResponse
//	@success	400	{object}	APIResponse
//	@router		/notes/{id} [put]
func (app *Config) Update(c *fiber.Ctx) error {
	userID := c.Locals("userID").(string)

	notesID := c.Params("id") // from query parameters

	var reqPayload NewNotesRequestPayload

	if err := c.BodyParser(&reqPayload); err != nil {
		return err
	}

	newNotesContent := user.Notes{
		ID:       notesID,
		Title:    reqPayload.Title,
		Body:     reqPayload.Body,
		AuthorID: userID,
	}

	if err := app.Repo.UpdateNotes(&newNotesContent); err != nil {
		return SendError(c, http.StatusBadRequest, err)
	}

	return c.Status(http.StatusOK).JSON(APIResponse{
		Message: "update notes with notesID",
	})
}
