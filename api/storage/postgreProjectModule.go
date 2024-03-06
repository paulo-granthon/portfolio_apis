package storage

import (
	"models"

	"github.com/jmoiron/sqlx"
)

type PostgreProjectModule struct {
	db *sqlx.DB
}

func NewPostgreProjectModule(db *sqlx.DB) (*PostgreProjectModule, error) {
	return &PostgreProjectModule{db: db}, nil
}

func (s *PostgreProjectModule) Migrate() error {
	_, err := s.db.Exec(`
		CREATE TABLE IF NOT EXISTS projects (
			id SERIAL PRIMARY KEY,
			name VARCHAR(50) NOT NULL,
			semester UINT1 NOT NULL,
			company VARCHAR(100) NOT NULL,
			teamId INT NOT NULL,
			summary TEXT NOT NULL,
			url VARCHAR(100) NOT NULL,
			FOREIGN KEY (teamId) REFERENCES teams(id)
		)
	`)
	if err != nil {
		return err
	}

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

func (s *PostgreProjectModule) Get() ([]*models.Project, error) {
	rows, err := s.db.Query(`
		SELECT id, name, semester, company, teamId, summary, url
		FROM projects
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var projects []*models.Project
	for rows.Next() {
		p := &models.Project{}
		if err := rows.Scan(&p.Id, &p.Name, &p.Semester, &p.Company, &p.TeamId, &p.Summary, &p.Url); err != nil {
			return nil, err
		}
		projects = append(projects, p)
	}
	return projects, nil
}

func (s *PostgreProjectModule) GetById(id uint64) (*models.Project, error) {
	p := &models.Project{}
	return p, s.db.QueryRow(`
		SELECT id, name, semester, company, teamId, summary, url
		FROM projects
		WHERE id = $1
	`, id).Scan(&p.Id, &p.Name, &p.Semester, &p.Company, &p.TeamId, &p.Summary, &p.Url)
}

func (s *PostgreProjectModule) GetByUserId(id uint64) ([]*models.Project, error) {
	rows, err := s.db.Query(`
		SELECT p.id, p.name, p.semester, p.company, p.teamId, p.summary, p.url
		FROM projects p
		JOIN teams t ON p.teamId = t.id
		JOIN teamUsers tu ON t.id = tu.teamId
		WHERE tu.userId = $1
	`, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var projects []*models.Project
	for rows.Next() {
		p := &models.Project{}
		if err := rows.Scan(&p.Id, &p.Name, &p.Semester, &p.Company, &p.TeamId, &p.Summary, &p.Url); err != nil {
			return nil, err
		}
		projects = append(projects, p)
	}
	return projects, nil
}

func (s *PostgreProjectModule) Create(p models.CreateProject) (*uint64, error) {
	var id uint64
	if err := s.db.QueryRow(`
		INSERT INTO projects (name, semester, company, teamId, summary, url)
		VALUES ($1, $2, $3)
		RETURNING id
	`, p.Name, p.Semester, p.Company, p.TeamId, p.Summary, p.Url,
	).Scan(&id); err != nil {
		return nil, err
	}

	return &id, nil
}

func (s *PostgreProjectModule) Update(p models.UpdateProject) error {
	_, err := s.db.Exec(`
		UPDATE projects
		SET name = $1, semester = $2, company = $3, teamId = $4, summary = $5, url = $6
		WHERE id = $4
	`, p.Name, p.Semester, p.Company, p.TeamId, p.Summary, p.Url, p.Id)
	return err
}

func (s *PostgreProjectModule) Delete(id uint64) error {
	_, err := s.db.Exec(`
		DELETE FROM projects
		WHERE id = $1
	`, id)
	return err
}

func (s *PostgreProjectModule) Close() {
	s.db.Close()
}
