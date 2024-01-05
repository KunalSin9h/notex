package data

import (
	"time"

	"github.com/kunalsin9h/notex/user"
)

/*
	Repository Method to make our handlers testable by mocking database.

	MongoDBRepository ----------
	                             \__ Repository
						         /
	MongoDBTestRepository-------

	MongoDBTestRepository will also define these methods but they will be dummy and return
	data without query the database, this is used for integration testing
*/

type Repository interface {
	InsertNewUser(u *user.User) error
	FindUser(u, p string) (*user.User, error)
	AddUserSession(accessToken, userID string, expiresTime time.Time) error
}
