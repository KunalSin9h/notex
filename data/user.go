package data

import (
	"context"

	"github.com/kunalsin9h/notex/user"
	"go.mongodb.org/mongo-driver/bson"
)

func (db *MongoDBRepository) InsertNewUser(user *user.User) error {
	_, err := db.Users.InsertOne(context.Background(), user)
	return err
}

// Check if user exists in the database
func (db *MongoDBRepository) FindUser(username string) (*user.User, error) {
	user := user.User{}

	err := db.Users.FindOne(context.Background(), bson.D{
		{Key: "username", Value: username},
	}).Decode(&user)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (db *MongoDBRepository) FindUserByID(id string) (*user.User, error) {
	user := user.User{}

	err := db.Users.FindOne(context.Background(), bson.D{
		{Key: "id", Value: id},
	}).Decode(&user)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

// add a new notes as access
// notes are independent objects units with IDs,
// each user will have access to notes if they have a notes ID	in their NotesAccess data field
// this function add a notes access to the user with given userID
func (db *MongoDBRepository) AddNotesAccess(userID, notesID string) error {
	userConcern, err := db.FindUserByID(userID)
	if err != nil {
		return err
	}

	userConcern.NotesAccess = append(userConcern.NotesAccess, notesID)

	return db.UpdateUser(userConcern)
}

// This function will replace the user with new data on same userID
// helpful to update some property of document
// but replacing entire document just to update some property will not be the best option
// to make the app simple i am doing this Unoptimized solution
func (db *MongoDBRepository) UpdateUser(user *user.User) error {
	filter := bson.D{{Key: "id", Value: user.ID}}

	_, err := db.Users.ReplaceOne(context.Background(), filter, user)
	return err
}

func (db *MongoDBRepository) GetAllNotes(userID string) ([]*user.Notes, error) {
	userConcern, err := db.FindUserByID(userID)
	if err != nil {
		return nil, err
	}

	var res []*user.Notes

	for _, notesID := range userConcern.NotesAccess {
		notes, err := db.GetNotesByID(notesID)
		if err != nil {
			return nil, err
		}

		res = append(res, notes)
	}

	return res, nil
}
