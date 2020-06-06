package mock

import (
	"time"

	"github.com/YuChaoGithub/YARC/backend/app/models"
)

// ArticleModel is the mock model for article.
type ArticleModel struct{}

// Insert returns errSubredditNotExist if the subreddit is "dankmeme",
// otherwise it returns "WX-78", nil.
func (m *ArticleModel) Insert(subName, postType, body, title, postedBy string) (string, error) {
	if subName == "dankmeme" {
		return "", models.ErrSubredditNotExist
	}

	return "WX-78", nil
}

// ModifyBody returns errArticleNotExist if articleID is "Dijkstra",
// other it returns nil.
func (m *ArticleModel) ModifyBody(articleID, username, body string) error {
	if articleID == "Dijkstra" {
		return models.ErrArticleNotExist
	}

	return nil
}

// Delete returns errArticleNotExist if articleID is "Dijkstra",
// otherwise it returns nil.
func (m *ArticleModel) Delete(articleID, username string) error {
	if articleID == "Dijkstra" {
		return models.ErrArticleNotExist
	}

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
func (m *ArticleModel) GetPoints(string) int {
	return 0
}
