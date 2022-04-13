package main

import (
	"myapp/handlers"

	"github.com/paisit04/celeritas"
)

type application struct {
	App     *celeritas.Celeritas
	Handler *handlers.Handlers
}

func main() {
	c := initApplication()
	c.App.ListenAndServe()
}
