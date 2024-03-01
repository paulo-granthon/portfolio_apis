package storage

import (
	"models"

	_ "github.com/lib/pq"
)

type Storage interface {
	Migrate() error
	GetProjectModule() (StorageModule[models.Project, models.CreateProject], error)
	GetUserModule() (StorageModule[models.User, models.CreateUser], error)
}

type StorageModule[T any, TCreate any] interface {
	Migrate() error
	Get() ([]*T, error)
	GetById(uint64) (*T, error)
	Create(TCreate) (*uint64, error)
	Update(T) error
	Delete(uint64) error
}
