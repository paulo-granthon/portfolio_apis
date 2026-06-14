package models

// Participation is the qualitative summary of a single user's overall work on a
// single project: one row per (user, project) pair. It complements the granular
// Contribution entries with a narrative overview.
type Participation struct {
	Id        uint64 `json:"id" db:"id"`
	ProjectId uint64 `json:"projectId" db:"projectId"`
	UserId    uint64 `json:"userId" db:"userId"`
	Summary   string `json:"summary" db:"summary"`
}

type CreateParticipation struct {
	ProjectId uint64 `json:"projectId" db:"projectId"`
	UserId    uint64 `json:"userId" db:"userId"`
	Summary   string `json:"summary" db:"summary"`
}

func NewCreateParticipation(
	projectId uint64,
	userId uint64,
	summary string,
) CreateParticipation {
	return CreateParticipation{
		ProjectId: projectId,
		UserId:    userId,
		Summary:   summary,
	}
}

// CreateParticipationByNames mirrors CreateContributionByNames: it lets seeds
// reference a project by name and a user by github username instead of ids.
type CreateParticipationByNames struct {
	Project string `json:"project"`
	User    string `json:"user"`
	Summary string `json:"summary"`
}

func NewCreateParticipationByNames(
	project string,
	user string,
	summary string,
) CreateParticipationByNames {
	return CreateParticipationByNames{
		Project: project,
		User:    user,
		Summary: summary,
	}
}
