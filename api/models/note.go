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

type NoteDetail struct {
	Id      uint64 `json:"id" db:"id"`
	Skill   string `json:"skill"`
	Project string `json:"project"`
	User    string `json:"user"`
	Content string `json:"content" db:"content"`
}

func NewNoteDetail(
	id uint64,
	skill string,
	project string,
	user string,
	content string,
) NoteDetail {
	return NoteDetail{
		Id:      id,
		Skill:   skill,
		Project: project,
		User:    user,
		Content: content,
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

type NoteFilter struct {
	Skill   *string `json:"skill"`
	Project *string `json:"project"`
	User    *string `json:"user"`
}

func NewNoteFilter(
	skill *string,
	project *string,
	user *string,
) NoteFilter {
	return NoteFilter{
		Skill:   skill,
		Project: project,
		User:    user,
	}
}
