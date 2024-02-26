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
