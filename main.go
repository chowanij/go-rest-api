package main

import "fmt"

type App struct {}

func (app *App) Run() error {
	fmt.Println("Setting up out application")
	return nil
}
 
func main() {
	fmt.Println("GOLANG service")	
	app := App{}
	if err := app.Run(); err!= nil {
		fmt.Println("error starting up our app")
		fmt.Println(err)
	}
}


