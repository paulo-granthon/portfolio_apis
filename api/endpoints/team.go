package endpoints

import (
	"models"
	"net/http"
	"schemas"
	"server"
	"strconv"
)

func TeamEndpoints() []server.Endpoint {
	return []server.Endpoint{
		{
			Path: "/teams",
			Methods: []server.Method{
				server.NewMethod("GET", GetTeams),
				server.NewMethod("POST", CreateTeam),
			},
		},
		{
			Path: "/teams/{id}",
			Methods: []server.Method{
				server.NewMethod("GET", GetTeam),
				server.NewMethod("PUT", UpdateTeam),
				server.NewMethod("DELETE", DeleteTeam),
			},
		},
		{
			Path: "/teams/{id}/members",
			Methods: []server.Method{
				server.NewMethod("GET", GetTeamUsers),
				server.NewMethod("POST", AddUserToTeam),
			},
		},
		{
			Path: "/teams/{id}/members/{userId}",
			Methods: []server.Method{
				server.NewMethod("DELETE", RemoveUserFromTeam),
			},
		},
		{
			Path: "/teams/{id}/projects",
			Methods: []server.Method{
				server.NewMethod("GET", GetTeamProjects),
			},
		},
	}
}

// GetTeams godoc
// @Summary get all teams
// @Tags    team
// @Produce json
// @Success 200  {array}  models.Team
// @Failure 500  {object}  error
// @Router  /teams [get]
func GetTeams(s server.Server, w http.ResponseWriter, r *http.Request) error {
	teamModule, err := s.Storage.GetTeamModule()
	if err != nil {
		return server.SendError(
			w, err, http.StatusInternalServerError,
			"storage misconfiguration",
		)
	}

	teams, err := teamModule.Get()
	if err != nil {
		return server.SendError(
			w, err, http.StatusInternalServerError,
			"error getting teams",
		)
	}

	return server.WriteJSON(
		w, http.StatusOK,
		teams,
	)
}

// GetTeam godoc
// @Summary get team by id
// @Tags    team
// @Produce json
// @Param   id     path    int     true  "team id"
// @Success 200  {object}  models.Team
// @Failure 400  {object}  error
// @Failure 500  {object}  error
// @Router  /teams/{id} [get]
func GetTeam(s server.Server, w http.ResponseWriter, r *http.Request) error {
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
			"Invalid id",
		)
	}

	teamModule, err := s.Storage.GetTeamModule()
	if err != nil {
		return server.SendError(
			w, err, http.StatusInternalServerError,
			"storage misconfiguration",
		)
	}

	team, err := teamModule.GetById(id)
	if err != nil {
		return server.SendError(
			w, err, http.StatusInternalServerError,
			"error getting team",
		)
	}

	return server.WriteJSON(
		w, http.StatusOK,
		team,
	)
}

// GetTeamUsers godoc
// @Summary get team users
// @Tags    team
// @Produce json
// @Param   id     path    int     true  "team id"
// @Success 200  {array}  models.User
// @Failure 400  {object}  error
// @Failure 500  {object}  error
// @Router  /teams/{id}/members [get]
func GetTeamUsers(s server.Server, w http.ResponseWriter, r *http.Request) error {
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
			"Invalid id",
		)
	}

	teamModule, err := s.Storage.GetTeamModule()
	if err != nil {
		return server.SendError(
			w, err, http.StatusInternalServerError,
			"storage misconfiguration",
		)
	}

	users, err := teamModule.GetUsers(id)
	if err != nil {
		return server.SendError(
			w, err, http.StatusInternalServerError,
			"error getting team users",
		)
	}

	return server.WriteJSON(
		w, http.StatusOK,
		users,
	)
}

// GetTeamProjects godoc
// @Summary get team projects
// @Tags    team
// @Produce json
// @Param   id     path    int     true  "team id"
// @Success 200  {array}  models.Project
// @Failure 400  {object}  error
// @Failure 500  {object}  error
// @Router  /teams/{id}/projects [get]
func GetTeamProjects(s server.Server, w http.ResponseWriter, r *http.Request) error {
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
			"Invalid id",
		)
	}

	projectModule, err := s.Storage.GetProjectModule()
	if err != nil {
		return server.SendError(
			w, err, http.StatusInternalServerError,
			"storage misconfiguration",
		)
	}

	projects, err := projectModule.GetByTeamId(id)
	if err != nil {
		return server.SendError(
			w, err, http.StatusInternalServerError,
			"error getting team projects",
		)
	}

	return server.WriteJSON(
		w, http.StatusOK,
		projects,
	)
}

// AddUserToTeam godoc
// @Summary add user to team
// @Tags    team
// @Accept  json
// @Produce json
// @Param   teamId     path    int     true  "team id"
// @Param   userId     body    int     true  "user id"
// @Success 200  {object}  string
// @Failure 400  {object}  error
// @Failure 500  {object}  error
// @Router  /teams/{teamId}/members [post]
func AddUserToTeam(s server.Server, w http.ResponseWriter, r *http.Request) error {
	idStr, err := server.GetRequestParam(r, "teamId")
	if err != nil {
		return server.SendError(
			w, err, http.StatusBadRequest,
			"Parameter teamId not found",
		)
	}

	id, err := strconv.ParseUint(*idStr, 10, 64)
	if err != nil {
		return server.SendError(
			w, err, http.StatusBadRequest,
			"Invalid id",
		)
	}

	var request schemas.AddUserToTeamRequest
	if err := server.ReadJSON(r, &request.UserId); err != nil {
		return server.SendError(
			w, err, http.StatusBadRequest,
			"error parsing request",
		)
	}

	teamModule, err := s.Storage.GetTeamModule()
	if err != nil {
		return server.SendError(
			w, err, http.StatusInternalServerError,
			"storage misconfiguration",
		)
	}

	err = teamModule.AddUsers(id, request.UserId)
	if err != nil {
		return server.SendError(
			w, err, http.StatusInternalServerError,
			"error adding user to team",
		)
	}

	return server.WriteJSON(
		w, http.StatusOK,
		nil,
	)
}

// RemoveUserFromTeam godoc
// @Summary remove user from team
// @Tags    team
// @Produce json
// @Param   id         path    int     true  "team id"
// @Param   userId     path    int     true  "user id"
// @Success 200  {object}  string
// @Failure 400  {object}  error
// @Failure 500  {object}  error
// @Router  /teams/{id}/members/{userId} [delete]
func RemoveUserFromTeam(s server.Server, w http.ResponseWriter, r *http.Request) error {
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
			"Invalid id",
		)
	}

	userIdStr, err := server.GetRequestParam(r, "userId")
	if err != nil {
		return server.SendError(
			w, err, http.StatusBadRequest,
			"Parameter userId not found",
		)
	}

	userId, err := strconv.ParseUint(*userIdStr, 10, 64)
	if err != nil {
		return server.SendError(
			w, err, http.StatusBadRequest,
			"Invalid userId",
		)
	}

	teamModule, err := s.Storage.GetTeamModule()
	if err != nil {
		return server.SendError(
			w, err, http.StatusInternalServerError,
			"storage misconfiguration",
		)
	}

	err = teamModule.RemoveUsers(id, userId)
	if err != nil {
		return server.SendError(
			w, err, http.StatusInternalServerError,
			"error removing user from team",
		)
	}

	return server.WriteJSON(
		w, http.StatusOK,
		nil,
	)
}

// CreateTeam godoc
// @Summary create a team
// @Tags    team
// @Accept  json
// @Produce json
// @Param   team     body    schemas.CreateTeamRequest     true  "team"
// @Success 201  {object}  schemas.CreateTeamResponse
// @Failure 400  {object}  error
// @Failure 500  {object}  error
// @Router  /teams [post]
func CreateTeam(s server.Server, w http.ResponseWriter, r *http.Request) error {
	var request schemas.CreateTeamRequest
	if err := server.ReadJSON(r, &request.Team); err != nil {
		return server.SendError(
			w, err, http.StatusBadRequest,
			"error parsing request",
		)
	}

	team := models.NewCreateTeam(
		request.Team.Name,
	)

	teamModule, err := s.Storage.GetTeamModule()
	if err != nil {
		return server.SendError(
			w, err, http.StatusInternalServerError,
			"storage misconfiguration",
		)
	}

	id, err := teamModule.Create(team)
	if err != nil {
		return server.SendError(
			w, err, http.StatusInternalServerError,
			"error creating team",
		)
	}

	return server.WriteJSON(
		w, http.StatusCreated,
		schemas.CreateTeamResponse{Id: *id},
	)
}

// UpdateTeam godoc
// @Summary update a team
// @Tags    team
// @Produce json
// @Param   id     path     string  true  "team id"
// @Param   team    body     schemas.UpdateTeamRequest  true  "team"
// @Success 200  {string}  string
// @Failure 400  {object}  error
// @Failure 500  {object}  error
// @Router  /teams/{id} [put]
func UpdateTeam(s server.Server, w http.ResponseWriter, r *http.Request) error {
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

	var request schemas.UpdateTeamRequest
	if err := server.ReadJSON(r, &request.Team); err != nil {
		return server.SendError(
			w, err, http.StatusBadRequest,
			"error parsing request",
		)
	}

	team := models.NewTeam(
		id,
		request.Team.Name,
	)

	teamModule, err := s.Storage.GetTeamModule()
	if err != nil {
		return server.SendError(
			w, err, http.StatusInternalServerError,
			"storage misconfiguration",
		)
	}

	if err := teamModule.Update(team); err != nil {
		return server.SendError(
			w, err, http.StatusInternalServerError,
			"error updating team",
		)
	}

	return server.WriteJSON(
		w, http.StatusOK,
		"team updated",
	)
}

// DeleteTeam godoc
// @Summary delete a team
// @Tags    team
// @Produce json
// @Param   id     path     string  true  "team id"
// @Success 200  {string}  string
// @Failure 400  {object}  error
// @Failure 500  {object}  error
// @Router  /teams/{id} [delete]
func DeleteTeam(s server.Server, w http.ResponseWriter, r *http.Request) error {
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
			"Invalid id",
		)
	}

	teamModule, err := s.Storage.GetTeamModule()
	if err != nil {
		return server.SendError(
			w, err, http.StatusInternalServerError,
			"storage misconfiguration",
		)
	}

	err = teamModule.Delete(id)
	if err != nil {
		return server.SendError(
			w, err, http.StatusInternalServerError,
			"error deleting team",
		)
	}

	return server.WriteJSON(
		w, http.StatusOK,
		nil,
	)
}
