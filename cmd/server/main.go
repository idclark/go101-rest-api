package main

import (
	"fmt"
	"net/http"

	"rest-api/internal/comment"
	"rest-api/internal/database"
	transport "rest-api/internal/transport/http"
)

type App struct{}

func (app *App) Run() error {
	fmt.Println("Setting up my app")

	var err error
	db, err := database.NewDatabase()
	if err != nil {
		return err
	}

	commentService := comment.NewService(db)

	handler := transport.NewHandler(commentService)
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
