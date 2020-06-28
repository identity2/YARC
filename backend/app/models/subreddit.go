package models

import (
	"database/sql"
	"errors"
	"regexp"
	"strings"

	"github.com/go-redis/redis"
	"github.com/lib/pq"
)

const (
	subNameMinLen     = 1
	subNameMaxLen     = 16
	descriptionMaxLen = 512
)

// SubredditModel defines the database which the functions operate on.
type SubredditModel struct {
	DB  *sql.DB
	RDB *redis.Client
}

// ErrSubNameInvalid means the subreddit name contains invalid characters or is too long/short.
var ErrSubNameInvalid = errors.New("the subreddit name should be 1-16 alphanumerical or underscore characters")

// ErrSubNameDup means the subreddit name already exists in the database.
var ErrSubNameDup = errors.New("the subreddit name is used")

// ErrDescriptionInvalid means the description of the subreddit is too long.
var ErrDescriptionInvalid = errors.New("the description should be at most 512 characters")

// ErrSubredditNotExist means the subreddit doesn't exist in the database.
var ErrSubredditNotExist = errors.New("the subreddit does not exist")

var subnameRegExp = regexp.MustCompile("^[a-zA-Z0-9_]*$")

// SubredditInfo is the public info of a subreddit.
type SubredditInfo struct {
	Name        string `json:"name"`
	Members     int    `json:"members"`
	Description string `json:"description"`
}

// List returns a list of all subreddit names.
func (m *SubredditModel) List() []string {
	var res []string

	stmt := `SELECT sub_name FROM subreddit ORDER BY sub_name ASC`
	rows, err := m.DB.Query(stmt)
	if err != nil {
		return res
	}

	for rows.Next() {
		var name string
		err = rows.Scan(&name)
		if err != nil {
			return res
		}
		res = append(res, name)
	}

	return res
}

// Insert adds a new subreddit to the database.
func (m *SubredditModel) Insert(name, description string) error {
	// Validate the subreddit name.
	nameLen := len(name)
	if nameLen < subNameMinLen || nameLen > subNameMaxLen || !subnameRegExp.MatchString(name) {
		return ErrSubNameInvalid
	}

	// Validate the description.
	if len(description) > descriptionMaxLen {
		return ErrDescriptionInvalid
	}

	// Insert into database.
	stmt := `INSERT INTO subreddit (sub_name, description) VALUES ($1, $2)`
	_, err := m.DB.Exec(stmt, name, description)
	if err, ok := err.(*pq.Error); ok {
		if strings.Contains(err.Message, "sub_name_unique") {
			return ErrSubNameDup
		}
		return err
	}

	// Successfully inserted to the database.
	return nil
}

// Get selects a subreddit in the database and returns its info.
func (m *SubredditModel) Get(name string) (SubredditInfo, error) {
	subInfo := SubredditInfo{}

	stmt := `SELECT S.sub_name, COALESCE(J.members, 0) AS members, S.description FROM subreddit S
	LEFT JOIN (SELECT sub_name, COUNT(*) AS members FROM join_sub GROUP BY sub_name) J ON S.sub_name = J.sub_name
	WHERE S.sub_name = $1`
	row := m.DB.QueryRow(stmt, name)

	err := row.Scan(&subInfo.Name, &subInfo.Members, &subInfo.Description)
	if err != nil {
		return subInfo, ErrSubredditNotExist
	}

	return subInfo, nil
}

// IncrVisitCount increments the visit count of the subreddit which affects the trending rankings.
func (m *SubredditModel) IncrVisitCount(name string) {
	// Does nothing if this goes wrong. (aka. no error check)
	m.RDB.ZIncrBy("subreddit", 1, name)
}

// GetTrending returns a list of trending subreddits limited by a number.
func (m *SubredditModel) GetTrending(limit int) ([]SubredditInfo, error) {
	res := []SubredditInfo{}

	// Get the top trending subreddits from Redis.
	subs, err := m.RDB.ZRevRange("subreddit", 0, int64(limit)-1).Result()
	if err != nil {
		return res, err
	}

	// Get the sub info from the persistent database.
	for _, val := range subs {
		s, err := m.Get(val)
		if err != nil {
			return []SubredditInfo{}, err
		}
		res = append(res, s)
	}

	return res, nil
}
