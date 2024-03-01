package storage

import (
	"models"

	_ "github.com/lib/pq"
)

type Storage interface {
	Migrate() error
	GetProjectModule() (StorageModule[models.Project, models.CreateProject], error)
	GetUserModule() (UserStorageModule, error)
}

type StorageModule[T any, TCreate any] interface {
	Migrate() error
	Get() ([]*T, error)
	GetById(uint64) (*T, error)
	Create(TCreate) (*uint64, error)
	Update(T) error
	Delete(uint64) error
}

type UserStorageModule interface {
	Migrate() error
	Get() ([]*models.User, error)
	GetById(uint64) (*models.User, error)
	Create(models.CreateUser) (*uint64, error)
	Register(models.RegisterUser) (*uint64, error)
	Update(models.UpdateUser) error
	Delete(uint64) error
}
