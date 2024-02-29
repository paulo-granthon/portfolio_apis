package storage

import (
	"models"

	_ "github.com/lib/pq"
)

type Storage interface {
	Migrate() error
	GetProjectModule() (ProjectModule, error)
}

type ProjectModule interface {
	Migrate() error
	GetProjects() ([]*models.Project, error)
	GetProject(uint64) (*models.Project, error)
	CreateProject(models.CreateProject) (*uint64, error)
	UpdateProject(models.Project) error
	DeleteProject(uint64) error
}
