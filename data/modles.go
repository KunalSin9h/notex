package data

import (
	"context"
	"time"

	"github.com/charmbracelet/log"
	"github.com/kunalsin9h/notex/user"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDBRepository struct {
	Conn          *mongo.Client
	Users         *mongo.Collection
	SessionTokens *mongo.Collection
	Notes         *mongo.Collection
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

	db := &MongoDBRepository{
		Conn:          client,
		Users:         client.Database("notex").Collection("users"),
		SessionTokens: client.Database("notex").Collection("sessions"),
		Notes:         client.Database("notex").Collection("notes"),
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

	indexName, err := db.Users.Indexes().CreateMany(context.Background(), []mongo.IndexModel{usernameIndex, emailIndex})

	if err != nil {
		panic(err)
	} else {
		log.Info("Created db index on username and email for user", "index name", indexName)
	}

	return db
}

func (db *MongoDBRepository) InsertNewUser(user *user.User) error {
	_, err := db.Users.InsertOne(context.Background(), user)
	return err
}

// Check if user exists in the database
func (db *MongoDBRepository) FindUser(username, password string) (*user.User, error) {
	user := user.User{}

	err := db.Users.FindOne(context.Background(), bson.D{
		{Key: "username", Value: username},
	}).Decode(&user)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

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
