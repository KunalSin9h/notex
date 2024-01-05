package data

import (
	"context"
	"fmt"
	"time"

	"github.com/kunalsin9h/notex/user"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Add a new Session
func (db *MongoDBRepository) AddUserSession(accessToken, userID string, expiresTime time.Time) error {
	session := user.Session{
		Token:          accessToken,
		UserID:         userID,
		ExpirationTime: expiresTime,
		ID:             primitive.NewObjectID(),
	}

	_, err := db.SessionTokens.InsertOne(context.Background(), session)
	return err
}

// Check is user is in session
func (db *MongoDBRepository) VerifySession(accessToken string) (string, error) {
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
