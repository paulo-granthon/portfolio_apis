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
