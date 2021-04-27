package main

import "fmt"
import "net/http"
import transport "rest-api/internal/transport/http"

type App struct{}

func (app *App) Run() error {
	fmt.Println("Setting up my app")
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
