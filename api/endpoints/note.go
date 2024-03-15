package endpoints

import (
	"models"
	"net/http"
	"server"
)

func NoteEndpoints() []server.Endpoint {
	return []server.Endpoint{
		{
			Path: "/notes",
			Methods: []server.Method{
				server.NewMethod("GET", GetNotesFilter),
			},
		},
	}
}

func GetNotesFilter(s server.Server, w http.ResponseWriter, r *http.Request) error {
	noteModule, err := s.Storage.GetNoteModule()
	if err != nil {
		return server.SendError(
			w, err, http.StatusInternalServerError,
			"storage misconfiguration",
		)
	}

	skill := r.URL.Query().Get("skill")
	project := r.URL.Query().Get("project")
	user := r.URL.Query().Get("user")

	notes, err := noteModule.GetFilter(models.NoteFilter{
		Skill:   &skill,
		Project: &project,
		User:    &user,
	})
	if err != nil {
		return server.SendError(
			w, err, http.StatusInternalServerError,
			"error getting notes",
		)
	}

	return server.WriteJSON(
		w, http.StatusOK,
		notes,
	)
}
