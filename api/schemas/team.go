package schemas

import "models"

type GetTeamRequest struct {
	Id uint64 `json:"id"`
}

type GetTeamResponse struct {
	Team models.Team `json:"team"`
}

type CreateTeamRequest struct {
	Team models.CreateTeam `json:"team"`
}

type CreateTeamResponse struct {
	Id uint64 `json:"id"`
}

type UpdateTeamRequest struct {
	Team models.Team `json:"team"`
}

type UpdateTeamResponse struct {
	Success bool `json:"success"`
}
