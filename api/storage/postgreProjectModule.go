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
			"A plataforma Khali permite a implementação do método de Avaliação 360° na Instituição de Ensino fictícia PBLTeX. Este projeto de API do 1º Semestre de Banco de Dados da Fatec - São José dos Campos possibilita uma abordagem abrangente na avaliação dos diversos aspectos da instituição, promovendo uma análise holística e aprimorando processos de gestão e desenvolvimento.",
			"github.com/taniacruzz/Khali",
		),
		models.NewCreateProject(
			"API2Semestre", 2, "2RP", 1,
			"Controle de Horas-Extras e Sobreavisos (desktop)",
			"A API desenvolvida no 2° semestre do curso de Banco de Dados na Fatec - SJC proporciona um sistema desktop especializado no registro de horas extras e sobreavisos pelos colaboradores, com funcionalidades de controle tanto para gestores (PO) quanto para administradores (RH e Financeiro). Essa solução oferece uma plataforma integrada e eficiente para gerenciamento de tempo e recursos humanos, contribuindo para uma gestão mais eficaz e transparente dentro da organização.",
			"github.com/projetoKhali/API2Semestre",
		),
		models.NewCreateProject(
			"api3", 3, "2RP", 1,
			"Controle de Horas-Extras e Sobreavisos (web)",
			"Sistema desenvolvido para auxiliar na gestão eficiente das horas trabalhadas pelos colaboradores de uma empresa. Ele automatiza a identificação e classificação de horas extras e sobreavisos, simplificando os processos de controle para os departamentos pessoal e financeiro.",
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
