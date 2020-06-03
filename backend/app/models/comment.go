package models

import (
	"database/sql"
	"time"
)

// CommentModel defines the database which the functions operate on.
type CommentModel struct {
	DB *sql.DB
}

// CommentInfo defines the public info of a comment.
type CommentInfo struct {
	Subreddit  string    `json:"subreddit"`
	ArticleID  string    `json:"articleID"`
	CommentID  string    `json:"commentID"`
	Body       string    `json:"body"`
	Points     int       `json:"points"`
	PostedBy   string    `json:"postedBy"`
	PostedTime time.Time `json:"postedTime"`
}

// Insert adds a new comment to the database, and returns its cid if successful.
func (m *CommentModel) Insert(subName, articleID, body, postedBy string) (string, error) {
	// TODO
	return "", nil
}

// ModifyBody modifies the body of a comment if the user is authorized to do so.
func (m *CommentModel) ModifyBody(commentID, username, newBody string) error {
	// TODO
	return nil
}

// Delete deletes a comment if the user is authorized to do so.
func (m *CommentModel) Delete(commentID, username string) error {
	// TODO
	return nil
}

// GetByArticle returns a list of comments in the article continued after a comment,
// and limited by a number.
func (m *CommentModel) GetByArticle(articleID, afterCommentID string, limit int) ([]CommentInfo, error) {
	// TODO
	return []CommentInfo{}, nil
}

// GetByUsername returns a list of comments posted by the user continued after a comment,
// and limited by a number.
func (m *CommentModel) GetByUsername(username, afterCommentID string, limit int) ([]CommentInfo, error) {
	// TODO
	return []CommentInfo{}, nil
}

// GetPoints returns the total points of a comment.
func (m *CommentModel) GetPoints(commentID string) (int, error) {
	// TODO
	return 0, nil
}
