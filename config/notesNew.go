package config

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/kunalsin9h/notex/user"
)

// New godoc
//
//	@summary	Create a new note for the authenticated user
//	@tags		notes
//	@Security	ApiKeyAuth
//	@accept		json
//	@param		request	body		NewNotesRequestPayload	true	"New notes request payload"
//	@success	200		{object}	APIResponse
//	@success	400		{object}	APIResponse
//	@success	401		{object}	APIResponse
//	@success	500		{object}	APIResponse
//	@router		/notes [post]
func (app *Config) New(c *fiber.Ctx) error {
	// The current logged in user, this variable is set at the Auth Middleware
	userID := c.Locals("userID").(string)

	var reqPayload NewNotesRequestPayload

	if err := c.BodyParser(&reqPayload); err != nil {
		return SendError(c, http.StatusBadRequest, err)
	}

	newNotes := user.Notes{
		ID:       uuid.NewString(),
		AuthorID: userID,
		Title:    reqPayload.Title,
		Body:     reqPayload.Body,
	}

	if err := app.Repo.InsertNewNotes(&newNotes); err != nil {
		return SendErrorWithMessage(c, http.StatusInternalServerError, err, "Failed to create new notes")
	}

	user, err := app.Repo.FindUserByID(userID)
	if err != nil {
		return SendError(c, http.StatusBadRequest, err)
	}

	if err := app.Repo.AddNotesAccess(string(user.Username), newNotes.ID); err != nil {
		return SendErrorWithMessage(c, http.StatusInternalServerError, err, "Failed to give access to notes")
	}

	return c.Status(http.StatusOK).JSON(APIResponse{
		Message: "New notes created",
		Data:    newNotes.ID,
	})
}

type NewNotesRequestPayload struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}
