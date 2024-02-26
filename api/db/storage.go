package db

import (
	"database/sql"
	"models"

	_ "github.com/lib/pq"
)

type Storage interface {
	Migrate() error
	CreateProject(*models.Project) error
	GetProject(uint64) (*models.Project, error)
	UpdateProject(*models.Project) error
	DeleteProject(uint64) error
}

type PostgresStorage struct {
	db *sql.DB
}

func NewPostgresStorage() (*PostgresStorage, error) {
	connStr := "user=postgres password=secret sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &PostgresStorage{
		db: db,
	}, nil
}

func (s *PostgresStorage) Migrate() error {
	_, err := s.db.Exec(`
		CREATE TABLE IF NOT EXISTS projects (
			id SERIAL PRIMARY KEY,
			name VARCHAR(50) NOT NULL,
			semester UINT1 NOT NULL,
			company VARCHAR(100) NOT NULL
		)
	`)
	return err
}

func (s *PostgresStorage) CreateProject(p *models.Project) error {
	return nil
}

func (s *PostgresStorage) GetProject(id uint64) (*models.Project, error) {
	return nil, nil
}

func (s *PostgresStorage) UpdateProject(p *models.Project) error {
	return nil
}

func (s *PostgresStorage) DeleteProject(id uint64) error {
	return nil
}

func (s *PostgresStorage) Close() {
	s.db.Close()
}
