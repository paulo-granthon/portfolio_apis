package service

import (
	"github.com/paulo-granthon/portfolio_apis/models"
	"github.com/paulo-granthon/portfolio_apis/storage"
	"sort"
	"strconv"

	"github.com/ztrue/tracerr"
)

// contributionGetter is the slice of ContributionService that PortfolioService
// depends on. Declaring it as an interface keeps Build unit-testable with a fake.
type contributionGetter interface {
	GetFilter(project *string, user *string) ([]models.ContributionDetail, error)
}

type PortfolioService struct {
	userStorage    storage.UserStorageModule
	projectStorage storage.ProjectStorageModule
	contributions  contributionGetter
}

func NewPortfolioService(
	s storage.Storage,
	contributions *ContributionService,
) (*PortfolioService, error) {
	userStorage, err := s.GetUserModule()
	if err != nil {
		return nil, tracerr.Errorf("failed to create portfolio service: failed to get UserModule from storage: %w", tracerr.Wrap(err))
	}

	projectStorage, err := s.GetProjectModule()
	if err != nil {
		return nil, tracerr.Errorf("failed to create portfolio service: failed to get ProjectModule from storage: %w", tracerr.Wrap(err))
	}

	return &PortfolioService{
		userStorage:    userStorage,
		projectStorage: projectStorage,
		contributions:  contributions,
	}, nil
}

// Build composes the portfolio document for a user: profile, their projects
// ordered by semester, and each project's contributions (with skills resolved).
// This composition is the single source of truth both the web view and the
// markdown export are rendered from.
func (s *PortfolioService) Build(userId uint64) (*models.Portfolio, error) {
	user, err := s.userStorage.GetById(userId)
	if err != nil {
		return nil, tracerr.Errorf("failed to build portfolio: failed to get user: %w", tracerr.Wrap(err))
	}

	projects, err := s.projectStorage.GetByUserId(userId)
	if err != nil {
		return nil, tracerr.Errorf("failed to build portfolio: failed to get projects: %w", tracerr.Wrap(err))
	}

	sort.SliceStable(projects, func(i, j int) bool {
		return projects[i].Semester < projects[j].Semester
	})

	userIdStr := strconv.FormatUint(userId, 10)
	portfolioProjects := make([]models.PortfolioProject, len(projects))
	for i, project := range projects {
		projectIdStr := strconv.FormatUint(project.Id, 10)
		details, err := s.contributions.GetFilter(&projectIdStr, &userIdStr)
		if err != nil {
			return nil, tracerr.Errorf("failed to build portfolio: failed to get contributions for project %d: %w", project.Id, tracerr.Wrap(err))
		}

		contributions := make([]models.PortfolioContribution, len(details))
		for j, detail := range details {
			contributions[j] = models.PortfolioContribution{
				Id:      detail.Id,
				Title:   detail.Title,
				Content: detail.Content,
				Skills:  detail.Skills,
			}
		}

		portfolioProjects[i] = models.PortfolioProject{
			Project:       project,
			Contributions: contributions,
		}
	}

	return &models.Portfolio{User: *user, Projects: portfolioProjects}, nil
}
