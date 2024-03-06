package storage

import (
	"fmt"
	"models"

	"github.com/jmoiron/sqlx"
)

type PostgreTeamModule struct {
	db *sqlx.DB
}

func NewPostgreTeamModule(db *sqlx.DB) (*PostgreTeamModule, error) {
	return &PostgreTeamModule{db: db}, nil
}

func (s *PostgreTeamModule) Migrate() error {
	if _, err := s.db.Exec(`
		CREATE TABLE IF NOT EXISTS teams (
			id SERIAL PRIMARY KEY,
			name VARCHAR(50) NOT NULL
		)

		CREATE TABLE IF NOT EXISTS team_users (
			team_id INT NOT NULL,
			user_id INT NOT NULL,
			PRIMARY KEY (team_id, user_id),
			FOREIGN KEY (team_id) REFERENCES teams(id),
			FOREIGN KEY (user_id) REFERENCES users(id)
		)
	`); err != nil {
		return fmt.Errorf("failed to create teams table: %w", err)
	}

	exampleTeams := []models.CreateTeam{
		models.NewCreateTeam("Khali"),
	}

	for _, p := range exampleTeams {
		if _, err := s.Create(p); err != nil {
			return fmt.Errorf("failed to insert team seeds: %w", err)
		}
	}

	return nil
}

func (s *PostgreTeamModule) Get() ([]*models.Team, error) {
	rows, err := s.db.Query(`
		SELECT id, name
		FROM teams
	`)
	if err != nil {
		return nil, fmt.Errorf("failed to get teams: %w", err)
	}
	defer rows.Close()

	teams := []*models.Team{}
	for rows.Next() {
		var t models.Team
		if err := rows.Scan(&t.Id, &t.Name); err != nil {
			return nil, fmt.Errorf("failed to scan team: %w", err)
		}
		teams = append(teams, &t)
	}

	return teams, nil
}

func (s *PostgreTeamModule) GetById(id uint64) (*models.Team, error) {
	var t models.Team
	if err := s.db.Get(&t, "SELECT id, name FROM teams WHERE id = $1", id); err != nil {
		return nil, fmt.Errorf("failed to get team by id: %w", err)
	}
	return &t, nil
}

func (s *PostgreTeamModule) Create(t models.CreateTeam) (*uint64, error) {
	var id uint64
	if err := s.db.QueryRow(`
		INSERT INTO teams (name)
		VALUES ($1)
		RETURNING id
	`, t.Name).Scan(&id); err != nil {
		return nil, fmt.Errorf("failed to create team: %w", err)
	}
	return &id, nil
}

func (s *PostgreTeamModule) AddUsers(teamId uint64, userIds ...uint64) error {
	query := "INSERT INTO team_users (team_id, user_id) VALUES"
	for i, userId := range userIds {
		query += fmt.Sprintf("(%d, %d)", teamId, userId)
		if i != len(userIds)-1 {
			query += ","
		}
	}

	if _, err := s.db.Exec(query); err != nil {
		return fmt.Errorf("failed to add users to team: %w", err)
	}
	return nil
}

func (s *PostgreTeamModule) RemoveUsers(teamId uint64, userIds ...uint64) error {
	query := "DELETE FROM team_users WHERE"
	for i, userId := range userIds {
		query += fmt.Sprintf(" (team_id = %d AND user_id = %d)", teamId, userId)
		if i != len(userIds)-1 {
			query += " OR"
		}
	}

	if _, err := s.db.Exec(query); err != nil {
		return fmt.Errorf("failed to remove users from team: %w", err)
	}
	return nil
}

func (s *PostgreTeamModule) Update(t models.Team) error {
	if _, err := s.db.Exec(`
		UPDATE teams
		SET name = $1
		WHERE id = $2
	`, t.Name, t.Id); err != nil {
		return fmt.Errorf("failed to update team: %w", err)
	}
	return nil
}

func (s *PostgreTeamModule) Delete(id uint64) error {
	if _, err := s.db.Exec("DELETE FROM teams WHERE id = $1", id); err != nil {
		return fmt.Errorf("failed to delete team: %w", err)
	}
	return nil
}
