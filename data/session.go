package data

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/kunalsin9h/notex/user"
	"go.mongodb.org/mongo-driver/bson"
)

// Add a new Session
func (db *MongoDB) AddUserSession(accessToken, userID string, expiresTime time.Time) error {
	session := user.Session{
		Token:          accessToken,
		UserID:         userID,
		ExpirationTime: expiresTime,
		ID:             uuid.NewString(),
	}

	_, err := db.SessionTokens.InsertOne(context.Background(), session)
	return err
}

// Check is user is in session
func (db *MongoDB) VerifySession(accessToken string) (string, error) {
	session := user.Session{}

	err := db.SessionTokens.FindOne(context.Background(), bson.D{
		{Key: "token", Value: accessToken},
	}).Decode(&session)

	if err != nil {
		return "", err
	}

	// Check if session is expired
	// if current time is after the expirationTime then session is expired
	if time.Now().After(session.ExpirationTime) {
		return "", fmt.Errorf("session expired")
	}

	return session.UserID, nil
}
