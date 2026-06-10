package endpoints

import (
	"github.com/paulo-granthon/portfolio_apis/server"
)

func CreateEndpoints() []server.Endpoint {
	return mergeEndpoints(
		ProjectEndpoints(),
		RootEndpoints(),
		UserEndpoints(),
		TeamEndpoints(),
		ContributionEndpoints(),
	)
}

func mergeEndpoints(endpoints ...[]server.Endpoint) []server.Endpoint {
	var merged []server.Endpoint
	for _, e := range endpoints {
		merged = append(merged, e...)
	}
	return merged
}
