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

func (s *PostgreProjectModule) Migrate() error {
	exampleProjects := []models.CreateProject{
		models.NewCreateProject(
			"Khali", 1, "FATEC", 1,
			"Avaliação 360",
			"github.com/taniacruzz/Khali",
		),
		models.NewCreateProject(
			"API2Semestre", 2, "2RP", 1,
			"Controle de Horas-Extras e Sobreavisos",
			"github.com/projetoKhali/API2Semestre",
		),
		models.NewCreateProject(
			"api3", 3, "2RP", 1,
			"Controle de Horas-Extras e Sobreavisos",
			"github.com/projetoKhali/api3",
		),
	}

	for _, p := range exampleProjects {
		if _, err := s.Create(p); err != nil {
			return err
		}
	}

	return nil
}

func (s *PostgreProjectModule) Get() ([]models.Project, error) {
	var projects []models.Project
	s.db.Find(&projects)
	return projects, nil
}

// function to get projects from any combination of optional parameters
// func (s *PostgreProjectModule) GetFilter(models.FilterProject) ([]models.Project, error) {
// }

func (s *PostgreProjectModule) GetById(id uint64) (*models.Project, error) {
	var project models.Project
	if err := s.db.First(&project, id).Error; err != nil {
		return nil, tracerr.Errorf("failed to get project by id: %w", tracerr.Wrap(err))
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
		Name:     p.Name,
		Semester: p.Semester,
		Company:  p.Company,
		TeamId:   p.TeamId,
		Summary:  p.Summary,
		Url:      p.Url,
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
