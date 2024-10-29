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

// GetContributionsFilter godoc
// @Summary get contributions of an user and/or project
// @Tags    contribution
// @Produce json
// @Param   project    query     string  false  "project search by project"
// @Param   user       query     string  false  "user search by user"
// @Success 200  {array}  models.ContributionDetail
// @Failure 500  {object}  error
// @Router  /contributions [get]
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
