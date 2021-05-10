package main

import (
	"net/http"

	"rest-api/internal/comment"
	"rest-api/internal/database"
	transport "rest-api/internal/transport/http"

	log "github.com/sirupsen/logrus"
)

type App struct {
	Name    string
	Version string
}

func (app *App) Run() error {
	log.SetFormatter(&log.JSONFormatter{})
	log.WithFields(
		log.Fields{
			"AppName":    app.Name,
			"AppVersion": app.Version,
		}).Info("Setting Up Our App")

	var err error
	db, err := database.NewDatabase()
	if err != nil {
		return err
	}

	err = database.MigrateDB(db)

	if err != nil {
		log.Error("failed to setup database")
		return err
	}

	commentService := comment.NewService(db)

	handler := transport.NewHandler(commentService)
	handler.SetupRoutes()

	if err := http.ListenAndServe(":8080", handler.Router); err != nil {
		log.Error("Failed to setup server")
		return err

	}

	return nil

}

func main() {
	log.Info("Starting Rest API!")

	app := App{
		Name:    "Comment Service",
		Version: "1.0.0",
	}

	if err := app.Run(); err != nil {
		log.Error(err)
		log.Fatal("Error starting up API")

	}
}
