package data

import (
	"context"

	"github.com/kunalsin9h/notex/user"
)

// Insert new notes
func (db *MongoDBRepository) InsertNewNotes(n *user.Notes) error {
	_, err := db.Notes.InsertOne(context.Background(), n)
	return err
}
