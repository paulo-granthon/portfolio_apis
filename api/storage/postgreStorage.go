package storage

import (
	"database/sql"
	"fmt"
)

type PostgreStorage struct {
	postgreProjectModule *PostgreProjectModule
	db                   *sql.DB
}

func NewPostgreStorage() (*PostgreStorage, error) {
	connStr := "user=postgres password=secret sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &PostgreStorage{
		db: db,
	}, nil
}

func (s *PostgreStorage) GetProjectModule() (ProjectModule, error) {
	if s.postgreProjectModule.db == nil {
		return nil, fmt.Errorf("projectModule not found")
	}

	return s.postgreProjectModule, nil
}

func (s *PostgreStorage) Migrate() error {
	projectModule, err := s.GetProjectModule()
	if err != nil {
		return err
	}

	if err := projectModule.Migrate(); err != nil {
		return err
	}

	return nil
}
