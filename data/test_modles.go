package data

import "github.com/kunalsin9h/notex/user"

type MongoDBTestRepository struct{}

func init() {
	// Check if MongoDBTestRepository fully implements the Repository interface
	var _ Repository = (*MongoDBTestRepository)(nil)
}

func (db *MongoDBTestRepository) InsertNewUser(u *user.User) error {
	return nil
}
