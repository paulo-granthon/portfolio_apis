package schemas

type CreateProjectRequest struct {
	Name     string `json:"name"`
	Semester uint8  `json:"semester"`
	Company  string `json:"company"`
	Summary  string `json:"summary"`
	Url      string `json:"url"`
}

type CreateProjectResponse struct {
	Id uint64 `json:"id"`
}
