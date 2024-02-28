package endpoints

import (
	"models"
	"net/http"
	"schemas"
	"server"
	"strconv"
)

func ProjectEndpoints() []server.Endpoint {
	return []server.Endpoint{
		{
			Path: "/projects",
			Methods: []server.Method{
				server.NewMethod("GET", GetProjects),
				server.NewMethod("POST", CreateProject),
			},
		},
		{
			Path: "/projects/{id}",
			Methods: []server.Method{
				server.NewMethod("GET", GetProject),
				server.NewMethod("PUT", UpdateProject),
				server.NewMethod("DELETE", DeleteProject),
			},
		},
	}
}

func GetProjects(s server.Server, w http.ResponseWriter, r *http.Request) error {
	projects, err := s.Storage.GetProjects()
	if err != nil {
		return server.SendError(w, err)
	}
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
	if error := server.ReadJSON(r, &request.Project); error != nil {
		return server.SendError(w, error)
	}

	project := models.NewCreateProject(
		request.Project.Name,
		request.Project.Semester,
		request.Project.Company,
		request.Project.Summary,
		request.Project.Url,
	)

	id, err := s.Storage.CreateProject(project)
	if err != nil {
		return server.SendError(w, err)
	}

	return server.WriteJSON(w, http.StatusCreated, schemas.CreateProjectResponse{Id: *id})
}

func UpdateProject(s server.Server, w http.ResponseWriter, r *http.Request) error {
	var request schemas.UpdateProjectRequest
	if err := server.ReadJSON(r, &request); err != nil {
		return server.SendError(w, err)
	}

	project := models.NewProject(
		request.Id,
		request.Project.Name,
		request.Project.Semester,
		request.Project.Company,
		request.Project.Summary,
		request.Project.Url,
	)

	if err := s.Storage.UpdateProject(project); err != nil {
		return server.SendError(w, err)
	}

	return server.WriteJSON(w, http.StatusOK, nil)
}
