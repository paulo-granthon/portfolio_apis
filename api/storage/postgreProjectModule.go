package storage

import (
	"models"

	"github.com/ztrue/tracerr"
	"gorm.io/gorm"
)

type PostgreProjectModule struct {
	db *gorm.DB
}

func NewPostgreProjectModule(db *gorm.DB) (*PostgreProjectModule, error) {
	return &PostgreProjectModule{db: db}, nil
}

func (s *PostgreProjectModule) Get() ([]models.Project, error) {
	var projects []models.Project
	s.db.Find(&projects)
	return projects, nil
}

func (s *PostgreProjectModule) GetById(id uint64) (*models.Project, error) {
	var project models.Project
	if err := s.db.First(&project, id).Error; err != nil {
		return nil, tracerr.Errorf("failed to get project by id: %w", tracerr.Wrap(err))
	}
	return &project, nil
}

func (s *PostgreProjectModule) GetByName(name string) (*models.Project, error) {
	var project models.Project
	if err := s.db.Where("name = ?", name).First(&project).Error; err != nil {
		return nil, tracerr.Errorf("failed to get project by name: %w", tracerr.Wrap(err))
	}
	return &project, nil
}

func (s *PostgreProjectModule) GetByUserId(id uint64) ([]models.Project, error) {
	var projects []models.Project
	subQuery := s.db.
		Select("team_id").
		Where("user_id = ?", id).
		Table("team_users ")
	s.db.
		Table("projects").
		Where("team_id = (?)", subQuery).
		Find(&projects)
	return projects, nil
}

func (s *PostgreProjectModule) GetByTeamId(id uint64) ([]models.Project, error) {
	var projects []models.Project
	s.db.Where("team_id = ?", id).Find(&projects)
	return projects, nil
}

func (s *PostgreProjectModule) Create(p models.CreateProject) (*uint64, error) {
	project := models.Project{
		Name:        p.Name,
		Image:       p.Image,
		Semester:    p.Semester,
		Company:     p.Company,
		TeamId:      p.TeamId,
		Summary:     p.Summary,
		Description: p.Description,
		Url:         p.Url,
	}

	if err := s.db.Create(&project).Error; err != nil {
		return nil, tracerr.Errorf("failed to create project: %w", tracerr.Wrap(err))
	}
	return &project.Id, nil
}

func (s *PostgreProjectModule) Update(p models.UpdateProject) error {
	if err := s.db.Model(&models.Project{}).Where("id = ?", p.Id).Updates(p).Error; err != nil {
		return tracerr.Errorf("failed to update project: %w", tracerr.Wrap(err))
	}
	return nil
}

func (s *PostgreProjectModule) Delete(id uint64) error {
	if err := s.db.Delete(&models.Project{}, id).Error; err != nil {
		return tracerr.Errorf("failed to delete project: %w", tracerr.Wrap(err))
	}
	return nil
}
