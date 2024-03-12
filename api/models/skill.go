package models

type Skill struct {
	Id   uint64 `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}

func NewSkill(
	id uint64,
	name string,
) Skill {
	return Skill{
		Id:   id,
		Name: name,
	}
}

type CreateSkill struct {
	Name string `json:"name" db:"name"`
}

func NewCreateSkill(
	name string,
) CreateSkill {
	return CreateSkill{
		Name: name,
	}
}
