package auth

import "github.com/gofiber/fiber/v2"

// Login godoc
//
//	@summary	Login user by using BasicAuth
//	@tags		auth
//	@Security	BasicAuth
//	@success	200
//	@router		/auth/login [post]
func Login(c *fiber.Ctx) error {
	return nil
}
