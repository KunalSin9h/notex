package tests

import (
	"net/http"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
)

// Fake mock data for testing signup and login
var testUser = map[string]string{
	"username": gofakeit.Username(),
	"email":    gofakeit.Email(),
	"password": gofakeit.Password(true, true, true, true, false, 10),
}

func TestSignUpSuccess(t *testing.T) {
	buffer := getBufferJsonData(testUser)

	req, err := http.NewRequest(http.MethodPost, "/api/auth/signup", buffer)
	panicIfError(err)

	req.Header.Set("Content-Type", "application/json")

	res, err := router.Test(req, -1)
	panicIfError(err)

	if res.StatusCode != http.StatusOK {
		t.Fatalf("Failed to signup got status code %d", res.StatusCode)
	}
}

// Test for SignUp Bad Request
func TestSignUpFailedBadRequest(t *testing.T) {
	// Fake mock data for testing signup
	payload := map[string]string{
		"username": "", // username is empty it must send bad request
		"email":    gofakeit.Email(),
		"password": gofakeit.Password(true, true, true, true, false, 10),
	}

	buffer := getBufferJsonData(payload)

	req, err := http.NewRequest(http.MethodPost, "/api/auth/signup", buffer)
	panicIfError(err)

	req.Header.Set("Content-Type", "application/json")

	res, err := router.Test(req, -1)
	panicIfError(err)

	if res.StatusCode != http.StatusBadRequest {
		t.Fatalf("Failed to signup got status code %d", res.StatusCode)
	}
}
