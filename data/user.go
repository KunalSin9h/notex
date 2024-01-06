package data

import (
	"context"

	"github.com/kunalsin9h/notex/user"
	"go.mongodb.org/mongo-driver/bson"
)

func (db *MongoDB) InsertNewUser(user *user.User) error {
	_, err := db.Users.InsertOne(context.Background(), user)
	return err
}

// Check if user exists in the database
func (db *MongoDB) FindUser(username string) (*user.User, error) {
	user := user.User{}

	err := db.Users.FindOne(context.Background(), bson.D{
		{Key: "username", Value: username},
	}).Decode(&user)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (db *MongoDB) FindUserByID(id string) (*user.User, error) {
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
func (db *MongoDB) AddNotesAccess(username, notesID string) error {
	userConcern, err := db.FindUser(username)
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
func (db *MongoDB) UpdateUser(user *user.User) error {
	filter := bson.D{{Key: "id", Value: user.ID}}

	_, err := db.Users.ReplaceOne(context.Background(), filter, user)
	return err
}
