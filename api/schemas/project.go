package schemas

import "models"

type CreateProjectRequest struct {
	Project models.CreateProject `json:"project"`
}

type CreateProjectResponse struct {
	Id uint64 `json:"id"`
}
