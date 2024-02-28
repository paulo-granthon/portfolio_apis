package models

type Project struct {
	Id       uint64 `json:"id"`
	Name     string `json:"name"`
	Semester uint8  `json:"semester"`
	Company  string `json:"company"`
}

func NewProject(id uint64, name string, semester uint8, company string) Project {
	return Project{
		Id:       id,
		Name:     name,
		Semester: semester,
		Company:  company,
	}
}

type CreateProject struct {
	Name     string `json:"name"`
	Semester uint8  `json:"semester"`
	Company  string `json:"company"`
}

func NewCreateProject(name string, semester uint8, company string) CreateProject {
	return CreateProject{
		Name:     name,
		Semester: semester,
		Company:  company,
	}
}
