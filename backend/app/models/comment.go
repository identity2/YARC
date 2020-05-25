package models

import "database/sql"

// CommentModel defines the database which the functions operate on.
type CommentModel struct {
	DB *sql.DB
}

// Insert adds a new comment to the database, and returns its cid if successful.
func (m *CommentModel) Insert(subName, aid, body, postedBy string) (string, error) {
	// TODO
	return "", nil
}

// GetAllByArticle selects all comments in a certain article, and returns them if successful.
func (m *CommentModel) GetAllByArticle(subName, aid string) {
	// TODO
}

// GetAllByUsername selects all comments posted by a certain user, and returns them if successful.
func (m *CommentModel) GetAllByUsername(username string) {
	// TODO
}
