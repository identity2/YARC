package models

import (
	"database/sql"
	"io/ioutil"
	"testing"

	"github.com/go-redis/redis"
)

const dbScriptPath = "../../../database/"

// newTestDB returns the mock database connection along with its teardown function.
func newTestDB(t *testing.T) (*sql.DB, func()) {
	// Database connection.
	db, err := sql.Open("postgres", "host=localhost port=5432 user=postgres password=password dbname=yarc_test sslmode=disable")
	if err != nil {
		t.Fatal(err)
	}

	// Database setup.
	script, err := ioutil.ReadFile(dbScriptPath + "setup.sql")
	if err != nil {
		t.Fatal(err)
	}

	_, err = db.Exec(string(script))
	if err != nil {
		t.Fatal(err)
	}

	// Insert mock data.
	script, err = ioutil.ReadFile(dbScriptPath + "mock_data.sql")
	if err != nil {
		t.Fatal(err)
	}

	_, err = db.Exec(string(script))
	if err != nil {
		t.Fatal(err)
	}

	// Return the db and the tear down function.
	return db, func() {
		script, err := ioutil.ReadFile(dbScriptPath + "teardown.sql")
		if err != nil {
			t.Fatal(err)
		}

		_, err = db.Exec(string(script))
		if err != nil {
			t.Fatal(err)
		}

		db.Close()
	}
}

// newTestRedis returns the mock Redis memory store along with its teardown function.
func newTestRedis(t *testing.T) (*redis.Client, func()) {
	// Redis connection.
	rdb := redis.NewClient(&redis.Options{Addr: "127.0.0.1:6379", DB: 7})

	// Data setup for subreddit.
	rdb.ZIncrBy("subreddit", 10, "dankmeme")
	rdb.ZIncrBy("subreddit", 7, "golang")
	rdb.ZIncrBy("subreddit", 3, "meirl")
	rdb.ZIncrBy("subreddit", 5, "PHP")

	// Return the db and the tear down function.
	return rdb, func() {
		rdb.FlushDB()
		rdb.Close()
	}
}
