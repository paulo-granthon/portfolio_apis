package endpoints

import (
	"net/http"
	"server"
)

func CreateEndpoints() []server.Endpoint {
	return mergeEndpoints(
		ProjectEndpoints(),
		RootEndpoints(),
	)
}

func mergeEndpoints(endpoints ...[]server.Endpoint) []server.Endpoint {
	var merged []server.Endpoint
	for _, e := range endpoints {
		merged = append(merged, e...)
	}
	return merged
}

func RootEndpoints() []server.Endpoint {
	return []server.Endpoint{
		{
			Path: "/",
			Methods: []server.Method{
				server.NewMethod("GET", GetRoot),
			},
		},
		{
			Path: "/migrate",
			Methods: []server.Method{
				server.NewMethod("POST", Migrate),
			},
		},
	}
}

func GetRoot(s server.Server, w http.ResponseWriter, r *http.Request) error {
	return server.WriteJSON(w, http.StatusOK, "Hello, world!")
}

func Migrate(s server.Server, w http.ResponseWriter, r *http.Request) error {
	err := s.Storage.Migrate()
	if err != nil {
		return server.SendError(w, err)
	}
	return server.WriteJSON(w, http.StatusOK, "Migrated")
}
