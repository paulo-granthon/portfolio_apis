package storage

import (
	"github.com/ztrue/tracerr"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgreStorage struct {
	postgreProjectModule *PostgreProjectModule
	postgreUserModule    *PostgreUserModule
	postgreTeamModule    *PostgreTeamModule
	postgreSkillModule   *PostgreSkillModule
	postgreNoteModule    *PostgreNoteModule
	db                   *gorm.DB
}

func NewPostgreStorage() (*PostgreStorage, error) {
	databaseCredentials, err := NewDatabaseCredentials()
	if err != nil {
		return nil, tracerr.Errorf("failed to create databaseCredentials: %w", tracerr.Wrap(err))
	}

	connectionString := databaseCredentials.GetConnectionString()

	db, err := gorm.Open(
		postgres.Open(
			connectionString,
		),
		&gorm.Config{},
	)
	if err != nil {
		return nil, tracerr.Errorf("failed to connect database", tracerr.Wrap(err))
	}

	postgreProjectModule, err := NewPostgreProjectModule(db)
	if err != nil {
		return nil, tracerr.Errorf("failed to create postgreProjectModule: %w", tracerr.Wrap(err))
	}

	postgreUserModule, err := NewPostgreUserModule(db)
	if err != nil {
		return nil, tracerr.Errorf("failed to create postgreUserModule: %w", tracerr.Wrap(err))
	}

	postgreTeamModule, err := NewPostgreTeamModule(db)
	if err != nil {
		return nil, tracerr.Errorf("failed to create postgreTeamModule: %w", tracerr.Wrap(err))
	}

	postgreSkillModule, err := NewPostgreSkillModule(db)
	if err != nil {
		return nil, tracerr.Errorf("failed to create postgreSkillModule: %w", tracerr.Wrap(err))
	}

	postgreNoteModule, err := NewPostgreNoteModule(db)
	if err != nil {
		return nil, tracerr.Errorf("failed to create postgreNoteModule: %w", tracerr.Wrap(err))
	}

	return &PostgreStorage{
		postgreProjectModule: postgreProjectModule,
		postgreUserModule:    postgreUserModule,
		postgreTeamModule:    postgreTeamModule,
		postgreSkillModule:   postgreSkillModule,
		postgreNoteModule:    postgreNoteModule,
		db:                   db,
	}, nil
}

func (s *PostgreStorage) GetProjectModule() (ProjectStorageModule, error) {
	if s.postgreProjectModule.db == nil {
		return nil, tracerr.Errorf("projectModule not found")
	}
	return s.postgreProjectModule, nil
}

func (s *PostgreStorage) GetUserModule() (UserStorageModule, error) {
	if s.postgreUserModule.db == nil {
		return nil, tracerr.Errorf("userModule not found")
	}
	return s.postgreUserModule, nil
}

func (s *PostgreStorage) GetTeamModule() (TeamStorageModule, error) {
	if s.postgreTeamModule.db == nil {
		return nil, tracerr.Errorf("teamModule not found")
	}
	return s.postgreTeamModule, nil
}

func (s *PostgreStorage) GetSkillModule() (SkillStorageModule, error) {
	if s.postgreSkillModule.db == nil {
		return nil, tracerr.Errorf("skillModule not found")
	}
	return s.postgreSkillModule, nil
}

func (s *PostgreStorage) GetNoteModule() (NoteStorageModule, error) {
	if s.postgreNoteModule.db == nil {
		return nil, tracerr.Errorf("noteModule not found")
	}
	return s.postgreNoteModule, nil
}

func (s *PostgreStorage) Migrate() error {
	rawDB, err := s.db.DB()
	if err != nil {
		err = tracerr.Errorf("PostgreStorage.Migrate: error getting sql.DB from gorm: %w", err)
		tracerr.PrintSourceColor(err)
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
		err = tracerr.Errorf("PostgreStorage.Migrate: error executing root migration: %w", err)
		tracerr.PrintSourceColor(err)
		return err
	}

	userModule, err := s.GetUserModule()
	if err != nil {
		err = tracerr.Errorf("PostgreStorage.Migrate: error getting userModule: %w", err)
		tracerr.PrintSourceColor(err)
		return err
	}

	if err := userModule.Migrate(); err != nil {
		err = tracerr.Errorf("PostgreStorage.Migrate: error migrating userModule: %w", err)
		tracerr.PrintSourceColor(err)
		return err
	}

	teamModule, err := s.GetTeamModule()
	if err != nil {
		err = tracerr.Errorf("PostgreStorage.Migrate: error getting teamModule: %w", err)
		tracerr.PrintSourceColor(err)
		return err
	}

	if err := teamModule.Migrate(); err != nil {
		err = tracerr.Errorf("PostgreStorage.Migrate: error migrating teamModule: %w", err)
		tracerr.PrintSourceColor(err)
		return err
	}

	projectModule, err := s.GetProjectModule()
	if err != nil {
		err = tracerr.Errorf("PostgreStorage.Migrate: error getting projectModule: %w", err)
		tracerr.PrintSourceColor(err)
		return err
	}

	if err := projectModule.Migrate(); err != nil {
		err = tracerr.Errorf("PostgreStorage.Migrate: error migrating projectModule: %w", err)
		tracerr.PrintSourceColor(err)
		return err
	}

	return nil
}
