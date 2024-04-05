package storage

import (
	"models"

	"gorm.io/gorm"
)

type PostgreSkillModule struct {
	db *gorm.DB
}

func NewPostgreSkillModule(db *gorm.DB) (*PostgreSkillModule, error) {
	return &PostgreSkillModule{db: db}, nil
}

func (s *PostgreSkillModule) Get() ([]models.Skill, error) {
	var skills []models.Skill
	s.db.Find(&skills)
	return skills, nil
}

func (s *PostgreSkillModule) GetById(id uint64) (*models.Skill, error) {
	var skill models.Skill
	if err := s.db.First(&skill, id).Error; err != nil {
		return nil, err
	}
	return &skill, nil
}

func (s *PostgreSkillModule) GetByName(name string) (*models.Skill, error) {
	var skill models.Skill
	if err := s.db.Where("name = ?", name).First(&skill).Error; err != nil {
		return nil, err
	}
	return &skill, nil
}

func (s *PostgreSkillModule) Create(sk models.CreateSkill) (*uint64, error) {
	skill := models.Skill{Name: sk.Name}
	if err := s.db.Create(&skill).Error; err != nil {
		return nil, err
	}
	return &skill.Id, nil
}

func (s *PostgreSkillModule) Delete(id uint64) error {
	if err := s.db.Delete(&models.Skill{}, id).Error; err != nil {
		return err
	}
	return nil
}
