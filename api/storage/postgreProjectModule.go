package storage

import (
	"database/sql"
	"models"
)

type PostgreProjectModule struct {
	db *sql.DB
}

func (s *PostgreProjectModule) Migrate() error {
	_, err := s.db.Exec(`
		CREATE TABLE IF NOT EXISTS projects (
			id SERIAL PRIMARY KEY,
			name VARCHAR(50) NOT NULL,
			semester UINT1 NOT NULL,
			company VARCHAR(100) NOT NULL
		)
	`)
	if err != nil {
		return err
	}

	exampleProjects := []models.CreateProject{
		models.NewCreateProject(
			"Khali", 1, "FATEC",
			"Avaliação 360",
			"github.com/taniacruzz/Khali",
		),
		models.NewCreateProject(
			"API2Semestre", 2, "2RP",
			"Controle de Horas-Extras e Sobreavisos",
			"github.com/projetoKhali/API2Semestre",
		),
		models.NewCreateProject(
			"api3", 3, "2RP",
			"Controle de Horas-Extras e Sobreavisos",
			"github.com/projetoKhali/api3",
		),
	}

	for _, p := range exampleProjects {
		if _, err := s.CreateProject(p); err != nil {
			return err
		}
	}

	return nil
}

func (s *PostgreProjectModule) GetProjects() ([]*models.Project, error) {
	rows, err := s.db.Query(`
		SELECT id, name, semester, company
		FROM projects
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var projects []*models.Project
	for rows.Next() {
		p := &models.Project{}
		if err := rows.Scan(&p.Id, &p.Name, &p.Semester, &p.Company); err != nil {
			return nil, err
		}
		projects = append(projects, p)
	}
	return projects, nil
}

func (s *PostgreProjectModule) GetProject(id uint64) (*models.Project, error) {
	p := &models.Project{}
	return p, s.db.QueryRow(`
		SELECT id, name, semester, company
		FROM projects
		WHERE id = $1
	`, id).Scan(&p.Id, &p.Name, &p.Semester, &p.Company)
}

func (s *PostgreProjectModule) CreateProject(p models.CreateProject) (*uint64, error) {
	var id uint64
	if err := s.db.QueryRow(`
		INSERT INTO projects (name, semester, company)
		VALUES ($1, $2, $3)
		RETURNING id
	`, p.Name, p.Semester, p.Company,
	).Scan(&id); err != nil {
		return nil, err
	}

	return &id, nil
}

func (s *PostgreProjectModule) UpdateProject(p models.Project) error {
	_, err := s.db.Exec(`
		UPDATE projects
		SET name = $1, semester = $2, company = $3
		WHERE id = $4
	`, p.Name, p.Semester, p.Company, p.Id)
	return err
}

func (s *PostgreProjectModule) DeleteProject(id uint64) error {
	_, err := s.db.Exec(`
		DELETE FROM projects
		WHERE id = $1
	`, id)
	return err
}

func (s *PostgreProjectModule) Close() {
	s.db.Close()
}
