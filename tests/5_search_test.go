package tests

import (
	"net/http"
	"testing"
)

func TestSearchNotes(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "/api/search?q=some_query", nil)
	panicIfError(err)
	authorizeRequest(req)

	res, err := router.Test(req, -1)
	panicIfError(err)

	if res.StatusCode != http.StatusOK {
		t.Fatalf("Failed to search notes: status code %d", res.StatusCode)
	}
}
