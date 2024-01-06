package main

import (
	"github.com/kunalsin9h/notex/config"
	"github.com/kunalsin9h/notex/data"
	_ "github.com/kunalsin9h/notex/docs"
)

// @title						Notex API
// @version					1.0
// @description				Yet another note taking app (for speer.io assignment)
// @host						localhost:7000
// @basePath					/api
// @schemes					http
// @securityDefinitions.basic	BasicAuth
// @securityDefinitions.apikey	ApiKeyAuth
// @in							header
// @name						X-API-Key
func main() {
	// MongoDB connection string and application port can also be
	// taken from some sort of config file or environment variable
	// to make application simple they are hard coded here
	mongoDBConnectionString := "mongodb://localhost:27017" // !

	app := &config.Config{
		Repo: data.NewMongoDB(mongoDBConnectionString, "notex"),
		Port: 7000, // !
	}

	app.Run()
}
