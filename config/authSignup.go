package config

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/kunalsin9h/notex/user"
)

// SignUp godoc
//
//	@summary	Create a new user
//	@tags		auth
//	@accept		json
//	@param		request	body	SignUpUserPayload	true	"Signup Request body"
//	@success	200 {object}    APIResponse
//	@failure	400	{object}	APIResponse
//	@failure	500	{object}	APIResponse
//	@router		/auth/signup [post]
func (app *Config) SignUp(c *fiber.Ctx) error {
	var reqBody SignUpUserPayload

	if err := c.BodyParser(&reqBody); err != nil {
		return SendError(c, http.StatusBadRequest, err)
	}

	var err error

	// Creating a new User
	newUser := user.User{}

	newUser.ID = uuid.New()

	newUser.Username, err = user.ParseUsername(reqBody.Username)
	if err != nil {
		return SendError(c, http.StatusBadRequest, err)
	}

	newUser.Email, err = user.ParseEmail(reqBody.Email)
	if err != nil {
		return SendError(c, http.StatusBadRequest, err)
	}

	newUser.PasswordHash, err = user.ParsePassword(reqBody.Password)
	if err != nil {
		return SendError(c, http.StatusBadRequest, err)
	}

	if err := newUser.HashPassword(string(newUser.PasswordHash)); err != nil {
		return SendError(c, http.StatusInternalServerError, err)
	}

	if err := app.Repo.InsertNewUser(&newUser); err != nil {
		return SendError(c, http.StatusBadRequest, err)
	}

	return nil
}

type SignUpUserPayload struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
