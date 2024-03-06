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
