package user

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Notes struct {
	ID       primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	Title    string             `json:"title"`
	Body     string             `json:"body"`
	AuthorID string             `json:"authorID"` // the author (the user who created this notes)
}
