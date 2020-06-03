package mock

import "github.com/YuChaoGithub/YARC/backend/app/models"

// SearchModel is a mock model for search.
type SearchModel struct{}

// GetResult TODO.
func (m *SearchModel) GetResult(query string) (models.SearchResults, error) {
	return models.SearchResults{}, nil
}
