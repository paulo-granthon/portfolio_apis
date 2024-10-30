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

// GetProjects godoc
// @Summary get all projects
// @Tags    project
// @Produce json
// @Success 200  {array}  models.Project
// @Failure 500  {object}  error
// @Router  /projects [get]
func GetProjects(s server.Server, w http.ResponseWriter, r *http.Request) error {
	projectModule, err := s.Storage.GetProjectModule()
	if err != nil {
		return server.SendError(
			w, err, http.StatusInternalServerError,
			"storage misconfiguration",
		)
	}

	projects, err := projectModule.Get()
	if err != nil {
		return server.SendError(
			w, err, http.StatusInternalServerError,
			"error getting projects",
		)
	}

	return server.WriteJSON(
		w, http.StatusOK,
		projects,
	)
}

// GetProject godoc
// @Summary get a project by id
// @Tags    project
// @Produce json
// @Param   id     path     string  true  "project id"
// @Success 200  {object}  models.Project
// @Failure 404  {object}  error
// @Failure 500  {object}  error
// @Router  /projects/{id} [get]
func GetProject(s server.Server, w http.ResponseWriter, r *http.Request) error {
	idStr, err := server.GetRequestParam(r, "id")
	if err != nil {
		return server.SendError(
			w, err, http.StatusBadRequest,
			"Parameter id not found",
		)
	}

	id, err := strconv.ParseUint(*idStr, 10, 64)
	if err != nil {
		return server.SendError(
			w, err, http.StatusBadRequest,
			"Parameter id is not a valid number",
		)
	}

	projectModule, err := s.Storage.GetProjectModule()
	if err != nil {
		return server.SendError(
			w, err, http.StatusInternalServerError,
			"storage misconfiguration",
		)
	}

	project, err := projectModule.GetById(id)
	if err != nil {
		return server.SendError(
			w, err, http.StatusNotFound,
			"project not found",
		)
	}

	return server.WriteJSON(
		w, http.StatusOK,
		project,
	)
}

func CreateProject(s server.Server, w http.ResponseWriter, r *http.Request) error {
	var request schemas.CreateProjectRequest
	if err := server.ReadJSON(r, &request.Project); err != nil {
		return server.SendError(
			w, err, http.StatusBadRequest,
			"error parsing request",
		)
	}

	project := models.NewCreateProject(
		request.Project.Name,
		request.Project.Image,
		request.Project.Semester,
		request.Project.Company,
		request.Project.TeamId,
		request.Project.Summary,
		request.Project.Description,
		request.Project.Url,
	)

	projectModule, err := s.Storage.GetProjectModule()
	if err != nil {
		return server.SendError(
			w, err, http.StatusInternalServerError,
			"storage misconfiguration",
		)
	}

	id, err := projectModule.Create(project)
	if err != nil {
		return server.SendError(
			w, err, http.StatusInternalServerError,
			"error creating project",
		)
	}

	return server.WriteJSON(
		w, http.StatusCreated,
		schemas.CreateProjectResponse{Id: *id},
	)
}

// UpdateProject godoc
// @Summary update a project
// @Tags    project
// @Produce json
// @Param   id     path     string  true  "project id"
// @Param   project    body     schemas.UpdateProjectRequest  true  "project"
// @Success 200  {string}  string
// @Failure 400  {object}  error
// @Failure 404  {object}  error
// @Failure 500  {object}  error
// @Router  /projects/{id} [put]
func UpdateProject(s server.Server, w http.ResponseWriter, r *http.Request) error {
	idStr, err := server.GetRequestParam(r, "id")
	if err != nil {
		return server.SendError(
			w, err, http.StatusBadRequest,
			"Parameter id not found",
		)
	}

	id, err := strconv.ParseUint(*idStr, 10, 64)
	if err != nil {
		return server.SendError(
			w, err, http.StatusBadRequest,
			"Parameter id is not a valid number",
		)
	}

	var request schemas.UpdateProjectRequest
	if err := server.ReadJSON(r, &request); err != nil {
		return server.SendError(
			w, err, http.StatusBadRequest,
			"error parsing request",
		)
	}

	project := models.NewUpdateProject(
		id,
		&request.Project.Name,
		request.Project.Image,
		&request.Project.Semester,
		&request.Project.Company,
		&request.Project.TeamId,
		&request.Project.Summary,
		&request.Project.Description,
		&request.Project.Url,
	)

	projectModule, err := s.Storage.GetProjectModule()
	if err != nil {
		return server.SendError(
			w, err, http.StatusInternalServerError,
			"storage misconfiguration",
		)
	}

	if err := projectModule.Update(project); err != nil {
		return server.SendError(
			w, err, http.StatusInternalServerError,
			"error updating project",
		)
	}

	return server.WriteJSON(
		w, http.StatusOK,
		"project updated",
	)
}

// DeleteProject godoc
// @Summary delete a project
// @Tags    project
// @Produce json
// @Param   id     path     string  true  "project id"
// @Success 200  {string}  string
// @Failure 400  {object}  error
// @Failure 500  {object}  error
// @Router  /projects/{id} [delete]
func DeleteProject(s server.Server, w http.ResponseWriter, r *http.Request) error {
	idStr, err := server.GetRequestParam(r, "id")
	if err != nil {
		return server.SendError(
			w, err, http.StatusBadRequest,
			"Parameter id not found",
		)
	}

	id, err := strconv.ParseUint(*idStr, 10, 64)
	if err != nil {
		return server.SendError(
			w, err, http.StatusBadRequest,
			"Parameter id is not a valid number",
		)
	}

	projectModule, err := s.Storage.GetProjectModule()
	if err != nil {
		return server.SendError(
			w, err, http.StatusInternalServerError,
			"storage misconfiguration",
		)
	}

	err = projectModule.Delete(id)
	if err != nil {
		return server.SendError(
			w, err, http.StatusInternalServerError,
			"error deleting project",
		)
	}

	return server.WriteJSON(
		w, http.StatusOK,
		"project deleted",
	)
}
