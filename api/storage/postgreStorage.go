package storage

import (
	"fmt"
	"log"
	"models"

	"github.com/jmoiron/sqlx"
)

type PostgreStorage struct {
	postgreProjectModule *PostgreProjectModule
	postgreUserModule    *PostgreUserModule
	db                   *sqlx.DB
}

func NewPostgreStorage() (*PostgreStorage, error) {
	db, err := sqlx.Connect("postgres", "port=3332 user=postgres password=secret sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	postgreProjectModule, err := NewPostgreProjectModule(db)
	if err != nil {
		return nil, err
	}

	postgreUserModule, err := NewPostgreUserModule(db)
	if err != nil {
		return nil, err
	}

	return &PostgreStorage{
		postgreProjectModule: postgreProjectModule,
		postgreUserModule:    postgreUserModule,
		db:                   db,
	}, nil
}

func (s *PostgreStorage) GetProjectModule() (StorageModule[models.Project, models.CreateProject], error) {
	if s.postgreProjectModule.db == nil {
		return nil, fmt.Errorf("projectModule not found")
	}
	return s.postgreProjectModule, nil
}

func (s *PostgreStorage) GetUserModule() (UserStorageModule, error) {
	if s.postgreUserModule.db == nil {
		return nil, fmt.Errorf("userModule not found")
	}
	return s.postgreUserModule, nil
}

func (s *PostgreStorage) Migrate() error {
	if _, err := s.db.Exec(`
		CREATE EXTENSION IF NOT EXISTS uint;
	`); err != nil {
		fmt.Println("PostgreStorage.Migrate: error executing root migration", err)
		return err
	}

	userModule, err := s.GetUserModule()
	if err != nil {
		fmt.Println("PostgreStorage.Migrate: error getting userModule", err)
		return err
	}

	if err := userModule.Migrate(); err != nil {
		fmt.Println("PostgreStorage.Migrate: error migrating userModule", err)
		return err
	}

	projectModule, err := s.GetProjectModule()
	if err != nil {
		fmt.Println("PostgreStorage.Migrate: error getting projectModule", err)
		return err
	}

	if err := projectModule.Migrate(); err != nil {
		fmt.Println("PostgreStorage.Migrate: error migrating projectModule", err)
		return err
	}

	return nil
}
