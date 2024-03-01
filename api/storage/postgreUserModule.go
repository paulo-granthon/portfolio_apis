package storage

import (
	"database/sql"
	"models"
)

type PostgreUserModule struct {
	db *sql.DB
}

func NewPostgreUserModule(db *sql.DB) (*PostgreUserModule, error) {
	return &PostgreUserModule{db: db}, nil
}

func (s *PostgreUserModule) Migrate() error {
	_, err := s.db.Exec(`
		CREATE TYPE YearSemester AS (
			year     uint2,
			yearSemester uint1
		)

		CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			name VARCHAR(50) NOT NULL,
			summary VARCHAR(100) NULL,
			yearSemester YearSemester NULL,
			githubUsername VARCHAR(39) NULL
			password VARCHAR(50) NOT NULL
		)
	`)
	if err != nil {
		return err
	}

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
			return err
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
		return nil, err
	}
	defer rows.Close()

	var users []*models.User
	for rows.Next() {
		p := &models.User{}
		if err := rows.Scan(
			&p.Id,
			&p.Name,
			&p.Summary,
			&p.SemesterMatriculed,
			&p.GithubUsername,
		); err != nil {
			return nil, err
		}
		users = append(users, p)
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
	if err := s.db.QueryRow(`
		INSERT INTO users (name, password, summary, yearSemester, githubUsername)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id
	`, p.Name, p.Password, p.Summary, p.SemesterMatriculed, p.GithubUsername,
	).Scan(&id); err != nil {
		return nil, err
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
		return nil, err
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
