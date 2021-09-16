package main

import (
	"fmt"
	"net/http"

	transportHttp "github.com/chowanij/go-rest-api/internal/transport/http"
)

type App struct{}

func (app *App) Run() error {
	fmt.Println("Setting up out application")

	handler := transportHttp.NewHandler()
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
