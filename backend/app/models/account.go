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

// ErrUsernameInvalid means the username contains illegal characters or is too long/short.
var ErrUsernameInvalid = errors.New("the username should be 3-20 alphanumerical or underscore characters")

// ErrUsernameDup means the username already exists in the database.
var ErrUsernameDup = errors.New("the username is used")

// ErrPasswordInvalid means the password contains illegal characters or is too long/short.
var ErrPasswordInvalid = errors.New("the password should be 6-20 alphanumerical or underscore characters")

// ErrEmailInvalid means the format of the email is invalid.
var ErrEmailInvalid = errors.New("the email has invalid format")

// ErrEmailDup means the email already exists in the database.
var ErrEmailDup = errors.New("the email is used")

// ErrAccountInfoMismatch means the username and password provided doesn't match.
var ErrAccountInfoMismatch = errors.New("the username and password does not match")

// ErrUsernameNotExist means the username doesn't exist in the database.
var ErrUsernameNotExist = errors.New("the username does not exist")

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
		return ErrUsernameInvalid
	}

	// Validate the email.
	if !emailRegExp.MatchString(email) {
		return ErrEmailInvalid
	}

	// Validate the password.
	passwordLen := len(password)
	if passwordLen < passwordMinLen || passwordLen > passwordMaxLen || !usernamePasswordRegExp.MatchString(password) {
		return ErrPasswordInvalid
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
			return ErrUsernameDup
		} else if strings.Contains(err.Message, "email_unique") {
			return ErrEmailDup
		}
		return err
	}

	// Successfully inserted to the database.
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
		return ErrUsernameNotExist
	}

	// Validate the password.
	err = bcrypt.CompareHashAndPassword(hashedPassword, []byte(password))
	if err != nil {
		return ErrAccountInfoMismatch
	}

	return nil
}

// ModifyBio modifies bio of the user.
func (m *AccountModel) ModifyBio(username, newBio string) error {
	// TODO
	return nil
}

// SaveArticle saves an article for the user.
func (m *AccountModel) SaveArticle(articleID, username string) error {
	// TODO
	return nil
}

// JoinSubreddit lets the user join a subreddit.
func (m *AccountModel) JoinSubreddit(subName, username string) error {
	// TODO
	return nil
}

// Get selects an account in the database and returns its info.
func (m *AccountModel) Get(username string) (UserInfo, error) {
	// TODO

	return UserInfo{}, nil
}
