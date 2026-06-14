package service

import (
	"github.com/paulo-granthon/portfolio_apis/models"
	"github.com/paulo-granthon/portfolio_apis/storage"
	"testing"
)

// Fakes embed the storage interfaces so only the methods Build uses need bodies;
// any other call would panic (and would signal Build doing more than expected).
type fakeUserStorage struct {
	storage.UserStorageModule
	user *models.User
}

func (f fakeUserStorage) GetById(uint64) (*models.User, error) { return f.user, nil }

type fakeProjectStorage struct {
	storage.ProjectStorageModule
	projects []models.Project
}

func (f fakeProjectStorage) GetByUserId(uint64) ([]models.Project, error) { return f.projects, nil }

type fakeContributions struct {
	byProjectId map[string][]models.ContributionDetail
}

func (f fakeContributions) GetFilter(project *string, _ *string) ([]models.ContributionDetail, error) {
	return f.byProjectId[*project], nil
}

type fakeParticipations struct {
	byProjectId map[uint64]string
}

func (f fakeParticipations) GetByUserAndProject(_ uint64, projectId uint64) (*models.Participation, error) {
	if summary, ok := f.byProjectId[projectId]; ok {
		return &models.Participation{ProjectId: projectId, Summary: summary}, nil
	}
	return nil, nil
}

func (f fakeParticipations) Create(models.CreateParticipation) (*uint64, error) { return nil, nil }

func TestBuild_OrdersBySemesterAndGroupsContributions(t *testing.T) {
	svc := &PortfolioService{
		userStorage: fakeUserStorage{user: &models.User{Id: 7, Name: "Paulo"}},
		projectStorage: fakeProjectStorage{projects: []models.Project{
			{Id: 30, Name: "api3", Semester: 3},
			{Id: 10, Name: "Khali", Semester: 1},
			{Id: 20, Name: "api2", Semester: 2},
		}},
		participations: fakeParticipations{byProjectId: map[uint64]string{
			10: "Resumo da participação no Khali.",
		}},
		contributions: fakeContributions{byProjectId: map[string][]models.ContributionDetail{
			"10": {{Id: 1, Title: "c1", Skills: []string{"Go"}}},
			"20": {{Id: 2, Title: "c2"}, {Id: 3, Title: "c3"}},
			// project 30 has no contributions
		}},
	}

	portfolio, err := svc.Build(7)
	if err != nil {
		t.Fatalf("Build returned error: %v", err)
	}

	if portfolio.User.Id != 7 {
		t.Errorf("expected user id 7, got %d", portfolio.User.Id)
	}

	gotOrder := make([]uint8, len(portfolio.Projects))
	for i, p := range portfolio.Projects {
		gotOrder[i] = p.Semester
	}
	if gotOrder[0] != 1 || gotOrder[1] != 2 || gotOrder[2] != 3 {
		t.Errorf("projects not ordered by semester ascending: got %v", gotOrder)
	}

	// Khali (semester 1, id 10) -> 1 contribution carrying skills, plus its participation summary.
	khali := portfolio.Projects[0]
	if khali.Participation != "Resumo da participação no Khali." {
		t.Errorf("expected Khali participation summary, got %q", khali.Participation)
	}
	if portfolio.Projects[2].Participation != "" {
		t.Errorf("expected empty participation for api3, got %q", portfolio.Projects[2].Participation)
	}
	if len(khali.Contributions) != 1 || khali.Contributions[0].Title != "c1" {
		t.Errorf("unexpected contributions for Khali: %+v", khali.Contributions)
	}
	if len(khali.Contributions[0].Skills) != 1 || khali.Contributions[0].Skills[0] != "Go" {
		t.Errorf("skills not propagated: %+v", khali.Contributions[0].Skills)
	}

	// api2 (semester 2, id 20) -> 2 contributions.
	if len(portfolio.Projects[1].Contributions) != 2 {
		t.Errorf("expected 2 contributions for api2, got %d", len(portfolio.Projects[1].Contributions))
	}

	// api3 (semester 3, id 30) -> 0 contributions, non-nil slice.
	if portfolio.Projects[2].Contributions == nil || len(portfolio.Projects[2].Contributions) != 0 {
		t.Errorf("expected empty contributions for api3, got %+v", portfolio.Projects[2].Contributions)
	}
}
