package models

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/lib/pq"
	"github.com/teris-io/shortid"
)

const (
	commentMaxLen = 512
)

// ErrCommentBodyInvalid means the comment body is too long.
var ErrCommentBodyInvalid = errors.New("the comment should contain 512 characters at most")

// ErrCommentNotExist means the comment doesn't exist.
var ErrCommentNotExist = errors.New("the comment does not exist")

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
func (m *CommentModel) Insert(articleID, body, postedBy string) (string, error) {
	// Validate body.
	if len(body) > commentMaxLen {
		return "", ErrCommentBodyInvalid
	}

	// Generate a short ID for the comment.
	commentID, err := shortid.Generate()
	if err != nil {
		return "", fmt.Errorf("error generating ID for comment")
	}

	// Insert into the datbase.
	stmt := `INSERT INTO comment (sub_name, aid, cid, body, posted_by, posted_time) VALUES 
		((SELECT sub_name FROM article WHERE aid = $1), $1, $2, $3, $4, $5)`
	_, err = m.DB.Exec(stmt, articleID, commentID, body, postedBy, time.Now().UTC())
	if err, ok := err.(*pq.Error); ok {
		if strings.Contains(err.Message, "sub_name") || strings.Contains(err.Message, "aid") {
			return "", ErrArticleNotExist
		} else if strings.Contains(err.Message, "posted_by") {
			return "", ErrUsernameNotExist
		}
		return "", err
	}

	// Successfully inserted into the database.
	return commentID, nil
}

// ModifyBody modifies the body of a comment if the user is authorized to do so.
func (m *CommentModel) ModifyBody(commentID, username, newBody string) error {
	// Validate body.
	if len(newBody) > commentMaxLen {
		return ErrCommentBodyInvalid
	}

	// Update the database.
	stmt := `UPDATE comment SET body = $1 WHERE cid = $2 AND posted_by = $3`
	res, err := m.DB.Exec(stmt, newBody, commentID, username)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil || rowsAffected == 0 {
		return ErrCommentNotExist
	}

	// Successfully updated the comment.
	return nil
}

// Delete deletes a comment if the user is authorized to do so.
func (m *CommentModel) Delete(commentID, username string) error {
	stmt := `DELETE FROM comment WHERE cid = $1 AND posted_by = $2`
	res, err := m.DB.Exec(stmt, commentID, username)
	if err != nil {
		return nil
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil || rowsAffected == 0 {
		return ErrCommentNotExist
	}

	// Successfully delete the comment.
	return nil
}

// Vote inserts a vote entry for the commentID by the username.
func (m *CommentModel) Vote(username, commentID string, point int) error {
	stmt := `INSERT INTO vote_comment SELECT $1, sub_name, aid, cid, $3
		FROM comment WHERE cid = $2
		ON CONFLICT (username, sub_name, aid, cid) DO UPDATE SET point = $3`

	res, err := m.DB.Exec(stmt, username, commentID, point)
	if err, ok := err.(*pq.Error); ok {
		if strings.Contains(err.Message, "username") {
			return ErrUsernameNotExist
		}
		return err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil || rowsAffected == 0 {
		return ErrCommentNotExist
	}

	// Successfully inserted into the database.
	return nil
}

// GetVote gets the voting action of the currently user.
// If the query fails, it returns 0.
func (m *CommentModel) GetVote(username, commentID string) int {
	res := 0
	stmt := `SELECT point FROM vote_comment WHERE username = $1 AND cid = $2`
	row := m.DB.QueryRow(stmt, username, commentID)
	err := row.Scan(&res)
	if err != nil {
		return 0
	}
	return res
}

// Get returns a comment by the commentID.
func (m *CommentModel) Get(commentID string) (CommentInfo, error) {
	c := CommentInfo{}
	stmt := `SELECT C.sub_name, C.aid, C.cid, C.body, COALESCE(SUM(VC.point), 0) AS points, C.posted_by, C.posted_time
		FROM comment C LEFT JOIN vote_comment VC ON C.cid = VC.cid WHERE C.cid = $1 GROUP BY C.sub_name, C.aid, C.cid`

	row := m.DB.QueryRow(stmt, commentID)
	err := row.Scan(&c.Subreddit, &c.ArticleID, &c.CommentID, &c.Body, &c.Points, &c.PostedBy, &c.PostedTime)
	if err != nil {
		return c, ErrCommentNotExist
	}

	// Successfully selected the comment.
	return c, nil
}

// GetByArticle returns a list of comments in the article continued after a comment,
// and limited by a number.
func (m *CommentModel) GetByArticle(articleID, afterCommentID string, limit int) ([]CommentInfo, error) {
	return m.getList(`C.aid = $1`, articleID, afterCommentID, limit)
}

// GetByUsername returns a list of comments posted by the user continued after a comment,
// and limited by a number.
func (m *CommentModel) GetByUsername(username, afterCommentID string, limit int) ([]CommentInfo, error) {
	return m.getList(`C.posted_by = $1`, username, afterCommentID, limit)
}

// getList perform the select query to list comments according to the SQL clauses provided.
// The list is always sorted by hot (i.e. by points DESC, posted_time DESC).
func (m *CommentModel) getList(where, firstParam, afterCommentID string, limit int) ([]CommentInfo, error) {
	var res []CommentInfo

	// Validate limit.
	if limit < listLimitMin || limit > listLimitMax {
		return res, ErrLimitInvalid
	}

	// Select query statement.
	stmt := fmt.Sprintf(`WITH temp AS (
		SELECT C.sub_name, C.aid, C.cid, C.body, COALESCE(VC.points, 0) AS points, C.posted_by, C.posted_time
		FROM comment C LEFT JOIN (select cid, SUM(point) AS points FROM vote_comment GROUP BY sub_name, aid, cid) VC ON C.cid = VC.cid
		WHERE %s)
		SELECT sub_name, aid, cid, body, points, posted_by, posted_time FROM temp
		WHERE COALESCE(points < (SELECT points FROM temp WHERE cid = $2) OR
		(points = (SELECT points FROM temp WHERE cid = $2) AND
		posted_time < (SELECT posted_time FROM temp WHERE cid = $2)), TRUE)
		ORDER BY points DESC, posted_time DESC LIMIT $3`, where)

	rows, err := m.DB.Query(stmt, firstParam, afterCommentID, limit)
	if err != nil {
		return res, err
	}
	for rows.Next() {
		c := CommentInfo{}
		err = rows.Scan(&c.Subreddit, &c.ArticleID, &c.CommentID, &c.Body, &c.Points, &c.PostedBy, &c.PostedTime)
		if err != nil {
			return []CommentInfo{}, err
		}
		res = append(res, c)
	}

	return res, nil
}
