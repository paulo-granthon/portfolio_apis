package models

type Note struct {
	Id        uint64 `json:"id" db:"id"`
	SkillId   uint64 `json:"skillId" db:"skillId"`
	ProjectId uint64 `json:"projectId" db:"projectId"`
	UserID    uint64 `json:"userId" db:"userId"`
	Content   string `json:"content" db:"content"`
}

func NewNote(
	id uint64,
	skillId uint64,
	projectId uint64,
	userID uint64,
	content string,
) Note {
	return Note{
		Id:        id,
		SkillId:   skillId,
		ProjectId: projectId,
		UserID:    userID,
		Content:   content,
	}
}

type CreateNote struct {
	SkillId   uint64 `json:"skillId" db:"skillId"`
	ProjectId uint64 `json:"projectId" db:"projectId"`
	UserID    uint64 `json:"userId" db:"userId"`
	Content   string `json:"content" db:"content"`
}

func NewCreateNote(
	skillId uint64,
	projectId uint64,
	userID uint64,
	content string,
) CreateNote {
	return CreateNote{
		SkillId:   skillId,
		ProjectId: projectId,
		UserID:    userID,
		Content:   content,
	}
}