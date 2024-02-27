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

func GetProjects(s server.Server, w http.ResponseWriter, r *http.Request) error {
	projects := s.Storage.GetProjects()
	return server.WriteJSON(w, http.StatusOK, projects)
}

func GetProject(s server.Server, w http.ResponseWriter, r *http.Request) error {
	idStr, err := server.GetRequestParam(r, "id")
	if err != nil {
		return server.SendError(w, err)
	}

	id, err := strconv.ParseUint(*idStr, 10, 64)
	if err != nil {
		return server.SendError(w, err)
	}

	project, err := s.Storage.GetProject(id)
	if err != nil {
		return server.WriteJSON(w, http.StatusNotFound, server.Error{Error: "models.Project not found"})
	}

	return server.WriteJSON(w, http.StatusOK, project)
}

func CreateProject(s server.Server, w http.ResponseWriter, r *http.Request) error {
	var request schemas.CreateProjectRequest
	if error := server.ReadJSON(r, &request); error != nil {
		return server.SendError(w, error)
	}

	var project models.Project
	err := s.Storage.CreateProject(
		request.Name,
		request.Semester,
		request.Company,
	)
	if err != nil {
		return server.SendError(w, err)
	}

	return server.WriteJSON(w, http.StatusCreated, schemas.CreateProjectResponse{Id: project.Id})
}
