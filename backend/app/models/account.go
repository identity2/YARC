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
	bioMaxLen      = 256
)

// UserInfo stores the public information of an account.
type UserInfo struct {
	Username string    `json:"username"`
	Karma    int       `json:"karma"`
	Bio      string    `json:"bio"`
	JoinTime time.Time `json:"joinTime"`
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
var ErrEmailDup = errors.New("the email is already used")

// ErrAccountInfoMismatch means the username and password provided doesn't match.
var ErrAccountInfoMismatch = errors.New("the username and password does not match")

// ErrUsernameNotExist means the username doesn't exist in the database.
var ErrUsernameNotExist = errors.New("the username does not exist")

// ErrBioInvalid means the bio of the user is not valid.
var ErrBioInvalid = errors.New("the short bio should contain at most 256 characters")

// ErrArticleAlreadySaved means the article is already saved by the user.
var ErrArticleAlreadySaved = errors.New("the article is already saved")

// ErrArticleNotSaved means the article is not saved by the user.
var ErrArticleNotSaved = errors.New("the article is not saved by the user")

// ErrSubAlreadyJoined means the user already joined the subreddit.
var ErrSubAlreadyJoined = errors.New("the user already joined the subreddit")

// ErrSubNotJoined means the user has not joined the subreddit.
var ErrSubNotJoined = errors.New("the user has not joined the subreddit")

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
	// Validate the bio.
	if len(newBio) > bioMaxLen {
		return ErrBioInvalid
	}

	// Update the database.
	stmt := `UPDATE account SET bio = $1 WHERE username = $2`
	res, err := m.DB.Exec(stmt, newBio, username)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil || rowsAffected == 0 {
		return ErrUsernameNotExist
	}

	// Successfully updated the bio.
	return nil
}

// SaveArticle saves an article for the user.
func (m *AccountModel) SaveArticle(articleID, username string) error {
	stmt := `INSERT INTO save_article (username, sub_name, aid) VALUES ($1, (SELECT sub_name FROM article WHERE aid = $2), $2)`
	_, err := m.DB.Exec(stmt, username, articleID)
	if err, ok := err.(*pq.Error); ok {
		if strings.Contains(err.Message, "unique") {
			return ErrArticleAlreadySaved
		} else if strings.Contains(err.Message, "sub_name") || strings.Contains(err.Message, "aid") {
			return ErrArticleNotExist
		} else if strings.Contains(err.Message, "username") {
			return ErrUsernameNotExist
		}
		return err
	}

	// Successfully inserted to the database.
	return nil
}

// UnsaveArticle unsaves an article for the user.
func (m *AccountModel) UnsaveArticle(articleID, username string) error {
	stmt := `DELETE FROM save_article WHERE aid = $1 AND username = $2`
	res, err := m.DB.Exec(stmt, articleID, username)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil || rowsAffected == 0 {
		return ErrArticleNotSaved
	}

	// Successfully deleted from the database.
	return nil
}

// JoinSubreddit lets the user join a subreddit.
func (m *AccountModel) JoinSubreddit(subName, username string) error {
	stmt := `INSERT INTO join_sub (username, sub_name) VALUES ($1, $2)`
	_, err := m.DB.Exec(stmt, username, subName)
	if err, ok := err.(*pq.Error); ok {
		if strings.Contains(err.Message, "unique") {
			return ErrSubAlreadyJoined
		} else if strings.Contains(err.Message, "username") {
			return ErrUsernameNotExist
		} else if strings.Contains(err.Message, "sub_name") {
			return ErrSubredditNotExist
		}
		return err
	}

	// Successfully inserted to the database.
	return nil
}

// LeaveSubreddit lets the user leave a subscribed subreddit.
func (m *AccountModel) LeaveSubreddit(subName, username string) error {
	stmt := `DELETE FROM join_sub WHERE username = $1 AND sub_name = $2`
	res, err := m.DB.Exec(stmt, username, subName)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil || rowsAffected == 0 {
		return ErrSubNotJoined
	}

	// Successfully deleted from the database.
	return nil
}

// Get selects an account in the database and returns its info.
func (m *AccountModel) Get(username string) (UserInfo, error) {
	u := UserInfo{}
	stmt := `SELECT
	username, (COALESCE(ART_PT.karma, 0) + COALESCE(COM_PT.karma, 0)) AS karma, bio, join_time
	FROM account ACC
	LEFT JOIN (
		SELECT posted_by, SUM(point) AS karma FROM vote_article VA
		INNER JOIN article ART ON VA.aid = ART.aid GROUP BY ART.posted_by
	) ART_PT
	ON ACC.username = ART_PT.posted_by
	LEFT JOIN (
		SELECT posted_by, SUM(point) AS karma FROM vote_comment VC
		INNER JOIN comment COM ON VC.cid = COM.cid GROUP BY COM.posted_by
	) COM_PT
	ON ACC.username = COM_PT.posted_by
	WHERE username = $1`
	row := m.DB.QueryRow(stmt, username)

	err := row.Scan(&u.Username, &u.Karma, &u.Bio, &u.JoinTime)
	if err != nil {
		return u, ErrUsernameNotExist
	}

	// Successfully selected the user.
	return u, nil
}
