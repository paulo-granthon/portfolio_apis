package main

type Project struct {
	Id       uint64 `json:"id"`
	Name     string `json:"name"`
	Semester uint8  `json:"semester"`
	Company  string `json:"company"`
}

func exampleProjects() []Project {
	return []Project{
		NewProject(1, "Khali", 1, "FATEC"),
		NewProject(2, "API2Semestre", 2, "2RP"),
		NewProject(3, "api3", 3, "2RP"),
	}
}

func NewProject(id uint64, name string, semester uint8, company string) Project {
	return Project{id, name, semester, company}
}
}
