package middleware

import (
	"myapp/data"

	"github.com/DavigGQK/celeritas"
)

type Middleware struct {
	App    *celeritas.Celeritas
	Models data.Models
}
