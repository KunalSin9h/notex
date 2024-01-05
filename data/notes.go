package data

import (
	"context"

	"github.com/kunalsin9h/notex/user"
	"go.mongodb.org/mongo-driver/bson"
)

// Insert new notes
func (db *MongoDBRepository) InsertNewNotes(n *user.Notes) error {
	_, err := db.Notes.InsertOne(context.Background(), n)
	return err
}

func (db *MongoDBRepository) GetAllNotes(userID string) ([]*user.Notes, error) {
	userConcern, err := db.FindUserByID(userID)
	if err != nil {
		return nil, err
	}

	var res []*user.Notes

	for _, notesID := range userConcern.NotesAccess {
		notes, err := db.GetNotesByID(notesID, userID)
		if err != nil {
			return nil, err
		}

		res = append(res, notes)
	}

	return res, nil
}

func (db *MongoDBRepository) GetNotesByID(id, userID string) (*user.Notes, error) {
	notes := user.Notes{}

	err := db.Notes.FindOne(context.Background(), bson.D{
		{Key: "id", Value: id},
	}).Decode(&notes)

	if err != nil {
		return nil, err
	}

	return &notes, nil
}

// Update notes just like update user function is not 100% optimized
// we can update each property, instead to make app simple i am
// replacing the entire document.
func (db *MongoDBRepository) UpdateNotes(notes *user.Notes) error {
	filter := bson.D{{Key: "id", Value: notes.ID}}

	_, err := db.Notes.ReplaceOne(context.Background(), filter, notes)
	return err
}

func (db *MongoDBRepository) DeleteNotes(notesID, userID string) error {
	filter := bson.D{
		{Key: "id", Value: notesID},
		{Key: "authorid", Value: userID},
	}

	_, err := db.Notes.DeleteOne(context.Background(), filter)

	return err
}

func (db *MongoDBRepository) ShareNotes(notesID, userID string, usersToShare []string) error {
	// This verifies that user owns the notes, i.e current logged in user is the
	// author of the notes
	_, err := db.GetNotesByID(notesID, userID)
	if err != nil {
		return err
	}

	// giving access to all the users specified
	for _, username := range usersToShare {
		if err := db.AddNotesAccess(username, notesID); err != nil {
			return err
		}
	}

	return nil
}
