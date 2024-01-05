package user

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Notes struct {
	ID primitive.ObjectID `bson:"_id" json:"id,omitempty"`
}
