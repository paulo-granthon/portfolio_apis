package models

type Contribution struct {
	Id        uint64 `json:"id" db:"id"`
	SkillId   uint64 `json:"skillId" db:"skillId"`
	ProjectId uint64 `json:"projectId" db:"projectId"`
	UserId    uint64 `json:"userId" db:"userId"`
	Title     string `json:"title" db:"title"`
	Content   string `json:"content" db:"content"`
}

func NewContribution(
	id uint64,
	skillId uint64,
	projectId uint64,
	userId uint64,
	title string,
	content string,
) Contribution {
	return Contribution{
		Id:        id,
		SkillId:   skillId,
		ProjectId: projectId,
		UserId:    userId,
		Title:     title,
		Content:   content,
	}
}

type ContributionDetail struct {
	Id      uint64 `json:"id" db:"id"`
	Skill   string `json:"skill"`
	Project string `json:"project"`
	User    string `json:"user"`
	Title   string `json:"title" db:"title"`
	Content string `json:"content" db:"content"`
}

func NewContributionDetail(
	id uint64,
	skill string,
	project string,
	user string,
	title string,
	content string,
) ContributionDetail {
	return ContributionDetail{
		Id:      id,
		Skill:   skill,
		Project: project,
		User:    user,
		Title:   title,
		Content: content,
	}
}

type CreateContribution struct {
	SkillId   uint64 `json:"skillId" db:"skillId"`
	ProjectId uint64 `json:"projectId" db:"projectId"`
	UserId    uint64 `json:"userId" db:"userId"`
	Title     string `json:"title" db:"title"`
	Content   string `json:"content" db:"content"`
}

func NewCreateContribution(
	skillId uint64,
	projectId uint64,
	userId uint64,
	title string,
	content string,
) CreateContribution {
	return CreateContribution{
		SkillId:   skillId,
		ProjectId: projectId,
		UserId:    userId,
		Title:     title,
		Content:   content,
	}
}

type CreateContributionByNames struct {
	Skill   string `json:"skill"`
	Project string `json:"project"`
	User    string `json:"user"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

func NewCreateContributionByNames(
	skill string,
	project string,
	user string,
	title string,
	content string,
) CreateContributionByNames {
	return CreateContributionByNames{
		Skill:   skill,
		Project: project,
		User:    user,
		Title:   title,
		Content: content,
	}
}

type ContributionFilter struct {
	Skill   *string `json:"skill"`
	Project *string `json:"project"`
	User    *string `json:"user"`
}

func NewContributionFilter(
	skill *string,
	project *string,
	user *string,
) ContributionFilter {
	return ContributionFilter{
		Skill:   skill,
		Project: project,
		User:    user,
	}
}
