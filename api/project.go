package main

import (
	"net/http"
	"server"
	"strconv"
)

type Project struct {
	Id       uint64 `json:"id"`
	Name     string `json:"name"`
	Semester uint8  `json:"semester"`
	Company  string `json:"company"`
}

func exampleProjects() []Project {
	return []Project{
		NewProject(1, "Khali", 1, "FATEC"),
		NewProject(2, "API2Semestre", 2, "2RP"),
		NewProject(3, "api3", 3, "2RP"),
	}
}

func NewProject(id uint64, name string, semester uint8, company string) Project {
	return Project{id, name, semester, company}
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

	return server.WriteJSON(w, http.StatusNotFound, server.Error{Error: "Project not found"})
}
