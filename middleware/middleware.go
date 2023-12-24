package middleware

import (
	"myapp/data"

	"github.com/DavidGQK/celeritas"
)

type Middleware struct {
	App    *celeritas.Celeritas
	Models data.Models
}
