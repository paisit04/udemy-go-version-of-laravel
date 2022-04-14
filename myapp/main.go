package main

import (
	"myapp/data"
	"myapp/handlers"

	"github.com/paisit04/celeritas"
)

type application struct {
	App     *celeritas.Celeritas
	Handler *handlers.Handlers
	Models  data.Models
}

func main() {
	c := initApplication()
	c.App.ListenAndServe()
}
