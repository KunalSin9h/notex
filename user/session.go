package user

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Session struct {
	ID             primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	Token          string             `json:"token"`
	ExpirationTime time.Time          `json:"expirationTime"`
	UserID         string             `json:"userID"`
}
