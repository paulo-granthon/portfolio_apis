package models

type YearSemester struct {
	Year     uint16 `json:"year"`
	Semester uint8  `json:"semester"`
}

func NewYearSemester(
	year uint16,
	semester uint8,
) YearSemester {
	return YearSemester{
		Year:     year,
		Semester: ((semester - 1) % 2) + 1, // ensure to 1 or 2
	}
}

type User struct {
	Id                 uint64        `json:"id"`
	Name               string        `json:"name"`
	Summary            *string       `json:"summary"`
	SemesterMatriculed *YearSemester `json:"semesterMatriculed"`
	GithubUsername     *string       `json:"githubUsername"`
}

func NewUser(
	id uint64,
	name string,
	summary *string,
	semesterMatriculed *YearSemester,
	githubUsername *string,
) User {
	return User{
		Id:                 id,
		Name:               name,
		Summary:            summary,
		SemesterMatriculed: semesterMatriculed,
		GithubUsername:     githubUsername,
	}
}

type RegisterUser struct {
	Name     string `json:"name"`
	Password string
}

func NewRegisterUser(
	name string,
	password string,
) RegisterUser {
	return RegisterUser{
		Name:     name,
		Password: password,
	}
}
