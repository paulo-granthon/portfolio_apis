package storage

import (
	"encoding/json"
	"fmt"
	"models"

	"github.com/jmoiron/sqlx"
)

type PostgreUserModule struct {
	db *sqlx.DB
}

func NewPostgreUserModule(db *sqlx.DB) (*PostgreUserModule, error) {
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

func (s *PostgreUserModule) Get() ([]*models.User, error) {
	rows, err := s.db.Query(`
		SELECT id, name, summary, yearSemester, githubUsername
		FROM users
	`)
	if err != nil {
		return nil, fmt.Errorf("failed to get users: %w", err)
	}
	defer rows.Close()

	var users []*models.User
	for rows.Next() {
		user := &models.User{}

		if err := rows.Scan(
			&user.Id,
			&user.Name,
			&user.Summary,
			&user.SemesterMatriculed,
			&user.GithubUsername,
		); err != nil {
			return nil, fmt.Errorf("failed to scan user: %w", err)
		}

		users = append(users, user)
	}

	return users, nil
}

func (s *PostgreUserModule) GetById(id uint64) (*models.User, error) {
	p := &models.User{}
	return p, s.db.QueryRow(`
		SELECT id, name, summary, yearSemester, githubUsername
		FROM users
		WHERE id = $1
	`, id).Scan(&p.Id, &p.Name, &p.Summary, &p.SemesterMatriculed, &p.GithubUsername)
}

func (s *PostgreUserModule) Create(p models.CreateUser) (*uint64, error) {
	var id uint64
	semesterMatriculed, err := json.Marshal(p.SemesterMatriculed)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal yearSemester: %w", err)
	}

	if err := s.db.QueryRow(`
		INSERT INTO users (name, password, summary, yearSemester, githubUsername)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id
	`, p.Name, p.Password, p.Summary, semesterMatriculed, p.GithubUsername,
	).Scan(&id); err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	return &id, nil
}

func (s *PostgreUserModule) Register(p models.RegisterUser) (*uint64, error) {
	var id uint64
	if err := s.db.QueryRow(`
		INSERT INTO users (name, password)
		VALUES ($1, $2)
		RETURNING id
	`, p.Name, p.Password,
	).Scan(&id); err != nil {
		return nil, fmt.Errorf("failed to register user: %w", err)
	}

	return &id, nil
}

func (s *PostgreUserModule) Update(p models.UpdateUser) error {
	_, err := s.db.Exec(`
		UPDATE users
		SET name = $1, summary = $2, yearSemester = $3, githubUsername = $4
		WHERE id = $5
	`, p.Name, p.Summary, p.SemesterMatriculed, p.GithubUsername, p.Id)
	return err
}

func (s *PostgreUserModule) Delete(id uint64) error {
	_, err := s.db.Exec(`
		DELETE FROM users
		WHERE id = $1
	`, id)
	return err
}

func (s *PostgreUserModule) Close() {
	s.db.Close()
}
