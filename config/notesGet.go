package config

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

// Get godoc
//
//	@summary	Get a list of all notes for the authenticated user
//	@tags		notes
//	@Security	ApiKeyAuth
//	@success	200 {object} APIResponse
//	@success	500 {object} APIResponse
//	@router		/notes [get]
func (app *Config) Get(c *fiber.Ctx) error {
	userID := c.Locals("userID").(string)
	data, err := app.Repo.GetAllNotes(userID)
	if err != nil {
		return SendError(c, http.StatusInternalServerError, err)
	}

	return c.Status(http.StatusOK).JSON(APIResponse{
		Message: "Successfully get all the notes which the user have access",
		Data:    data,
	})
}
