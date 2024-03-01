package endpoints

import (
	"net/http"
	"server"
)

func CreateEndpoints() []server.Endpoint {
	return append(
		ProjectEndpoints(),
		RootEndpoint(),
	)
}

func mergeEndpoints(endpoints ...[]server.Endpoint) []server.Endpoint {
	var merged []server.Endpoint
	for _, e := range endpoints {
		merged = append(merged, e...)
	}
	return merged
}

func RootEndpoint() server.Endpoint {
	return server.Endpoint{
		Path: "/",
		Methods: []server.Method{
			server.NewMethod("GET", GetRoot),
		},
	}
}

func GetRoot(s server.Server, w http.ResponseWriter, r *http.Request) error {
	return server.WriteJSON(w, http.StatusOK, "Hello, world!")
}
