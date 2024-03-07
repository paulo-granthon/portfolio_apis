package storage

import (
	"fmt"
	"models"

	"gorm.io/gorm"
)

type PostgreUserModule struct {
	db *gorm.DB
}

func NewPostgreUserModule(db *gorm.DB) (*PostgreUserModule, error) {
	return &PostgreUserModule{db: db}, nil
}

func (s *PostgreUserModule) Migrate() error {
	summary := "Backend developer intern at @gorilainvest | Database technologist student at FATEC | Self titled full-stack developer"
	yearSemesterMatriculed := models.NewYearSemester(uint16(2022), uint8(2))
	githubUsername := "paulo-granthon"

	exampleUsers := []models.CreateUser{
		models.NewCreateUser(
			"Paulo Granthon",
			"123456",
			&summary,
			&yearSemesterMatriculed,
			&githubUsername,
		),
	}

	for _, p := range exampleUsers {
		if _, err := s.Create(p); err != nil {
			return fmt.Errorf("failed to insert user seeds: %w", err)
		}
	}

	return nil
}

func (s *PostgreUserModule) Get() ([]models.User, error) {
	var users []models.User
	s.db.Find(&users)
	return users, nil
}

func (s *PostgreUserModule) GetById(id uint64) (*models.User, error) {
	var user models.User
	if err := s.db.First(&user, &id).Error; err != nil {
		return nil, fmt.Errorf("failed to get user by id: %w", err)
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
		return nil, fmt.Errorf("failed to create user: %w", err)
	}
	return &user.Id, nil
}

func (s *PostgreUserModule) Register(p models.RegisterUser) (*uint64, error) {
	user := models.FullUser{
		Name:     p.Name,
		Password: p.Password,
	}

	if err := s.db.Create(&user).Error; err != nil {
		return nil, fmt.Errorf("failed to register user: %w", err)
	}
	return &user.Id, nil
}

func (s *PostgreUserModule) Update(p models.UpdateUser) error {
	if err := s.db.Model(&models.FullUser{}).Where("id = ?", p.Id).Updates(&p).Error; err != nil {
		return fmt.Errorf("failed to update user: %w", err)
	}
	return nil
}

func (s *PostgreUserModule) Delete(id uint64) error {
	if err := s.db.Delete(&models.FullUser{}, id).Error; err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}
	return nil
}
