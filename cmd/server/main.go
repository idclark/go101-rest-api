package main

import (
	"fmt"
	"net/http"

	"rest-api/internal/database"
	transport "rest-api/internal/transport/http"
)

type App struct{}

func (app *App) Run() error {
	fmt.Println("Setting up my app")

	var err error
	_, err = database.NewDatabase()
	if err != nil {
		return err
	}

	handler := transport.NewHandler()
	handler.SetupRoutes()

	if err := http.ListenAndServe(":8080", handler.Router); err != nil {
		return err

	}

	return nil

}

func main() {
	fmt.Println("Go Rest API!")
	app := App{}

	if err := app.Run(); err != nil {
		fmt.Println("=== Error Starting Application ===")
		fmt.Println(err)
	}
}
