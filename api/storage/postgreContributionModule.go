package storage

import (
	"models"

	"github.com/ztrue/tracerr"
	"gorm.io/gorm"
)

type PostgreContributionModule struct {
	db *gorm.DB
}

func NewPostgreContributionModule(db *gorm.DB) (*PostgreContributionModule, error) {
	return &PostgreContributionModule{db: db}, nil
}

func (s *PostgreContributionModule) Get() ([]models.Contribution, error) {
	var contributions []models.Contribution
	s.db.Find(&contributions)
	return contributions, nil
}

func (s *PostgreContributionModule) GetById(id uint64) (*models.Contribution, error) {
	var contribution models.Contribution
	if err := s.db.First(&contribution, id).Error; err != nil {
		return nil, tracerr.Errorf("failed to get contribution by id: %w", tracerr.Wrap(err))
	}
	return &contribution, nil
}

func (s *PostgreContributionModule) GetFilter(f models.ContributionFilter) ([]models.ContributionDetail, error) {
	var contributions []models.ContributionDetail

	query := s.db.
		Table("contributions").
		Select(`
			contributions.id,
			contributions.title,
			contributions.content,
			users.name AS user,
			projects.name AS project
		`).
		Joins("JOIN projects ON contributions.project_id = projects.id").
		Joins("JOIN users ON contributions.user_id = users.id")

	if f.Project != nil && *f.Project != "" {
		query = query.Where("contributions.project_id = ?", *f.Project)
	}

	if f.User != nil && *f.User != "" {
		query = query.Where("contributions.user_id = ?", *f.User)
	}

	result, err := query.Rows()
	if err != nil {
		return nil, tracerr.Errorf("failed to get contributions by filter: %w", tracerr.Wrap(err))
	}

	for result.Next() {
		var contribution models.ContributionDetail
		if err := result.Scan(
			&contribution.Id,
			&contribution.Title,
			&contribution.Content,
			&contribution.User,
			&contribution.Project,
		); err != nil {
			return nil, tracerr.Errorf("failed to scan contribution: %w", tracerr.Wrap(err))
		}
		contributions = append(contributions, contribution)
	}

	return contributions, nil
}

func (s *PostgreContributionModule) GetByProjectId(id uint64) ([]models.Contribution, error) {
	var contributions []models.Contribution
	s.db.Where("project_id = ?", id).Find(&contributions)
	return contributions, nil
}

func (s *PostgreContributionModule) GetByUserId(id uint64) ([]models.Contribution, error) {
	var contributions []models.Contribution
	s.db.Where("user_id = ?", id).Find(&contributions)
	return contributions, nil
}

func (s *PostgreContributionModule) AddSkills(id uint64, skills ...uint64) error {
	var contributionSkills []models.ContributionSkill
	for _, skillId := range skills {
		contributionSkills = append(contributionSkills, models.ContributionSkill{
			ContributionId: id,
			SkillId:        skillId,
		})
	}

	if err := s.db.Create(&contributionSkills).Error; err != nil {
		return tracerr.Errorf("failed to add skills to contribution: %w", tracerr.Wrap(err))
	}

	return nil
}

func (s *PostgreContributionModule) RemoveSkills(id uint64, skills ...uint64) error {
	var contributionSkills []models.ContributionSkill
	for _, skillId := range skills {
		contributionSkills = append(contributionSkills, models.ContributionSkill{
			ContributionId: id,
			SkillId:        skillId,
		})
	}

	if err := s.db.Delete(&contributionSkills).Error; err != nil {
		return tracerr.Errorf("failed to remove skills from contribution: %w", tracerr.Wrap(err))
	}

	return nil
}

func (s *PostgreContributionModule) Create(n models.CreateContribution) (*uint64, error) {
	contribution := models.Contribution{
		ProjectId: n.ProjectId,
		UserId:    n.UserId,
		Title:     n.Title,
		Content:   n.Content,
	}
	if err := s.db.Create(&contribution).Error; err != nil {
		return nil, tracerr.Errorf("failed to create contribution: %w", tracerr.Wrap(err))
	}
	return &contribution.Id, nil
}

func (s *PostgreContributionModule) Delete(id uint64) error {
	if err := s.db.Delete(&models.Contribution{}, id).Error; err != nil {
		return tracerr.Errorf("failed to delete contribution: %w", tracerr.Wrap(err))
	}
	return nil
}
