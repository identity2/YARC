package mock

import (
	"time"

	"github.com/YuChaoGithub/YARC/backend/app/models"
)

// ArticleModel is the mock model for article.
type ArticleModel struct{}

// Insert returns ErrSubredditNotExist if the subreddit is "dankmeme",
// otherwise it returns "WX-78", nil.
func (m *ArticleModel) Insert(subName, postType, body, title, postedBy string) (string, error) {
	if subName == "dankmeme" {
		return "", models.ErrSubredditNotExist
	}

	return "WX-78", nil
}

// ModifyBody returns ErrArticleNotExist if articleID is "Dijkstra",
// other it returns nil.
func (m *ArticleModel) ModifyBody(articleID, username, body string) error {
	if articleID == "Dijkstra" {
		return models.ErrArticleNotExist
	}

	return nil
}

// Delete returns ErrArticleNotExist if articleID is "Dijkstra",
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
			Comments:   30,
			PostedBy:   "Bellamy",
			PostedTime: time.Date(2019, 11, 13, 0, 0, 0, 0, time.UTC),
		}, nil
	}

	return models.ArticleInfo{}, models.ErrArticleNotExist
}

// GetAll always return a valid list of article.
func (m *ArticleModel) GetAll(afterArticleID, sortedBy string, limit int) ([]models.ArticleInfo, error) {
	return []models.ArticleInfo{
		{
			Subreddit:  "radiohead",
			ArticleID:  "WX-78",
			Type:       "text",
			Body:       "1",
			Title:      "2",
			Points:     1,
			Comments:   2,
			PostedBy:   "Tom",
			PostedTime: time.Date(2000, 12, 12, 12, 12, 12, 0, time.UTC),
		},
		{
			Subreddit:  "meirl",
			ArticleID:  "tyty",
			Type:       "text",
			Body:       "2",
			Title:      "11",
			Points:     -3,
			Comments:   1,
			PostedBy:   "Tim",
			PostedTime: time.Date(2020, 1, 1, 1, 1, 1, 0, time.UTC),
		},
	}, nil
}

// GetBySubscribed will not be called for now.
func (m *ArticleModel) GetBySubscribed(username, afterArticleID, sortedBy string, limit int) ([]models.ArticleInfo, error) {
	return []models.ArticleInfo{}, nil
}

// GetBySubreddit will only return a result if subName=radiohead, sortedBy=hot, limit=1, afterArticleID=ok.
func (m *ArticleModel) GetBySubreddit(subName, afterArticleID, sortedBy string, limit int) ([]models.ArticleInfo, error) {
	if afterArticleID == "ok" && subName == "radiohead" && sortedBy == "hot" && limit == 1 {
		return []models.ArticleInfo{
			{
				Subreddit:  "radiohead",
				ArticleID:  "WX-78",
				Type:       "text",
				Body:       "I'm a reasonable man, get off my case.",
				Title:      "Tin can",
				Points:     42,
				Comments:   3,
				PostedBy:   "Bellamy",
				PostedTime: time.Date(2019, 11, 13, 0, 0, 0, 0, time.UTC),
			},
		}, nil
	}

	return []models.ArticleInfo{}, nil
}

// GetByUser will only return a result if username=Jonny, sortedBy=new, limit=25, afterArticleID=bad.
func (m *ArticleModel) GetByUser(username, afterArticleID, sortedBy string, limit int) ([]models.ArticleInfo, error) {
	if afterArticleID == "bad" && username == "Jonny" && sortedBy == "new" && limit == 25 {
		return []models.ArticleInfo{
			{
				Subreddit:  "a",
				ArticleID:  "a",
				Type:       "text",
				Body:       "a",
				Title:      "a",
				Points:     -12,
				Comments:   0,
				PostedBy:   "Freeman",
				PostedTime: time.Date(2019, 11, 13, 0, 0, 0, 0, time.UTC),
			},
		}, nil
	}

	return []models.ArticleInfo{}, nil
}

// GetSavedByUser will only return a result if username=Albarn, sortedBy=old, limit=2, afterArticleID="".
func (m *ArticleModel) GetSavedByUser(username, afterArticleID, sortedBy string, limit int) ([]models.ArticleInfo, error) {
	if afterArticleID == "" && username == "Albarn" && sortedBy == "old" && limit == 2 {
		return []models.ArticleInfo{
			{
				Subreddit:  "b",
				ArticleID:  "b",
				Type:       "text",
				Body:       "b",
				Title:      "b",
				Points:     420,
				Comments:   69,
				PostedBy:   "mememan",
				PostedTime: time.Date(2019, 11, 13, 0, 0, 0, 0, time.UTC),
			},
		}, nil
	}

	return []models.ArticleInfo{}, nil
}
