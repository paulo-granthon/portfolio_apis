package schemas

import "models"

type GetProjectRequest struct {
	Id uint64 `json:"id"`
}

type CreateProjectRequest struct {
	Project models.CreateProject `json:"project"`
}

type CreateProjectResponse struct {
	Id uint64 `json:"id"`
}

type UpdateProjectRequest struct {
	Project models.Project `json:"project"`
}
