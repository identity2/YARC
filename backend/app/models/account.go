package models

import (
	"database/sql"
	"errors"
	"regexp"
	"strings"
	"time"

	"github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

const (
	usernameMinLen = 3
	usernameMaxLen = 20
	passwordMinLen = 6
	passwordMaxLen = 20
)

// UserInfo stores the public information of an account.
type UserInfo struct {
	Username string `json:"username"`
	Karma    int    `json:"karma"`
	Bio      string `json:"bio"`
}

var errUsernameInvalid = errors.New("the username should be 3-20 alphanumerical or underscore characters")
var errUsernameDup = errors.New("the username is used")
var errPasswordInvalid = errors.New("the password should be 6-20 alphanumerical or underscore characters")
var errEmailInvalid = errors.New("the email has invalid format")
var errEmailDup = errors.New("the email is used")
var errAccountInfoMismatch = errors.New("the username and password does not match")
var errUsernameNotExist = errors.New("the username does not exist")

var usernamePasswordRegExp = regexp.MustCompile("^[a-zA-Z0-9_]*$")
var emailRegExp = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

// AccountModel defines the database which the functions operate on.
type AccountModel struct {
	DB *sql.DB
}

// Insert adds a new account to the database.
func (m *AccountModel) Insert(username, email, password string) error {
	// Validate the username.
	usernameLen := len(username)
	if usernameLen < usernameMinLen || usernameLen > usernameMaxLen || !usernamePasswordRegExp.MatchString(username) {
		return errUsernameInvalid
	}

	// Validate the email.
	if !emailRegExp.MatchString(email) {
		return errEmailInvalid
	}

	// Validate the password.
	passwordLen := len(password)
	if passwordLen < passwordMinLen || passwordLen > passwordMaxLen || !usernamePasswordRegExp.MatchString(password) {
		return errPasswordInvalid
	}

	// Hash the password.
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return err
	}

	// Insert into the database.
	stmt := `INSERT INTO account (username, hashed_password, email, join_time) VALUES ($1, $2, $3, $4)`
	_, err = m.DB.Exec(stmt, username, hashedPassword, email, time.Now().UTC())
	if err, ok := err.(*pq.Error); ok {
		if strings.Contains(err.Message, "username_unique") {
			return errUsernameDup
		} else if strings.Contains(err.Message, "email_unique") {
			return errEmailDup
		}
		return err
	}

	// Successfully insert to the database.
	return nil
}

// Authenticate authenticates a user by checking if the username and password
// matches.
func (m *AccountModel) Authenticate(username, password string) error {
	var hashedPassword []byte

	// Get the user account from the database.
	stmt := `SELECT hashed_password FROM account WHERE username = $1`
	row := m.DB.QueryRow(stmt, username)

	err := row.Scan(&hashedPassword)
	if err != nil {
		return errUsernameNotExist
	}

	// Validate the password.
	err = bcrypt.CompareHashAndPassword(hashedPassword, []byte(password))
	if err != nil {
		return errAccountInfoMismatch
	}

	// TODO: Issue a JWT token. (Maybe do this in another function).

	return nil
}

// Get selects an account in the database and returns its info.
func (m *AccountModel) Get(username string) (UserInfo, error) {
	// TODO

	return UserInfo{}, nil
}
