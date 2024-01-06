package tests

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"testing"
)

var accessToken string

// Login Require BasicAuth
func TestLoginSuccess(t *testing.T) {
	token := encodeToBasicAuth(testUser["username"], testUser["password"])

	req, err := http.NewRequest(http.MethodPost, "/api/auth/login", nil)
	panicIfError(err)

	req.Header.Set("Authorization", "Basic "+token)

	res, err := router.Test(req, -1)
	panicIfError(err)

	data := getResponseData(res)

	if res.StatusCode != http.StatusOK {
		t.Fatalf("Failed to login, got status code: %d", res.StatusCode)
	} else {
		// Setting up access token return from login to that we can use
		// it in other apis
		accessToken = data.Data.(string)
	}
}

// Login Require BasicAuth
func TestLoginBadRequest(t *testing.T) {
	token := encodeToBasicAuth(testUser["username"], "") // empty password will lead to bad request

	req, err := http.NewRequest(http.MethodPost, "/api/auth/login", nil)
	panicIfError(err)

	req.Header.Set("Authorization", "Basic "+token)

	res, err := router.Test(req, -1)
	panicIfError(err)

	// data := getResponseData(res)

	if res.StatusCode != http.StatusBadRequest {
		t.Fatalf("Failed to login, got status code: %d", res.StatusCode)
	}
}

// A function used to convert username, password to basic auth token
func encodeToBasicAuth(username, password string) string {
	userPass := fmt.Sprintf("%s:%s", username, password)

	return base64.StdEncoding.EncodeToString([]byte(userPass))

}

func authorizeRequest(req *http.Request) {
	req.Header.Set("X-API-Key", accessToken)
}
