package models

type Contribution struct {
	Id        uint64 `json:"id" db:"id"`
	ProjectId uint64 `json:"projectId" db:"projectId"`
	UserId    uint64 `json:"userId" db:"userId"`
	Title     string `json:"title" db:"title"`
	Content   string `json:"content" db:"content"`
}

func NewContribution(
	id uint64,
	projectId uint64,
	userId uint64,
	title string,
	content string,
) Contribution {
	return Contribution{
		Id:        id,
		ProjectId: projectId,
		UserId:    userId,
		Title:     title,
		Content:   content,
	}
}

type ContributionDetail struct {
	Id      uint64 `json:"id" db:"id"`
	Project string `json:"project"`
	User    string `json:"user"`
	Title   string `json:"title" db:"title"`
	Content string `json:"content" db:"content"`
}

func NewContributionDetail(
	id uint64,
	project string,
	user string,
	title string,
	content string,
) ContributionDetail {
	return ContributionDetail{
		Id:      id,
		Project: project,
		User:    user,
		Title:   title,
		Content: content,
	}
}

type CreateContribution struct {
	ProjectId uint64 `json:"projectId" db:"projectId"`
	UserId    uint64 `json:"userId" db:"userId"`
	Title     string `json:"title" db:"title"`
	Content   string `json:"content" db:"content"`
}

func NewCreateContribution(
	projectId uint64,
	userId uint64,
	title string,
	content string,
) CreateContribution {
	return CreateContribution{
		ProjectId: projectId,
		UserId:    userId,
		Title:     title,
		Content:   content,
	}
}

type CreateContributionByNames struct {
	Project string `json:"project"`
	User    string `json:"user"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

func NewCreateContributionByNames(
	project string,
	user string,
	title string,
	content string,
) CreateContributionByNames {
	return CreateContributionByNames{
		Project: project,
		User:    user,
		Title:   title,
		Content: content,
	}
}

type ContributionFilter struct {
	Project *string `json:"project"`
	User    *string `json:"user"`
}

func NewContributionFilter(
	project *string,
	user *string,
) ContributionFilter {
	return ContributionFilter{
		Project: project,
		User:    user,
	}
}
