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

// GetRoot godoc
// @Summary get root
// @Tags    root
// @Produce json
// @Success 200  {string}  string
// @Failure 500  {object}  error
// @Router  / [get]
func GetRoot(s server.Server, w http.ResponseWriter, r *http.Request) error {
	return server.WriteJSON(
		w, http.StatusOK,
		"Hello, world!",
	)
}
