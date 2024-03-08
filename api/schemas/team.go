package schemas

import "models"

type GetTeamRequest struct {
	Id uint64 `json:"id"`
}

type GetTeamsResponse struct {
	Teams []models.Team `json:"teams"`
}

type GetTeamResponse struct {
	Team models.Team `json:"team"`
}

type AddUserToTeamRequest struct {
	TeamId uint64 `json:"teamId"`
	UserId uint64 `json:"userId"`
}

type AddUserToTeamResponse struct {
	Success bool `json:"success"`
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
