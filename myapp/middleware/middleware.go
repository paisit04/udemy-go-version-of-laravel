package middleware

import (
	"myapp/data"

	"github.com/paisit04/celeritas"
)

type Middleware struct {
	App    *celeritas.Celeritas
	Models data.Models
}
