package models

type Project struct {
	Id       uint64 `json:"id"`
	Name     string `json:"name"`
	Semester uint8  `json:"semester"`
	Company  string `json:"company"`
	Summary  string `json:"summary"`
	Url      string `json:"url"`
}

func NewProject(
	id uint64,
	name string,
	semester uint8,
	company string,
	summary string,
	url string,
) Project {
	return Project{
		Id:       id,
		Name:     name,
		Semester: semester,
		Company:  company,
		Summary:  summary,
		Url:      url,
	}
}

type CreateProject struct {
	Name     string `json:"name"`
	Semester uint8  `json:"semester"`
	Company  string `json:"company"`
	Summary  string `json:"summary"`
	Url      string `json:"url"`
}

func NewCreateProject(
	name string,
	semester uint8,
	company string,
	Summary string,
	Url string,
) CreateProject {
	return CreateProject{
		Name:     name,
		Semester: semester,
		Company:  company,
		Summary:  Summary,
		Url:      Url,
	}
}
