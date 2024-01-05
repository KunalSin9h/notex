package data

import (
	"context"
	"time"

	"github.com/kunalsin9h/notex/user"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDBRepository struct {
	Conn *mongo.Client
}

func init() {
	// Check if MongoDBRepository fully implements the Repository interface
	var _ Repository = (*MongoDBRepository)(nil)
}

func NewMongoDBRepository(connectionString string) *MongoDBRepository {
	// Wait for 10 seconds until fail to connect with mongodb
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel() // when context ends cancel (close / end) the context

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connectionString))
	if err != nil {
		// Crash program if could not to mongodb
		// this assumes mongodb is up and running before we run the app
		// this can be a drawback, and can be solved by Exponential Backoff And Jitter
		// where we retry to connect do db in delayed intervals, like 1sec -> 2sec -> 4sec -> 8sec
		// this way our app waits for db to start and operational.
		panic(err)
	}

	return &MongoDBRepository{
		Conn: client,
	}
}

func (db *MongoDBRepository) InsertNewUser(user *user.User) error {
	return nil
}
