package mock

import (
	"time"

	"github.com/YuChaoGithub/YARC/backend/app/models"
)

// CommentModel is the mock model for comment.
type CommentModel struct{}

// Insert TODO.
func (m *CommentModel) Insert(subName, articleID, body, postedBy string) (string, error) {
	return "mAXwEll", nil
}

// ModifyBody TODO.
func (m *CommentModel) ModifyBody(commentID, username, newBody string) error {
	return nil
}

// Delete TODO.
func (m *CommentModel) Delete(commentID, username string) error {
	return nil
}

// GetByArticle TODO.
func (m *CommentModel) GetByArticle(articleID, afterCommentID string, limit int) ([]models.CommentInfo, error) {
	return []models.CommentInfo{
		{
			Subreddit:  "blur",
			ArticleID:  "mAGiC",
			CommentID:  "wHiPp",
			Body:       "I was only 21 when I watched it on TV.",
			Points:     689,
			PostedBy:   "GorillaZ",
			PostedTime: time.Date(1999, time.September, 9, 0, 0, 0, 0, time.UTC),
		},
		{
			Subreddit:  "thesmiths",
			ArticleID:  "tENtoN",
			CommentID:  "tRUcC",
			Body:       "The pain was enough to make a shy bald buddist reflect.",
			Points:     333,
			PostedBy:   "BernardSumner",
			PostedTime: time.Date(2012, time.April, 1, 0, 0, 0, 0, time.UTC),
		},
	}, nil
}

// GetByUsername TODO.
func (m *CommentModel) GetByUsername(username, afterCommentID string, limit int) ([]models.CommentInfo, error) {
	return []models.CommentInfo{
		{
			Subreddit:  "blur",
			ArticleID:  "mAGiC",
			CommentID:  "wHiPp",
			Body:       "I was only 21 when I watched it on TV.",
			Points:     689,
			PostedBy:   "GorillaZ",
			PostedTime: time.Date(1999, time.September, 9, 0, 0, 0, 0, time.UTC),
		},
		{
			Subreddit:  "thesmiths",
			ArticleID:  "tENtoN",
			CommentID:  "tRUcC",
			Body:       "The pain was enough to make a shy bald buddist reflect.",
			Points:     333,
			PostedBy:   "BernardSumner",
			PostedTime: time.Date(2012, time.April, 1, 0, 0, 0, 0, time.UTC),
		},
	}, nil
}
