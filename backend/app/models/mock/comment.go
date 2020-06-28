package mock

import (
	"fmt"
	"time"

	"github.com/YuChaoGithub/YARC/backend/app/models"
)

// CommentModel is the mock model for comment.
type CommentModel struct{}

// Insert returns ErrArticleNotExist if articleID is "88888".
func (m *CommentModel) Insert(articleID, body, postedBy string) (string, error) {
	if articleID == "88888" {
		return "", models.ErrArticleNotExist
	}

	return "Maxwell", nil
}

// ModifyBody returns ErrCommentNotExist if commentID is "Willow".
func (m *CommentModel) ModifyBody(commentID, username, newBody string) error {
	if commentID == "Willow" {
		return models.ErrCommentNotExist
	}

	return nil
}

// Delete returns ErrCommentNotExist if commentID is "Willow".
func (m *CommentModel) Delete(commentID, username string) error {
	if commentID == "Willow" {
		return models.ErrCommentNotExist
	}

	return nil
}

// Vote returns nil only if username=Jonny, commentID=88888, point=-1.
func (m *CommentModel) Vote(username, commentID string, point int) error {
	if username == "Jonny" && commentID == "88888" && point == -1 {
		return nil
	}

	return fmt.Errorf("")
}

// GetVote returns -1 if username=Jonny, commentID=88888.
func (m *CommentModel) GetVote(username, commentID string) int {
	if username == "Jonny" && commentID == "88888" {
		return -1
	}
	return 0
}

// Get has only the comment with ID 12345. It returns ErrCommentNotExist otherwise.
func (m *CommentModel) Get(commentID string) (models.CommentInfo, error) {
	if commentID == "12345" {
		return models.CommentInfo{
			Subreddit:  "radiohead",
			ArticleID:  "WX-78",
			CommentID:  "12345",
			Body:       "Flan",
			Points:     3,
			PostedBy:   "Jonny",
			PostedTime: time.Date(2003, 11, 24, 0, 0, 0, 0, time.UTC),
		}, nil
	}

	return models.CommentInfo{}, models.ErrCommentNotExist
}

// GetByArticle will return a result only if articleID=124215545, afterCommentID=451542212, limit=3.
func (m *CommentModel) GetByArticle(articleID, afterCommentID string, limit int) ([]models.CommentInfo, error) {
	if articleID == "124215545" && afterCommentID == "451542212" && limit == 3 {
		return []models.CommentInfo{
			{
				Subreddit:  "blur",
				ArticleID:  "mAGiC",
				CommentID:  "wHiPp",
				Body:       "I was only 21 when I watched it on TV.",
				Points:     689,
				PostedBy:   "GorillaZ",
				PostedTime: time.Date(1999, 9, 9, 0, 0, 0, 0, time.UTC),
			},
			{
				Subreddit:  "thesmiths",
				ArticleID:  "tENtoN",
				CommentID:  "tRUcC",
				Body:       "The pain was enough to make a shy bald buddist reflect.",
				Points:     333,
				PostedBy:   "BernardSumner",
				PostedTime: time.Date(2012, 4, 1, 0, 0, 0, 0, time.UTC),
			},
		}, nil
	}

	return []models.CommentInfo{}, nil
}

// GetByUsername will only return a result if username=god, afterCommentID="", limit=25.
func (m *CommentModel) GetByUsername(username, afterCommentID string, limit int) ([]models.CommentInfo, error) {
	if username == "god" && afterCommentID == "" && limit == 25 {
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
		}, nil
	}

	return []models.CommentInfo{}, nil
}
