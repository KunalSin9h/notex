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

func (db *MongoDBTestRepository) FindUser(username, password string) (*user.User, error) {
	return &user.User{
		Username:     "dummy",
		PasswordHash: "dummy",
		Email:        "dummy",
	}, nil
}


func  (db *MongoDBTestRepository)	AddUserSession(accessToken, userID string, expiresTime time.Time) error {
	return nil
}