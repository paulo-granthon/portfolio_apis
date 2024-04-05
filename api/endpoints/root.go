package endpoints

import (
	"net/http"
	"server"
)

func RootEndpoints() []server.Endpoint {
	return []server.Endpoint{
		{
			Path: "/",
			Methods: []server.Method{
				server.NewMethod("GET", GetRoot),
			},
		},
	}
}

func GetRoot(s server.Server, w http.ResponseWriter, r *http.Request) error {
	return server.WriteJSON(
		w, http.StatusOK,
		"Hello, world!",
	)
}
