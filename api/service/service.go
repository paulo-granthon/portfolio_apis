package service

import (
	"storage"

	"github.com/ztrue/tracerr"
)

type Service struct {
	ContributionService *ContributionService
}

func NewService(s storage.Storage) (*Service, error) {
	contributionService, err := NewContributionService(s)
	if err != nil {
		return nil, tracerr.Errorf("failed to create contribution service: %v", err)
	}

	return &Service{
		ContributionService: contributionService,
	}, nil
}
