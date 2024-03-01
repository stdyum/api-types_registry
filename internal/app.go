package internal

import (
	"github.com/stdyum/api-common/server"
)

func App() server.Routes {
	routes, _, err := Configure()
	if err != nil {
		return server.Routes{Error: err}
	}

	return routes
}
