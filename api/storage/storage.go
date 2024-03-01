package storage

import (
	"models"

	_ "github.com/lib/pq"
)

type Storage interface {
	Migrate() error
	GetProjectModule() (StorageModule[models.Project, models.CreateProject], error)
}

type StorageModule[T any, TCreate any] interface {
	Migrate() error
	GetProjects() ([]*T, error)
	GetProject(uint64) (*T, error)
	CreateProject(TCreate) (*uint64, error)
	UpdateProject(T) error
	DeleteProject(uint64) error
}
