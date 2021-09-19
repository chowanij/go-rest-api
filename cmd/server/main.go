package main

import (
	"fmt"
	"net/http"

	"github.com/chowanij/go-rest-api/internal/comment"
	"github.com/chowanij/go-rest-api/internal/database"
	transportHttp "github.com/chowanij/go-rest-api/internal/transport/http"
)

// App - structure for storing envs and settings
type App struct{}

// Run - app setup
func (app *App) Run() error {
	fmt.Println("Setting up out application")

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
		fmt.Println("Failed to set up server")
		return err
	}

	return nil
}

func main() {
	fmt.Println("GOLANG service")

	app := App{}
	if err := app.Run(); err != nil {
		fmt.Println("error starting up our app")
		fmt.Println(err)
	}
}
