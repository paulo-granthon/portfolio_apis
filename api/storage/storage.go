package storage

import (
	"models"

	_ "github.com/lib/pq"
)

type Storage interface {
	Migrate() error
	GetProjectModule() (ProjectStorageModule, error)
	GetUserModule() (UserStorageModule, error)
	GetTeamModule() (TeamStorageModule, error)
}

type StorageModule[T any, TCreate any, TUpdate any] interface {
	Migrate() error
	Get() ([]T, error)
	GetById(uint64) (*T, error)
	Create(TCreate) (*uint64, error)
	Update(TUpdate) error
	Delete(uint64) error
}

type ProjectStorageModule interface {
	Migrate() error
	Get() ([]models.Project, error)
	GetById(uint64) (*models.Project, error)
	GetByUserId(uint64) ([]models.Project, error)
	GetByTeamId(uint64) ([]models.Project, error)
	Create(models.CreateProject) (*uint64, error)
	Update(models.UpdateProject) error
	Delete(uint64) error
}

type UserStorageModule interface {
	Migrate() error
	Get() ([]models.User, error)
	GetById(uint64) (*models.User, error)
	Create(models.CreateUser) (*uint64, error)
	Register(models.RegisterUser) (*uint64, error)
	Update(models.UpdateUser) error
	Delete(uint64) error
}

type TeamStorageModule interface {
	Migrate() error
	Get() ([]models.Team, error)
	GetById(uint64) (*models.Team, error)
	GetUsers(uint64) ([]models.User, error)
	AddUsers(uint64, ...uint64) error
	RemoveUsers(uint64, ...uint64) error
	Create(models.CreateTeam) (*uint64, error)
	Update(models.Team) error
	Delete(uint64) error
}
