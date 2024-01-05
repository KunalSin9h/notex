package data

import (
	"time"

	"github.com/kunalsin9h/notex/user"
)

type MongoDBTestRepository struct{}

func init() {
	// Check if MongoDBTestRepository fully implements the Repository interface
	var _ Repository = (*MongoDBTestRepository)(nil)
}

// Mock DB Operations for Testing

func (db *MongoDBTestRepository) InsertNewUser(u *user.User) error {
	return nil
}

func (db *MongoDBTestRepository) FindUser(username string) (*user.User, error) {
	return &user.User{
		Username:     "dummy",
		PasswordHash: "dummy",
		Email:        "dummy",
	}, nil
}

func (db *MongoDBTestRepository) FindUserByID(id string) (*user.User, error) {
	return nil, nil
}

func (db *MongoDBTestRepository) AddUserSession(accessToken, userID string, expiresTime time.Time) error {
	return nil
}

func (db *MongoDBTestRepository) VerifySession(accessToken string) (string, error) {
	return "", nil
}

func (db *MongoDBTestRepository) InsertNewNotes(n *user.Notes) error {
	return nil
}

func (db *MongoDBTestRepository) AddNotesAccess(userID, notesID string) error {
	return nil
}

func (db *MongoDBTestRepository) GetAllNotes(userID string) ([]*user.Notes, error) {
	return nil, nil
}

func (db *MongoDBTestRepository) GetNotesByID(id string) (*user.Notes, error) {
	return nil, nil
}

func (db *MongoDBTestRepository) UpdateNotes(notes *user.Notes) error {
	return nil
}

func (db *MongoDBTestRepository) DeleteNotes(notesID, userID string) error {
	return nil
}

func (db *MongoDBTestRepository) ShareNotes(notesID, userID string, usersToShare []string) error {
	return nil
}
