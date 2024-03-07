package storage

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgreStorage struct {
	postgreProjectModule *PostgreProjectModule
	postgreUserModule    *PostgreUserModule
	postgreTeamModule    *PostgreTeamModule
	db                   *gorm.DB
}

func NewPostgreStorage() (*PostgreStorage, error) {
	databaseCredentials, err := NewDatabaseCredentials()
	if err != nil {
		return nil, err
	}

	connectionString := databaseCredentials.GetConnectionString()

	db, err := gorm.Open(
		postgres.Open(
			connectionString,
		),
		&gorm.Config{},
	)
	if err != nil {
		log.Fatalln(err)
	}

	postgreProjectModule, err := NewPostgreProjectModule(db)
	if err != nil {
		return nil, err
	}

	postgreUserModule, err := NewPostgreUserModule(db)
	if err != nil {
		return nil, err
	}

	postgreTeamModule, err := NewPostgreTeamModule(db)
	if err != nil {
		return nil, err
	}

	return &PostgreStorage{
		postgreProjectModule: postgreProjectModule,
		postgreUserModule:    postgreUserModule,
		postgreTeamModule:    postgreTeamModule,
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

func (s *PostgreStorage) GetTeamModule() (TeamStorageModule, error,
) {
	if s.postgreTeamModule.db == nil {
		return nil, fmt.Errorf("teamModule not found")
	}
	return s.postgreTeamModule, nil
}

func (s *PostgreStorage) Migrate() error {
	rawDB, err := s.db.DB()
	if err != nil {
		fmt.Println("PostgreStorage.Migrate: error getting sql.DB from gorm", err)
		return err
	}

	if _, err := rawDB.Exec(`
		CREATE EXTENSION IF NOT EXISTS uint;

		DROP TABLE IF EXISTS users CASCADE;
		DROP TABLE IF EXISTS teams CASCADE;
		DROP TABLE IF EXISTS team_users CASCADE;
		DROP TABLE IF EXISTS projects CASCADE;

		CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			name VARCHAR(50) NOT NULL,
			summary VARCHAR(200) NULL,
			semester_matriculed JSONB NULL,
			github_username VARCHAR(39) NULL,
			password VARCHAR(50) NOT NULL
		);

		CREATE TABLE IF NOT EXISTS teams (
			id SERIAL PRIMARY KEY,
			name VARCHAR(50) NOT NULL
		);

		CREATE TABLE IF NOT EXISTS team_users (
			team_id INT NOT NULL,
			user_id INT NOT NULL,
			PRIMARY KEY (team_id, user_id),
			FOREIGN KEY (team_id) REFERENCES teams(id),
			FOREIGN KEY (user_id) REFERENCES users(id)
		);

		CREATE TABLE IF NOT EXISTS projects (
			id SERIAL PRIMARY KEY,
			name VARCHAR(50) NOT NULL,
			semester UINT1 NOT NULL,
			company VARCHAR(100) NOT NULL,
			team_id INT NOT NULL,
			summary TEXT NOT NULL,
			url VARCHAR(100) NOT NULL,
			FOREIGN KEY (team_id) REFERENCES teams(id)
		);

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

	teamModule, err := s.GetTeamModule()
	if err != nil {
		fmt.Println("PostgreStorage.Migrate: error getting teamModule", err)
		return err
	}

	if err := teamModule.Migrate(); err != nil {
		fmt.Println("PostgreStorage.Migrate: error migrating teamModule", err)
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
