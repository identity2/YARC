package mock

import (
	"time"

	"github.com/YuChaoGithub/YARC/backend/app/models"
)

// ArticleModel is the mock model for article.
type ArticleModel struct{}

// Insert always returns "WX-78", nil.
func (m *ArticleModel) Insert(string, string, string, string, string) (string, error) {
	return "WX-78", nil
}

// ModifyBody always returns nil.
func (m *ArticleModel) ModifyBody(string, string, string) error {
	return nil
}

// Delete always returns nil.
func (m *ArticleModel) Delete(string, string) error {
	return nil
}

// Get has only the article with ID WX-78.
func (m *ArticleModel) Get(articleID string) (models.ArticleInfo, error) {
	if articleID == "WX-78" {
		return models.ArticleInfo{
			Subreddit:  "radiohead",
			ArticleID:  "WX-78",
			Type:       "text",
			Body:       "I'm a reasonable man, get off my case.",
			Title:      "Tin can",
			Points:     42,
			PostedBy:   "Bellamy",
			PostedTime: time.Date(2019, time.November, 13, 0, 0, 0, 0, time.UTC),
		}, nil
	}

	return models.ArticleInfo{}, models.ErrArticleNotExist
}

// GetBySubreddit TODO.
func (m *ArticleModel) GetBySubreddit(string, string, string, int) ([]models.ArticleInfo, error) {
	return []models.ArticleInfo{
		{
			Subreddit:  "radiohead",
			ArticleID:  "WX-78",
			Type:       "text",
			Body:       "I'm a reasonable man, get off my case.",
			Title:      "Tin can",
			Points:     42,
			PostedBy:   "Bellamy",
			PostedTime: time.Date(2019, time.November, 13, 0, 0, 0, 0, time.UTC),
		},
		{
			Subreddit:  "theclash",
			ArticleID:  "bRiXtoN",
			Type:       "image",
			Body:       "On your head or trigger?",
			Title:      "On the pavement",
			Points:     66,
			PostedBy:   "Lyndon",
			PostedTime: time.Date(1989, time.June, 4, 0, 0, 0, 0, time.UTC),
		},
	}, nil
}

// GetByUser TODO.
func (m *ArticleModel) GetByUser(string, string, string, int) ([]models.ArticleInfo, error) {
	return []models.ArticleInfo{
		{
			Subreddit:  "radiohead",
			ArticleID:  "WX-78",
			Type:       "text",
			Body:       "I'm a reasonable man, get off my case.",
			Title:      "Tin can",
			Points:     42,
			PostedBy:   "Bellamy",
			PostedTime: time.Date(2019, time.November, 13, 0, 0, 0, 0, time.UTC),
		},
		{
			Subreddit:  "theclash",
			ArticleID:  "bRiXtoN",
			Type:       "image",
			Body:       "On your head or trigger?",
			Title:      "On the pavement",
			Points:     66,
			PostedBy:   "Lyndon",
			PostedTime: time.Date(1989, time.June, 4, 0, 0, 0, 0, time.UTC),
		},
	}, nil
}

// GetSavedByUser TODO.
func (m *ArticleModel) GetSavedByUser(string, string, string, int) ([]models.ArticleInfo, error) {
	return []models.ArticleInfo{
		{
			Subreddit:  "radiohead",
			ArticleID:  "WX-78",
			Type:       "text",
			Body:       "I'm a reasonable man, get off my case.",
			Title:      "Tin can",
			Points:     42,
			PostedBy:   "Bellamy",
			PostedTime: time.Date(2019, time.November, 13, 0, 0, 0, 0, time.UTC),
		},
		{
			Subreddit:  "theclash",
			ArticleID:  "bRiXtoN",
			Type:       "image",
			Body:       "On your head or trigger?",
			Title:      "On the pavement",
			Points:     66,
			PostedBy:   "Lyndon",
			PostedTime: time.Date(1989, time.June, 4, 0, 0, 0, 0, time.UTC),
		},
	}, nil
}

// GetPoints TODO.
func (m *ArticleModel) GetPoints(string) (int, error) {
	return 0, nil
}
