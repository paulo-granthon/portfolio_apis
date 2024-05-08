package endpoints

import (
	"models"
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
	contributionModule, err := s.Storage.GetContributionModule()
	if err != nil {
		return server.SendError(
			w, err, http.StatusInternalServerError,
			"storage misconfiguration",
		)
	}

	skill := r.URL.Query().Get("skill")
	project := r.URL.Query().Get("project")
	user := r.URL.Query().Get("user")

	contributions, err := contributionModule.GetFilter(models.ContributionFilter{
		Skill:   &skill,
		Project: &project,
		User:    &user,
	})
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
