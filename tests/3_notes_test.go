package tests

import (
	"net/http"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
)

func TestGetNotesSuccess(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "/api/notes", nil)
	panicIfError(err)

	authorizeRequest(req)

	res, err := router.Test(req, -1)
	panicIfError(err)

	if res.StatusCode != http.StatusOK {
		t.Fatalf("Failed to get all notes, got status code %d", res.StatusCode)
	}
}

func TestGetNotesUnauthorized(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "/api/notes", nil)
	panicIfError(err)

	// not setting access token will not allow to hit api
	// this will return authorized error

	// req.Header.Set("X-API-Key", accessToken)

	res, err := router.Test(req, -1)
	panicIfError(err)

	if res.StatusCode != http.StatusUnauthorized {
		t.Fatalf("Failed to get all notes, got status code %d", res.StatusCode)
	}
}

var notesId string

func TestCreateNewNotesSuccess(t *testing.T) {
	reqBody := map[string]string{
		"body":  gofakeit.Paragraph(3, 5, 500, " "),
		"title": gofakeit.BookTitle(),
	}

	buffer := getBufferJsonData(reqBody)

	req, err := http.NewRequest(http.MethodPost, "/api/notes", buffer)
	panicIfError(err)

	req.Header.Set("Content-Type", "application/json")
	authorizeRequest(req)

	res, err := router.Test(req, -1)
	panicIfError(err)

	if res.StatusCode != http.StatusOK {
		t.Fatalf("Failed to create new notes, got status code %d", res.StatusCode)
	} else {
		data := getResponseData(res)
		notesId = data.Data.(string)
	}
}

func TestGetNotesByID(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "/api/notes/"+notesId, nil)
	panicIfError(err)

	authorizeRequest(req)

	res, err := router.Test(req, -1)
	panicIfError(err)

	if res.StatusCode != http.StatusOK {
		t.Fatalf("Failed to get notes by id, got status code %d", res.StatusCode)
	}
}

// Delete Notes by ID will be tested after we share the notes
// so TestDeleteNotesByID will be in file 4_share_test.go
