package models

import "database/sql"

const (
	searchSubredditResultLimit = 5
	searchArticleResultLimit   = 50
)

// SearchModel defines the database which the search functions operate on.
type SearchModel struct {
	DB *sql.DB
}

// SearchResults contain an array of SubredditInfo and an array of ArticleInfo.
type SearchResults struct {
	Subreddits []SubredditInfo `json:"subreddits"`
	Articles   []ArticleInfo   `json:"articles"`
}

// GetResult returns the search result of a given query string.
func (m *SearchModel) GetResult(query string) (SearchResults, error) {
	res := SearchResults{
		Subreddits: []SubredditInfo{},
		Articles:   []ArticleInfo{},
	}

	// Search for subreddit name and description.
	stmt := `SELECT sub_name, description FROM subreddit WHERE sub_name ILIKE '%' || $1 || '%' OR description ILIKE '%' ||  $1 || '%' ORDER BY sub_name ASC LIMIT $2`
	rows, err := m.DB.Query(stmt, query, searchSubredditResultLimit)
	if err != nil {
		return res, err
	}

	for rows.Next() {
		s := SubredditInfo{}
		err = rows.Scan(&s.Name, &s.Description)
		if err != nil {
			return res, err
		}
		res.Subreddits = append(res.Subreddits, s)
	}

	// Search for article title.
	stmt = `SELECT
	ART.sub_name, ART.aid, ART.type, ART.body, ART.title, COALESCE(VA.points, 0) AS points, COALESCE(COM.comments, 0) AS comments, ART.posted_by, ART.posted_time
	FROM article ART
	LEFT JOIN (SELECT aid, SUM(point) AS points FROM vote_article GROUP BY sub_name, aid) VA ON ART.aid = VA.aid
	LEFT JOIN (SELECT aid, COUNT(*) AS comments FROM comment GROUP BY aid) COM ON COM.aid = ART.aid
	WHERE ART.title ILIKE '%' || $1 || '%' ORDER BY points DESC, ART.posted_time DESC LIMIT $2`
	rows, err = m.DB.Query(stmt, query, searchArticleResultLimit)
	if err != nil {
		return res, err
	}

	for rows.Next() {
		a := ArticleInfo{}
		err = rows.Scan(&a.Subreddit, &a.ArticleID, &a.Type, &a.Body, &a.Title, &a.Points, &a.Comments, &a.PostedBy, &a.PostedTime)
		if err != nil {
			return res, err
		}
		res.Articles = append(res.Articles, a)
	}

	// Successfully getting the search results.
	return res, nil
}
