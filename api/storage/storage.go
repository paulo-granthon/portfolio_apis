package storage

import (
	"models"

	_ "github.com/lib/pq"
)

type Storage interface {
	GetProjectModule() (ProjectStorageModule, error)
	GetUserModule() (UserStorageModule, error)
	GetTeamModule() (TeamStorageModule, error)
	GetSkillModule() (SkillStorageModule, error)
	GetNoteModule() (NoteStorageModule, error)
}

type StorageModule[T any, TCreate any, TUpdate any] interface {
	Get() ([]T, error)
	GetById(uint64) (*T, error)
	GetByName(string) (*T, error)
	Create(TCreate) (*uint64, error)
	Update(TUpdate) error
	Delete(uint64) error
}

type SkillStorageModule interface {
	Get() ([]models.Skill, error)
	GetById(uint64) (*models.Skill, error)
	GetByName(string) (*models.Skill, error)
	Create(models.CreateSkill) (*uint64, error)
	Delete(uint64) error
}

type NoteStorageModule interface {
	Get() ([]models.Note, error)
	GetById(uint64) (*models.Note, error)
	GetFilter(models.NoteFilter) ([]models.NoteDetail, error)
	GetByProjectId(uint64) ([]models.Note, error)
	GetBySkillId(uint64) ([]models.Note, error)
	GetByUserId(uint64) ([]models.Note, error)
	Create(models.CreateNote) (*uint64, error)
	Delete(uint64) error
}

type ProjectStorageModule interface {
	Get() ([]models.Project, error)
	GetById(uint64) (*models.Project, error)
	GetByName(string) (*models.Project, error)
	GetByUserId(uint64) ([]models.Project, error)
	GetByTeamId(uint64) ([]models.Project, error)
	Create(models.CreateProject) (*uint64, error)
	Update(models.UpdateProject) error
	Delete(uint64) error
}

type UserStorageModule interface {
	Get() ([]models.User, error)
	GetById(uint64) (*models.User, error)
	GetByName(string) (*models.User, error)
	Create(models.CreateUser) (*uint64, error)
	Register(models.RegisterUser) (*uint64, error)
	Update(models.UpdateUser) error
	Delete(uint64) error
}

type TeamStorageModule interface {
	Get() ([]models.Team, error)
	GetById(uint64) (*models.Team, error)
	GetUsers(uint64) ([]models.User, error)
	AddUsers(uint64, ...uint64) error
	RemoveUsers(uint64, ...uint64) error
	Create(models.CreateTeam) (*uint64, error)
	Update(models.Team) error
	Delete(uint64) error
}
