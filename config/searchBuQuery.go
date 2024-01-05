package config

import "github.com/gofiber/fiber/v2"

// New godoc
//
//	@summary	Search for notes based on keywords for the authenticated user
//	@tags		search
//	@Security	ApiKeyAuth
//	@Param		q	query	string	false	"keyword to search notes on"
//	@success	200
//	@router		/search [get]
func (app *Config) ByQuery(c *fiber.Ctx) error {
	return nil
}
