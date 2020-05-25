package models

import "database/sql"

// SubredditModel defines the database which the functions operate on.
type SubredditModel struct {
	DB *sql.DB
}
