package storage

import (
	"database/sql"

	"github.com/ztrue/tracerr"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgreStorage struct {
	postgreProjectModule      *PostgreProjectModule
	postgreUserModule         *PostgreUserModule
	postgreTeamModule         *PostgreTeamModule
	postgreSkillModule        *PostgreSkillModule
	postgreContributionModule *PostgreContributionModule
	db                        *gorm.DB
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
		return nil, tracerr.Errorf("failed to connect database: %w", tracerr.Wrap(err))
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

	postgreContributionModule, err := NewPostgreContributionModule(db)
	if err != nil {
		return nil, tracerr.Errorf("failed to create postgreContributionModule: %w", tracerr.Wrap(err))
	}

	return &PostgreStorage{
		postgreProjectModule:      postgreProjectModule,
		postgreUserModule:         postgreUserModule,
		postgreTeamModule:         postgreTeamModule,
		postgreSkillModule:        postgreSkillModule,
		postgreContributionModule: postgreContributionModule,
		db:                        db,
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

func (s *PostgreStorage) GetContributionModule() (ContributionStorageModule, error) {
	if s.postgreContributionModule.db == nil {
		return nil, tracerr.Errorf("contributionModule not found")
	}
	return s.postgreContributionModule, nil
}

func (s *PostgreStorage) GetRawDB() (*sql.DB, error) {
	return s.db.DB()
}
