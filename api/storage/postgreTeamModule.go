package storage

import (
	"models"

	"github.com/ztrue/tracerr"
	"gorm.io/gorm"
)

type PostgreTeamModule struct {
	db *gorm.DB
}

func NewPostgreTeamModule(db *gorm.DB) (*PostgreTeamModule, error) {
	return &PostgreTeamModule{db: db}, nil
}

func (s *PostgreTeamModule) Migrate() error {
	exampleTeams := []models.CreateTeam{
		models.NewCreateTeam("Khali"),
	}

	var teamIdToAddMember *uint64

	for _, p := range exampleTeams {
		insertedTeamId, err := s.Create(p)
		if err != nil {
			if teamIdToAddMember == nil {
				teamIdToAddMember = insertedTeamId
			}
			return tracerr.Errorf("failed to insert team seeds: %w", tracerr.Wrap(err))
		}
	}

	if teamIdToAddMember == nil {
		existingTeam, err := s.GetById(1)
		if err != nil {
			return tracerr.Errorf("no team was inserted neither found in the database: %w", tracerr.Wrap(err))
		}

		teamIdToAddMember = &existingTeam.Id
	}

	if err := s.AddUsers(*teamIdToAddMember, 1); err != nil {
		return tracerr.Errorf("failed to add users to team: %w", tracerr.Wrap(err))
	}

	return nil
}

func (s *PostgreTeamModule) Get() ([]models.Team, error) {
	var teams []models.Team
	s.db.Find(&teams)
	return teams, nil
}

func (s *PostgreTeamModule) GetById(id uint64) (*models.Team, error) {
	var team models.Team
	if err := s.db.First(&team, &id).Error; err != nil {
		return nil, tracerr.Errorf("failed to get team by id: %w", tracerr.Wrap(err))
	}
	return &team, nil
}

func (s *PostgreTeamModule) GetUsers(id uint64) ([]models.User, error) {
	var users []models.User
	subQuery := s.db.
		Table("team_users").
		Select("user_id").
		Where("team_id = ?", id)
	s.db.
		Where("id IN (?)", subQuery).
		Find(&users)
	return users, nil
}

func (s *PostgreTeamModule) Create(t models.CreateTeam) (*uint64, error) {
	team := models.Team{
		Name: t.Name,
	}

	if err := s.db.Create(&team).Error; err != nil {
		return nil, tracerr.Errorf("failed to create team: %w", tracerr.Wrap(err))
	}
	return &team.Id, nil
}

func (s *PostgreTeamModule) AddUsers(teamId uint64, userIds ...uint64) error {
	var teamUsers []models.TeamUser
	for _, userId := range userIds {
		teamUsers = append(teamUsers, models.TeamUser{
			TeamId: teamId,
			UserId: userId,
		})
	}

	if err := s.db.Create(&teamUsers).Error; err != nil {
		return tracerr.Errorf("failed to add users to team: %w", tracerr.Wrap(err))
	}

	return nil
}

func (s *PostgreTeamModule) RemoveUsers(teamId uint64, userIds ...uint64) error {
	var teamUsers []models.TeamUser
	for _, userId := range userIds {
		teamUsers = append(teamUsers, models.TeamUser{
			TeamId: teamId,
			UserId: userId,
		})
	}

	if err := s.db.Delete(&teamUsers).Error; err != nil {
		return tracerr.Errorf("failed to remove users from team: %w", tracerr.Wrap(err))
	}

	return nil
}

func (s *PostgreTeamModule) Update(t models.Team) error {
	if err := s.db.Model(&models.Team{}).Where("id = ?", t.Id).Updates(&t).Error; err != nil {
		return tracerr.Errorf("failed to update team: %w", tracerr.Wrap(err))
	}
	return nil
}

func (s *PostgreTeamModule) Delete(id uint64) error {
	if err := s.db.Delete(&models.Team{}, id).Error; err != nil {
		return tracerr.Errorf("failed to delete team: %w", tracerr.Wrap(err))
	}
	return nil
}
