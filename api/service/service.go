package service

import (
	"github.com/paulo-granthon/portfolio_apis/storage"
	"github.com/ztrue/tracerr"
)

type Service struct {
	ContributionService *ContributionService
	PortfolioService    *PortfolioService
}

func NewService(s storage.Storage) (*Service, error) {
	contributionService, err := NewContributionService(s)
	if err != nil {
		return nil, tracerr.Errorf("failed to create contribution service: %v", err)
	}

	portfolioService, err := NewPortfolioService(s, contributionService)
	if err != nil {
		return nil, tracerr.Errorf("failed to create portfolio service: %v", err)
	}

	return &Service{
		ContributionService: contributionService,
		PortfolioService:    portfolioService,
	}, nil
}
