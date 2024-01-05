package config

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

// ByQuery godoc
//
//	@summary	Search for notes based on keywords for the authenticated user
//	@tags		search
//	@Security	ApiKeyAuth
//	@Param		q	query	string	false	"keyword to search notes on"
//	@success	200 {object} APIResponse
//	@success	500 {object} APIResponse
//	@router		/search [get]
func (app *Config) ByQuery(c *fiber.Ctx) error {
	userID := c.Locals("userID").(string)

	query := c.Query("q", "")

	if query == "" {
		return SendError(c, http.StatusBadRequest, fmt.Errorf("missing query"))
	}

	data, err := app.Repo.SearchNotes(userID, query)
	if err != nil {
		return SendError(c, http.StatusInternalServerError, err)
	}

	return c.Status(http.StatusOK).JSON(APIResponse{
		Message: "Successfully get notes searched by query",
		Data:    data,
	})
}
