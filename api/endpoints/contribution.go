package endpoints

import (
	"net/http"
	"server"
)

func ContributionEndpoints() []server.Endpoint {
	return []server.Endpoint{
		{
			Path: "/contributions",
			Methods: []server.Method{
				server.NewMethod("GET", GetContributionsFilter),
			},
		},
	}
}

func GetContributionsFilter(s server.Server, w http.ResponseWriter, r *http.Request) error {
	project := r.URL.Query().Get("project")
	user := r.URL.Query().Get("user")

	constributionService := s.Service.ContributionService

	contributions, err := constributionService.GetFilter(
		&project,
		&user,
	)
	if err != nil {
		return server.SendError(
			w, err, http.StatusInternalServerError,
			"error getting contributions",
		)
	}

	return server.WriteJSON(
		w, http.StatusOK,
		contributions,
	)
}
