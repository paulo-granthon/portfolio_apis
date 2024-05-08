package storage

import (
	"models"

	"github.com/ztrue/tracerr"
	"gorm.io/gorm"
)

type PostgreUserModule struct {
	db *gorm.DB
}

func NewPostgreUserModule(db *gorm.DB) (*PostgreUserModule, error) {
	return &PostgreUserModule{db: db}, nil
}

func (s *PostgreUserModule) Get() ([]models.User, error) {
	var users []models.User
	s.db.Find(&users)
	return users, nil
}

func (s *PostgreUserModule) GetById(id uint64) (*models.User, error) {
	var user models.User
	if err := s.db.First(&user, &id).Error; err != nil {
		return nil, tracerr.Errorf("failed to get user by id: %w", tracerr.Wrap(err))
	}
	return &user, nil
}

func (s *PostgreUserModule) GetByName(name string) (*models.User, error) {
	var user models.User
	if err := s.db.Where("name = ?", name).First(&user).Error; err != nil {
		return nil, tracerr.Errorf("failed to get user by name: %w", tracerr.Wrap(err))
	}
	return &user, nil
}

func (s *PostgreUserModule) GetByUsername(username string) (*models.User, error) {
	var user models.User
	if err := s.db.Where("github_username = ?", username).First(&user).Error; err != nil {
		return nil, tracerr.Errorf("failed to get user by github username: %w", tracerr.Wrap(err))
	}
	return &user, nil
}

func (s *PostgreUserModule) Create(p models.CreateUser) (*uint64, error) {
	user := models.FullUser{
		Name:               p.Name,
		Password:           p.Password,
		Summary:            p.Summary,
		SemesterMatriculed: p.SemesterMatriculed,
		GithubUsername:     p.GithubUsername,
	}

	if err := s.db.Table("users").Create(&user).Error; err != nil {
		return nil, tracerr.Errorf("failed to create user: %w", tracerr.Wrap(err))
	}
	return &user.Id, nil
}

func (s *PostgreUserModule) Register(p models.RegisterUser) (*uint64, error) {
	user := models.FullUser{
		Name:     p.Name,
		Password: p.Password,
	}

	if err := s.db.Create(&user).Error; err != nil {
		return nil, tracerr.Errorf("failed to register user: %w", tracerr.Wrap(err))
	}
	return &user.Id, nil
}

func (s *PostgreUserModule) Update(p models.UpdateUser) error {
	if err := s.db.Model(&models.FullUser{}).Where("id = ?", p.Id).Updates(&p).Error; err != nil {
		return tracerr.Errorf("failed to update user: %w", tracerr.Wrap(err))
	}
	return nil
}

func (s *PostgreUserModule) Delete(id uint64) error {
	if err := s.db.Delete(&models.FullUser{}, id).Error; err != nil {
		return tracerr.Errorf("failed to delete user: %w", tracerr.Wrap(err))
	}
	return nil
}
