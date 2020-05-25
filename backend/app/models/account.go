package models

import "database/sql"

// AccountModel defines the database which the functions operate on.
type AccountModel struct {
	DB *sql.DB
}

// Insert adds a new account to the database.
func (m *AccountModel) Insert(username, email, password string) error {
	// TODO
	return nil
}

// Authenticate authenticates a user by checking if the username and password
// matches.
func (m *AccountModel) Authenticate(username, password string) error {
	// TODO
	return nil
}

// Get selects a user by the username from the database and returns it.
func (m *AccountModel) Get(username string) {
	// TODO
}
