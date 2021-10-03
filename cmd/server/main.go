package main

import (
	"net/http"

	"github.com/chowanij/go-rest-api/internal/comment"
	"github.com/chowanij/go-rest-api/internal/database"
	transportHttp "github.com/chowanij/go-rest-api/internal/transport/http"
	log "github.com/sirupsen/logrus"
)

// App - structure for storing application information
type App struct{
	Name string
	Version string
}

// Run - app setup
func (app *App) Run() error {
	log.Info("Setting up out application")

	db, err := database.NewDatabaseConnection()
	if err != nil {
		return err
	}

	err = database.MigrateDB(db)
	if err != nil {
		return err
	}

	commentService := comment.NewService(db)

	handler := transportHttp.NewHandler(commentService)
	handler.SetupRoutes()

	if err := http.ListenAndServe(":8080", handler.Router); err != nil {
		log.Error("Failed to set up server")
		return err
	}

	return nil
}

func main() {
	app := App{
		Name: "Commentig app",
		Version: "1.0.0",
	}

	log.SetFormatter(&log.JSONFormatter{})
	log.WithFields(
		log.Fields{
			"AppName":    app.Name,
			"AppVersion": app.Version,
		}).Info("Setting Up Our APP")
	if err := app.Run(); err != nil {
		log.Error("error starting up our app")
		log.Fatal(err)
	}
}
