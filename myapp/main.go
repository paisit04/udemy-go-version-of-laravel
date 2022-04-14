package main

import (
	"myapp/data"
	"myapp/handlers"
	"myapp/middleware"

	"github.com/paisit04/celeritas"
)

type application struct {
	App        *celeritas.Celeritas
	Handler    *handlers.Handlers
	Models     data.Models
	Middleware *middleware.Middleware
}

func main() {
	c := initApplication()
	c.App.ListenAndServe()
}
