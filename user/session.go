package user

import (
	"time"
)

type Session struct {
	ID             string    `json:"id"`
	Token          string    `json:"token"`
	ExpirationTime time.Time `json:"expirationTime"`
	UserID         string    `json:"userID"`
}
