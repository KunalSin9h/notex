package config

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

// Share godoc
//
//	@summary	Share a note with another user for the authenticated user
//	@tags		notes
//	@accept		json
//	@param		id		path	string					true	"Notes ID"
//	@param		request	body	UsersToShareNotesData	true	"usernames os users to share notes"
//	@Security	ApiKeyAuth
//	@success	200	{object}	APIResponse
//	@success	401 {object} APIResponse
//	@success	400	{object}	APIResponse
//	@router		/notes/{id}/share [post]
func (app *Config) Share(c *fiber.Ctx) error {
	currentLoggedInUser := c.Locals("userID").(string)
	notesID := c.Params("id") // notes to share

	var reqPayload UsersToShareNotesData

	if err := c.BodyParser(&reqPayload); err != nil {
		return SendError(c, http.StatusBadRequest, err)
	}

	if err := app.Repo.ShareNotes(notesID, currentLoggedInUser, reqPayload.Users); err != nil {
		return SendError(c, http.StatusBadRequest, err)
	}

	return c.Status(http.StatusOK).JSON(APIResponse{
		Message: "Successfully Shared notes to all users",
	})
}

type UsersToShareNotesData struct {
	Users []string `json:"users"`
}
