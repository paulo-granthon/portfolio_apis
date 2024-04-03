package models

type Project struct {
	Id          uint64 `json:"id" db:"id"`
	Name        string `json:"name" db:"name"`
	Semester    uint8  `json:"semester" db:"semester"`
	Company     string `json:"company" db:"company"`
	TeamId      uint64 `json:"teamId" db:"teamId"`
	Summary     string `json:"summary" db:"summary"`
	Description string `json:"description" db:"description"`
	Url         string `json:"url" db:"url"`
}

func NewProject(
	id uint64,
	name string,
	semester uint8,
	company string,
	teamId uint64,
	summary string,
	description string,
	url string,
) Project {
	return Project{
		Id:          id,
		Name:        name,
		Semester:    semester,
		Company:     company,
		TeamId:      teamId,
		Summary:     summary,
		Description: description,
		Url:         url,
	}
}

type CreateProject struct {
	Name        string `json:"name" db:"name"`
	Semester    uint8  `json:"semester" db:"semester"`
	Company     string `json:"company" db:"company"`
	TeamId      uint64 `json:"teamId" db:"teamId"`
	Summary     string `json:"summary" db:"summary"`
	Description string `json:"description" db:"description"`
	Url         string `json:"url" db:"url"`
}

func NewCreateProject(
	name string,
	semester uint8,
	company string,
	teamId uint64,
	summary string,
	description string,
	url string,
) CreateProject {
	return CreateProject{
		Name:        name,
		Semester:    semester,
		Company:     company,
		TeamId:      teamId,
		Summary:     summary,
		Description: description,
		Url:         url,
	}
}

type UpdateProject struct {
	Id          uint64  `json:"id" db:"id"`
	Name        *string `json:"name" db:"name"`
	Semester    *uint8  `json:"semester" db:"semester"`
	Company     *string `json:"company" db:"company"`
	TeamId      *uint64 `json:"teamId" db:"teamId"`
	Summary     *string `json:"summary" db:"summary"`
	Description *string `json:"description" db:"description"`
	Url         *string `json:"url" db:"url"`
}

func NewUpdateProject(
	id uint64,
	name *string,
	semester *uint8,
	company *string,
	teamId *uint64,
	summary *string,
	description *string,
	url *string,
) UpdateProject {
	return UpdateProject{
		Id:          id,
		Name:        name,
		Semester:    semester,
		Company:     company,
		TeamId:      teamId,
		Summary:     summary,
		Description: description,
		Url:         url,
	}
}

type FilterProject struct {
	Id          *uint64 `json:"id" db:"id"`
	Name        *string `json:"name" db:"name"`
	Semester    *uint8  `json:"semester" db:"semester"`
	Company     *string `json:"company" db:"company"`
	TeamId      *uint64 `json:"teamId" db:"teamId"`
	Summary     *string `json:"summary" db:"summary"`
	Description *string `json:"description" db:"description"`
	Url         *string `json:"url" db:"url"`
}

func NewFilterProject(
	id *uint64,
	name *string,
	semester *uint8,
	company *string,
	teamId *uint64,
	summary *string,
	description *string,
	url *string,
) FilterProject {
	return FilterProject{
		Id:          id,
		Name:        name,
		Semester:    semester,
		Company:     company,
		TeamId:      teamId,
		Summary:     summary,
		Description: description,
		Url:         url,
	}
}
