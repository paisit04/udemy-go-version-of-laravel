package main

import (
	"log"
	"myapp/data"
	"myapp/handlers"
	"os"

	"github.com/paisit04/celeritas"
)

func initApplication() *application {
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	// init celeritas
	cel := &celeritas.Celeritas{}
	err = cel.New(path)
	if err != nil {
		log.Fatal(err)
	}

	cel.AppName = "myapp"

	myHandlers := &handlers.Handlers{
		App: cel,
	}

	app := &application{
		App:     cel,
		Handler: myHandlers,
	}

	app.App.Routes = app.routes()

	app.Models = data.New(app.App.DB.Pool)
	myHandlers.Models = app.Models

	return app
}
