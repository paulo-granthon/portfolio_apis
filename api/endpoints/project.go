package endpoints

import (
	"models"
	"net/http"
	"schemas"
	"server"
	"strconv"
)

func exampleProjects() []models.Project {
	return []models.Project{
		models.NewProject(1, "Khali", 1, "FATEC"),
		models.NewProject(2, "API2Semestre", 2, "2RP"),
		models.NewProject(3, "api3", 3, "2RP"),
	}
}

func ProjectEndpoints() []server.Endpoint {
	return []server.Endpoint{
		{
			Path: "/projects",
			Methods: []server.Method{
				{
					Method: "GET",
					Func:   GetProjects,
				},
				{
					Method: "POST",
					Func:   CreateProject,
				},
			},
		},
		{
			Path: "/projects/{id}",
			Methods: []server.Method{
				{
					Method: "GET",
					Func:   GetProject,
				},
			},
		},
	}
}

func GetProjects(w http.ResponseWriter, r *http.Request) error {
	return server.WriteJSON(w, http.StatusOK, exampleProjects())
}

func GetProject(w http.ResponseWriter, r *http.Request) error {
	idStr, error := server.GetRequestParam(r, "id")
	if error != nil {
		return server.SendError(w, error)
	}

	id, error := strconv.ParseUint(*idStr, 10, 64)
	if error != nil {
		return server.SendError(w, error)
	}

	for _, project := range exampleProjects() {
		if project.Id != id {
			continue
		}

		return server.WriteJSON(w, http.StatusOK, project)
	}

	return server.WriteJSON(w, http.StatusNotFound, server.Error{Error: "models.Project not found"})
}

func CreateProject(w http.ResponseWriter, r *http.Request) error {
	var request schemas.CreateProjectRequest
	if error := server.ReadJSON(r, &request); error != nil {
		return server.SendError(w, error)
	}

	project := models.NewProject(
		4,
		request.Name,
		request.Semester,
		request.Company,
	)

	return server.WriteJSON(w, http.StatusCreated, schemas.CreateProjectResponse{Id: project.Id})
}
