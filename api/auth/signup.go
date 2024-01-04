package auth

import "github.com/gofiber/fiber/v2"

// SignUp godoc
//
//	@summary	Create a new user
//	@tags		auth
//	@accept		json
//	@param		request	body	SignUpUserPayload	true	"Signup Request body"
//	@success	200
//	@router		/auth/signup [post]
func SignUp(c *fiber.Ctx) error {
	return nil
}

type SignUpUserPayload struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
