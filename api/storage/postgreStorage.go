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
	postgreTeamModule    *PostgreTeamModule
	db                   *sqlx.DB
}

func NewPostgreStorage() (*PostgreStorage, error) {
	databaseCredentials, err := NewDatabaseCredentials()
	if err != nil {
		return nil, err
	}

	connectionString := databaseCredentials.GetConnectionString()

	db, err := sqlx.Connect("postgres", connectionString)
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

func (s *PostgreStorage) GetProjectModule() (ProjectStorageModule, error) {
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

func (s *PostgreStorage) GetTeamModule() (
	StorageModule[models.Team, models.CreateTeam, models.Team],
	error,
) {
	if s.postgreTeamModule.db == nil {
		return nil, fmt.Errorf("teamModule not found")
	}
	return s.postgreTeamModule, nil
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

	teamModule, err := s.GetTeamModule()
	if err != nil {
		fmt.Println("PostgreStorage.Migrate: error getting teamModule", err)
		return err
	}

	if err := teamModule.Migrate(); err != nil {
		fmt.Println("PostgreStorage.Migrate: error migrating teamModule", err)
		return err
	}

	return nil
}
