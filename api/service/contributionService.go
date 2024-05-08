package service

import (
	"models"
	"storage"

	"github.com/ztrue/tracerr"
)

type ContributionService struct {
	contributionStorage    *storage.ContributionStorageModule
	skillStorage   *storage.SkillStorageModule
	projectStorage *storage.ProjectStorageModule
	userStorage    *storage.UserStorageModule
}

func NewContributionService(
	storage storage.Storage,
) (*ContributionService, error) {
	contributionstorage, err := storage.GetContributionModule()
	if err != nil {
		return nil, tracerr.Errorf("failed to create contribution service: failed to get ContributionModule from storage: %w", tracerr.Wrap(err))
	}

	skillStorage, err := storage.GetSkillModule()
	if err != nil {
		return nil, tracerr.Errorf("failed to create contribution service: failed to get SkillModule from storage: %w", tracerr.Wrap(err))
	}

	projectStorage, err := storage.GetProjectModule()
	if err != nil {
		return nil, tracerr.Errorf("failed to create contribution service: failed to get ProjectModule from storage: %w", tracerr.Wrap(err))
	}

	userStorage, err := storage.GetUserModule()
	if err != nil {
		return nil, tracerr.Errorf("failed to create contribution service: failed to get UserModule from storage: %w", tracerr.Wrap(err))
	}

	return &ContributionService{
		contributionStorage:    &contributionstorage,
		skillStorage:   &skillStorage,
		projectStorage: &projectStorage,
		userStorage:    &userStorage,
	}, nil
}

func (s *ContributionService) Create(n models.CreateContributionByNames) (*uint64, error) {
	skillStorage := *s.skillStorage
	skill, err := skillStorage.GetByName(n.Skill)
	if err != nil {
		return nil, tracerr.Errorf("failed to create contribution: failed to get skill by name: %w", tracerr.Wrap(err))
	}

	projectStorage := *s.projectStorage
	project, err := projectStorage.GetByName(n.Project)
	if err != nil {
		return nil, tracerr.Errorf("failed to create contribution: failed to get project by name: %w", tracerr.Wrap(err))
	}

	userStorage := *s.userStorage
	user, err := userStorage.GetByUsername(n.User)
	if err != nil {
		return nil, tracerr.Errorf("failed to create contribution: failed to get user by github username: %w", tracerr.Wrap(err))
	}

	contribution := models.CreateContribution{
		SkillId:   skill.Id,
		ProjectId: project.Id,
		UserId:    user.Id,
		Content:   n.Content,
	}

	contributionStorage := *s.contributionStorage

	id, err := contributionStorage.Create(contribution)
	if err != nil {
		return nil, tracerr.Errorf("failed to create contribution: %w", tracerr.Wrap(err))
	}

	return id, nil
}
