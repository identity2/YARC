package models

import "database/sql"

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
	// TODO
	return SearchResults{}, nil
}
