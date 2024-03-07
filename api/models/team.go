package models

type Team struct {
	Id   uint64 `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}

func NewTeam(
	id uint64,
	name string,
) Team {
	return Team{
		Id:   id,
		Name: name,
	}
}

type CreateTeam struct {
	Name string `json:"name" db:"name"`
}

func NewCreateTeam(
	name string,
) CreateTeam {
	return CreateTeam{
		Name: name,
	}
}

type TeamUser struct {
	TeamId uint64 `json:"teamId" db:"teamId"`
	UserId uint64 `json:"userId" db:"userId"`
}
