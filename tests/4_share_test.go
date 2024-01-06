package tests

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
)

func TestShareNotes(t *testing.T) {
	// create a new user
	// the current logged in user will share notes to this new user
	newTestUser := map[string]string{
		"username": gofakeit.Username(),
		"email":    gofakeit.Email(),
		"password": gofakeit.Password(true, true, true, true, false, 10),
	}

	buffer := getBufferJsonData(newTestUser)

	// Creating a new user
	req, err := http.NewRequest(http.MethodPost, "/api/auth/signup", buffer)
	panicIfError(err)

	req.Header.Set("Content-Type", "application/json")
	authorizeRequest(req)

	res, err := router.Test(req, -1)
	panicIfError(err)

	if res.StatusCode != http.StatusOK {
		t.Fatalf("Failed to signup got status code %d", res.StatusCode)
	}

	// new user created
	// now lets share the notes
	reqPayload := map[string][]string{
		"users": {
			newTestUser["username"], // username of new user
		},
	}

	fmt.Println(reqPayload)
	buffer = getBufferJsonData(reqPayload)

	req, err = http.NewRequest(http.MethodPost, fmt.Sprintf("/api/notes/%s/share", notesId), buffer)
	panicIfError(err)

	req.Header.Set("Content-Type", "application/json")
	authorizeRequest(req)

	res, err = router.Test(req, -1)
	panicIfError(err)

	if res.StatusCode != http.StatusOK {
		t.Fatalf("Failed to share notes to user,  status code %d", res.StatusCode)
	}
}

// Testing Delete notes by id
// after sharing the notes to other user in above test
func TestDeleteNotesByID(t *testing.T) {
	req, err := http.NewRequest(http.MethodDelete, "/api/notes/"+notesId, nil)
	panicIfError(err)

	authorizeRequest(req)

	res, err := router.Test(req, -1)
	panicIfError(err)

	if res.StatusCode != http.StatusOK {
		t.Fatalf("Failed to delete notes by id, got status code %d", res.StatusCode)
	}
}
