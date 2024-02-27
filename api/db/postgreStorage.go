package db

import (
	"database/sql"
	"models"
)

type PostgreStorage struct {
	db *sql.DB
}

func NewPostgreStorage() (*PostgreStorage, error) {
	connStr := "user=postgres password=secret sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &PostgreStorage{
		db: db,
	}, nil
}

func (s *PostgreStorage) Migrate() error {
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

	exampleProjects := []models.Project{
		models.NewProject(1, "Khali", 1, "FATEC"),
		models.NewProject(2, "API2Semestre", 2, "2RP"),
		models.NewProject(3, "api3", 3, "2RP"),
	}

	for _, p := range exampleProjects {
		if err := s.CreateProject(&p); err != nil {
			return err
		}
	}

	return nil
}

func (s *PostgreStorage) GetProjects() ([]*models.Project, error) {
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

func (s *PostgreStorage) GetProject(id uint64) (*models.Project, error) {
	p := &models.Project{}
	return p, s.db.QueryRow(`
		SELECT id, name, semester, company
		FROM projects
		WHERE id = $1
	`, id).Scan(&p.Id, &p.Name, &p.Semester, &p.Company)
}

func (s *PostgreStorage) CreateProject(p *models.Project) error {
	return s.db.QueryRow(`
		INSERT INTO projects (name, semester, company)
		VALUES ($1, $2, $3)
		RETURNING id
	`, p.Name, p.Semester, p.Company,
	).Scan(&p.Id)
}

func (s *PostgreStorage) UpdateProject(p *models.Project) error {
	_, err := s.db.Exec(`
		UPDATE projects
		SET name = $1, semester = $2, company = $3
		WHERE id = $4
	`, p.Name, p.Semester, p.Company, p.Id)
	return err
}

func (s *PostgreStorage) DeleteProject(id uint64) error {
	_, err := s.db.Exec(`
		DELETE FROM projects
		WHERE id = $1
	`, id)
	return err
}

func (s *PostgreStorage) Close() {
	s.db.Close()
}
