package models

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

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

func (y YearSemester) Value() (driver.Value, error) {
	return json.Marshal(y)
}

func (y *YearSemester) Scan(value interface{}) error {
	yearSemesterBytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("yearSemester must be a byte array")
	}
	return json.Unmarshal(yearSemesterBytes, y)
}

type User struct {
	Id                 uint64        `json:"id" db:"id"`
	Name               string        `json:"name" db:"name"`
	Summary            *string       `json:"summary" db:"summary"`
	SemesterMatriculed *YearSemester `json:"semesterMatriculed" db:"semesterMatriculed"`
	GithubUsername     *string       `json:"githubUsername" db:"githubUsername"`
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
	Name     string `json:"name" db:"name"`
	Password string `json:"password" db:"password"`
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

type CreateUser struct {
	Name               string        `json:"name" db:"name"`
	Password           string        `json:"password" db:"password"`
	Summary            *string       `json:"summary" db:"summary"`
	SemesterMatriculed *YearSemester `json:"semesterMatriculed" db:"semesterMatriculed"`
	GithubUsername     *string       `json:"githubUsername" db:"githubUsername"`
}

func NewCreateUser(
	name string,
	password string,
	summary *string,
	semesterMatriculed *YearSemester,
	githubUsername *string,
) CreateUser {
	return CreateUser{
		Name:               name,
		Password:           password,
		Summary:            summary,
		SemesterMatriculed: semesterMatriculed,
		GithubUsername:     githubUsername,
	}
}

type UpdateUser struct {
	Id                 uint64        `json:"id" db:"id"`
	Name               *string       `json:"name" db:"name"`
	Summary            *string       `json:"summary" db:"summary"`
	SemesterMatriculed *YearSemester `json:"semesterMatriculed" db:"semesterMatriculed"`
	GithubUsername     *string       `json:"githubUsername" db:"githubUsername"`
}

func NewUpdateUser(
	id uint64,
	name *string,
	summary *string,
	semesterMatriculed *YearSemester,
	githubUsername *string,
) UpdateUser {
	return UpdateUser{
		Id:                 id,
		Name:               name,
		Summary:            summary,
		SemesterMatriculed: semesterMatriculed,
		GithubUsername:     githubUsername,
	}
}

type FullUser struct {
	Id                 uint64        `json:"id" db:"id"`
	Name               string        `json:"name" db:"name"`
	Password           string        `json:"password" db:"password"`
	Summary            *string       `json:"summary" db:"summary"`
	SemesterMatriculed *YearSemester `json:"semesterMatriculed" db:"semesterMatriculed"`
	GithubUsername     *string       `json:"githubUsername" db:"githubUsername"`
}
