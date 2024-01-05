package config

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// Login godoc
//
//	@summary	Login user by using BasicAuth
//	@tags		auth
//	@Security	BasicAuth
//	@success	200 {object} APIResponse
//	@success	400 {object} APIResponse
//	@success	500 {object} APIResponse
//	@router		/auth/login [post]
func (app *Config) Login(c *fiber.Ctx) error {
	username, password := extractFromBasic(c)

	if username == "" || password == "" {
		return SendError(c, http.StatusBadRequest, fmt.Errorf("invalid username or password"))
	}

	user, err := app.Repo.FindUser(username, password)

	if err != nil {
		return SendErrorWithMessage(c, http.StatusBadRequest, err, "Username does not exists, signup first")
	}

	// Timeing Attach Proof Password verifier
	match, err := user.VerifyPassword(password)
	if err != nil {
		return SendError(c, http.StatusInternalServerError, err)
	}
	if !match {
		return SendErrorWithMessage(c, http.StatusBadRequest, err, "Incorrect password")
	}

	accessToken := "api_" + uuid.NewString()

	// Access Token will be expired after 1 hour
	expirationTime := time.Now().Add(1 * time.Hour)

	// Asynchronously Adding user to session
	go app.Repo.AddUserSession(accessToken, user.ID.String(), expirationTime)

	// Additionally Setting up cookies with accessToken
	// so the browser client does not have to send X-Api-Key header for api key
	// we can get it from cookies
	// for non-browser clients, like CLI, Mobile applications, we can sending Access Token
	// which they can use
	c.Cookie(&fiber.Cookie{
		Name:     "accessToken",
		Value:    accessToken,
		Expires:  expirationTime,
		Secure:   true,
		SameSite: "Strict",
		HTTPOnly: true,
	})

	return c.Status(http.StatusOK).JSON(APIResponse{
		Message: "Logged in successfully",
		Data:    accessToken,
	})
}

// Helper function to extract username and password
// from basicAuth Header
func extractFromBasic(c *fiber.Ctx) (string, string) {
	value := string(c.Request().Header.Peek("Authorization"))

	token, _ := strings.CutPrefix(value, "Basic ")

	creds, err := base64.StdEncoding.DecodeString(token)
	if err != nil {
		return "", ""
	}

	credentials := string(creds)

	userPass := strings.SplitN(credentials, ":", 2)
	if len(userPass) != 2 {
		return "", ""
	}

	return userPass[0], userPass[1]
}
