package mock

import (
	"time"

	"github.com/YuChaoGithub/YARC/backend/app/models"
)

// SearchModel is a mock model for search.
type SearchModel struct{}

// GetResult only returns result when query="ghost town".
func (m *SearchModel) GetResult(query string) (models.SearchResults, error) {
	if query == "ghost town" {
		return models.SearchResults{
			Subreddits: []models.SubredditInfo{
				{Name: "PHP", Members: 0, Description: "$_POST['cool']"},
				{Name: "golang", Members: 0, Description: "if err != nil"},
			},
			Articles: []models.ArticleInfo{
				{
					Subreddit:  "Python",
					ArticleID:  "1",
					Type:       "text",
					Title:      "title",
					Points:     3,
					Comments:   1,
					PostedBy:   "Collin",
					PostedTime: time.Date(1995, 11, 24, 0, 0, 0, 0, time.UTC),
				},
			},
		}, nil
	}

	return models.SearchResults{}, nil
}
