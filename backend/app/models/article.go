package models

import (
	"database/sql"
	"errors"
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/lib/pq"
	"github.com/teris-io/shortid"
)

const (
	titleMaxLen    = 128
	textPostMaxLen = 1024
	linkPostMaxLen = 512
	defaultTitle   = "Untitled"
	textPost       = "text"
	linkPost       = "link"
	imagePost      = "image"
)

// ErrTitleInvalid means the title length is too long.
var ErrTitleInvalid = errors.New("the title should be at most 128 characters")

// ErrTextBodyInvalid means the text body of the article is too long.
var ErrTextBodyInvalid = errors.New("the body of the text post should be at most 512 characters")

// ErrLinkBodyInvalid means the link in the post is not a valid URL.
var ErrLinkBodyInvalid = errors.New("the body of the link post is invalid")

// ErrArticleNotExist means the article doesn't exist in the database.
var ErrArticleNotExist = errors.New("the article does not exist")

// ErrGetPoints means some error occurred when getting points from the database.
var ErrGetPoints = errors.New("error getting the points of the article")

// ErrInvalidPostType means the post type is invalid.
var ErrInvalidPostType = errors.New("the post type should be either 'text', 'image' or 'link'")

// ArticleModel defines the database which the functions operate on.
type ArticleModel struct {
	DB *sql.DB
}

// ArticleInfo is the public interface of an article.
type ArticleInfo struct {
	Subreddit  string    `json:"subreddit"`
	ArticleID  string    `json:"articleID"`
	Type       string    `json:"type"`
	Body       string    `json:"body"`
	Title      string    `json:"title"`
	Points     int       `json:"points"`
	PostedBy   string    `json:"postedBy"`
	PostedTime time.Time `json:"postedTime"`
}

// Insert adds a new article to the database, and returns its aid if successful.
func (m *ArticleModel) Insert(subName, postType, body, title, postedBy string) (string, error) {
	// Validate title.
	titleLen := len(title)
	if titleLen > titleMaxLen {
		return "", ErrTitleInvalid
	} else if titleLen == 0 {
		title = defaultTitle
	}

	// Validate body.
	if postType == textPost {
		// Text post.
		if len(body) > textPostMaxLen {
			return "", ErrTextBodyInvalid
		}
	} else if postType == linkPost || postType == imagePost {
		// Link or image post. (Image post is essentially a link to the image.)
		u, err := url.Parse(body)
		if err != nil || u.Scheme == "" || u.Host == "" || u.Path == "" {
			return "", ErrLinkBodyInvalid
		}
	} else {
		return "", ErrInvalidPostType
	}

	// Generate a short ID for the article.
	articleID, err := shortid.Generate()
	if err != nil {
		return "", fmt.Errorf("error generating ID for article")
	}

	// Insert into the database.
	stmt := `INSERT INTO article (sub_name, aid, type, body, title, posted_by, posted_time) VALUES ($1, $2, $3, $4, $5, $6, $7)`
	_, err = m.DB.Exec(stmt, subName, articleID, postType, body, title, postedBy, time.Now().UTC())
	if err, ok := err.(*pq.Error); ok {
		if strings.Contains(err.Message, "sub_name") {
			return "", ErrSubredditNotExist
		} else if strings.Contains(err.Message, "posted_by") {
			return "", ErrUsernameNotExist
		}
		return "", err
	}

	// Successfully inserted to the database.
	return articleID, nil
}

// ModifyBody modifies the body of an article if the user is authorized to do so,
// and the article is of text type.
func (m *ArticleModel) ModifyBody(articleID, username, body string) error {
	// Validate body.
	if len(body) > textPostMaxLen {
		return ErrTextBodyInvalid
	}

	// Update the database.
	stmt := `UPDATE article SET body = $1 WHERE aid = $2 AND posted_by = $3 AND type = 'text'`
	res, err := m.DB.Exec(stmt, body, articleID, username)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil || rowsAffected == 0 {
		return ErrArticleNotExist
	}

	// Successfully updated the article.
	return nil
}

// Delete deletes an article if the user is authorized to do so.
func (m *ArticleModel) Delete(articleID, username string) error {
	stmt := `DELETE FROM article WHERE aid = $1 AND posted_by = $2`
	res, err := m.DB.Exec(stmt, articleID, username)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil || rowsAffected == 0 {
		return ErrArticleNotExist
	}

	// Successfully deleted the article.
	return nil
}

// Get returns the article specified by the articleID.
func (m *ArticleModel) Get(articleID string) (ArticleInfo, error) {
	a := ArticleInfo{}
	stmt := `SELECT
	ART.sub_name, ART.aid, ART.type, ART.body, ART.title, COALESCE(SUM(VA.point), 0) AS points, ART.posted_by, ART.posted_time
	FROM article ART LEFT JOIN vote_article VA ON ART.aid = VA.aid
	WHERE ART.aid = $1
	GROUP BY ART.sub_name, ART.aid`
	row := m.DB.QueryRow(stmt, articleID)

	err := row.Scan(&a.Subreddit, &a.ArticleID, &a.Type, &a.Body, &a.Title, &a.Points, &a.PostedBy, &a.PostedTime)
	if err != nil {
		return a, ErrArticleNotExist
	}

	// Successfully selected the article.
	return a, nil
}

// GetBySubreddit returns a list of articles in the subreddit continued after an article,
// and limited by a number, sorted by hot/new/old.
func (m *ArticleModel) GetBySubreddit(subName, afterArticleID, sortedBy string, limit int) ([]ArticleInfo, error) {
	// TODO
	return []ArticleInfo{}, nil
}

// GetByUser returns a list of articles posted by the user continued after an article,
// and limited by a number, sorted by hot/new/old.
func (m *ArticleModel) GetByUser(username, afterArticleID, sortedBy string, limit int) ([]ArticleInfo, error) {
	// TODO
	return []ArticleInfo{}, nil
}

// GetSavedByUser returns a list of articles saved by the user continued after an article,
// and limited by a number, sorted by hot/new/old.
func (m *ArticleModel) GetSavedByUser(username, afterArticleID, sortedBy string, limit int) ([]ArticleInfo, error) {
	// TODO
	return []ArticleInfo{}, nil
}
