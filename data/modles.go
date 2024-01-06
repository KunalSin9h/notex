package data

import (
	"context"
	"time"

	"github.com/charmbracelet/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
	Conn          *mongo.Client
	Users         *mongo.Collection
	SessionTokens *mongo.Collection
	Notes         *mongo.Collection
}

// dbName will depends on environment, like for testing we use notexTest
// and for production we use notex
func NewMongoDB(connectionString, dbName string) *MongoDB {
	log.Info("Connecting to mongodb database, if its not running, i will hold for 10 seconds :)")
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

	db := &MongoDB{
		Conn:          client,
		Users:         client.Database(dbName).Collection("users"),
		SessionTokens: client.Database(dbName).Collection("sessions"),
		Notes:         client.Database(dbName).Collection("notes"),
	}

	uniqueConstraintUsername := options.IndexOptions{}
	uniqueConstraintEmail := options.IndexOptions{}

	uniqueConstraintUsername.SetUnique(true)
	uniqueConstraintEmail.SetUnique(true)

	usernameIndex := mongo.IndexModel{
		Keys:    bson.D{{Key: "username", Value: 1}},
		Options: &uniqueConstraintUsername,
	}

	emailIndex := mongo.IndexModel{
		Keys:    bson.D{{Key: "email", Value: 1}},
		Options: &uniqueConstraintEmail,
	}

	accessToken := mongo.IndexModel{
		Keys: bson.D{{Key: "token", Value: 1}},
	}

	indexName, err := db.Users.Indexes().CreateMany(context.Background(), []mongo.IndexModel{usernameIndex, emailIndex})

	if err != nil {
		panic(err)
	} else {
		log.Info("Created db index on username and email for user", "index name", indexName)
	}

	sessionIndex, err := db.SessionTokens.Indexes().CreateOne(context.Background(), accessToken)
	if err != nil {
		panic(err)
	} else {
		log.Info("Created session token index", "index name", sessionIndex)
	}

	// TEXT Indexing
	// used for text searching
	textSearchIndexing := mongo.IndexModel{
		Keys: bson.D{
			{Key: "title", Value: "text"},
			{Key: "body", Value: "text"},
		},
	}

	notesIndex, err := db.Notes.Indexes().CreateOne(context.Background(), textSearchIndexing)

	if err != nil {
		panic(err)
	} else {
		log.Info("Created text searching index on notes", "index name", notesIndex)
	}

	return db
}
