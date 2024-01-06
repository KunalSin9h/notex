package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"os"
	"testing"

	"github.com/charmbracelet/log"
	"github.com/gofiber/fiber/v2"
	"github.com/kunalsin9h/notex/config"
	"github.com/kunalsin9h/notex/data"
)

/*
	Setup Test
	This file will run (only once) before any other test will run.

	It must named be as `setup_test.go`

	We are going to setup the MongoDBTestRepository which mock MongoDB Database for testing
*/

var router *fiber.App // TestApp Router (Api Endpoints) used by testApp

func TestMain(m *testing.M) {
	// TestApp is using MongoDBTestRepository
	// which is a mock database for testing

	mongoDBConnectionString := "mongodb://localhost:27017"

	testAPP := config.Config{
		Repo: data.NewMongoDB(mongoDBConnectionString, "notex_testing"),
		Port: 7001,
	}

	log.Info("Created TestApp on port 7001")

	router = testAPP.Routes()

	os.Exit(m.Run())
}

func getBufferJsonData(data any) *bytes.Buffer {
	buf := new(bytes.Buffer)

	if err := json.NewEncoder(buf).Encode(data); err != nil {
		log.Error("Failed to encode data into bytes buffer")
		panic(err)
	}

	return buf
}

func getResponseData(res *http.Response) *config.APIResponse {
	defer res.Body.Close()

	var result config.APIResponse

	err := json.NewDecoder(res.Body).Decode(&result)
	panicIfError(err)

	return &result
}

func panicIfError(err error) {
	if err != nil {
		log.Error("Got error")
		panic(err)
	}
}
