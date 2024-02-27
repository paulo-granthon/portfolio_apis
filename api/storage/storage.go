package storage

import (
	"models"

	_ "github.com/lib/pq"
)

type Storage interface {
	Migrate() error
	GetProjects() ([]*models.Project, error)
	GetProject(uint64) (*models.Project, error)
	CreateProject(*models.Project) error
	UpdateProject(*models.Project) error
	DeleteProject(uint64) error
}
