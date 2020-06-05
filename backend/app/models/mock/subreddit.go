package mock

import "github.com/YuChaoGithub/YARC/backend/app/models"

// SubredditModel is a mock model for subreddit.
type SubredditModel struct{}

// Insert always returns nil.
func (m *SubredditModel) Insert(name, description string) error {
	return nil
}

// Get only returns a result when name is "radiohead".
func (m *SubredditModel) Get(name string) (models.SubredditInfo, error) {
	if name == "radiohead" {
		return models.SubredditInfo{
			Name:        "radiohead",
			Description: "Dedicated to all human beings.",
		}, nil
	}
	return models.SubredditInfo{}, models.ErrSubredditNotExist
}

// GetTrending TODO.
func (m *SubredditModel) GetTrending(limit int) ([]models.SubredditInfo, error) {
	return []models.SubredditInfo{
		{
			Name:        "radiohead",
			Description: "Dedicated to all human beings.",
		},
		{
			Name:        "underrated",
			Description: "Let down and hanging around.",
		},
	}, nil
}
