package storage

import (
	"errors"

	"github.com/paulo-granthon/portfolio_apis/models"

	"github.com/ztrue/tracerr"
	"gorm.io/gorm"
)

type PostgreParticipationModule struct {
	db *gorm.DB
}

func NewPostgreParticipationModule(db *gorm.DB) (*PostgreParticipationModule, error) {
	return &PostgreParticipationModule{db: db}, nil
}

// GetByUserAndProject returns the participation summary for a (user, project)
// pair, or nil when none has been recorded (so callers can render projects
// without a summary cleanly).
func (s *PostgreParticipationModule) GetByUserAndProject(userId uint64, projectId uint64) (*models.Participation, error) {
	var participation models.Participation
	err := s.db.
		Table("participations").
		Where("user_id = ? AND project_id = ?", userId, projectId).
		First(&participation).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, tracerr.Errorf("failed to get participation: %w", tracerr.Wrap(err))
	}
	return &participation, nil
}

func (s *PostgreParticipationModule) Create(p models.CreateParticipation) (*uint64, error) {
	participation := models.Participation{
		ProjectId: p.ProjectId,
		UserId:    p.UserId,
		Summary:   p.Summary,
	}
	if err := s.db.Create(&participation).Error; err != nil {
		return nil, tracerr.Errorf("failed to create participation: %w", tracerr.Wrap(err))
	}
	return &participation.Id, nil
}
