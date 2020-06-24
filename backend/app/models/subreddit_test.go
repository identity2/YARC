package models

import (
	"reflect"
	"testing"
)

func TestSubredditInsert(t *testing.T) {
	// Testcases.
	tests := []struct {
		name        string
		subName     string
		description string
		wantError   error
	}{
		{"Valid", "news", "News about the world (aka USA).", nil},
		{"Subreddit already exists", "dankmeme", "The dankest meme.", ErrSubNameDup},
		{"Empty subreddit name", "", "ok", ErrSubNameInvalid},
	}

	// Perform tests.
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Stub and driver.
			db, teardown := newTestDB(t)
			defer teardown()

			m := SubredditModel{db, nil}

			// When.
			err := m.Insert(tc.subName, tc.description)

			// Want.
			if err != tc.wantError {
				t.Errorf("want %v; got %v", tc.wantError, err)
			}
		})
	}
}

// Combination test.
func TestSubredditGet(t *testing.T) {
	// Testcases.
	tests := []struct {
		name          string
		subName       string
		wantSubreddit SubredditInfo
		wantError     error
	}{
		{"Valid", "golang", SubredditInfo{"golang", "Hey, ho! Let's go!"}, nil},
		{"Not Exist", "subsifellfor", SubredditInfo{}, ErrSubredditNotExist},
	}

	// Perform tests.
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Stub and driver.
			db, teardown := newTestDB(t)
			defer teardown()

			m := SubredditModel{db, nil}

			// When.
			subreddit, err := m.Get(tc.subName)

			// Want.
			if !reflect.DeepEqual(subreddit, tc.wantSubreddit) {
				t.Errorf("want:\n%v\ngot:\n%v", tc.wantSubreddit, subreddit)
			}

			if err != tc.wantError {
				t.Errorf("want %v; got %v", tc.wantError, err)
			}
		})
	}
}

func TestInsertAndGet(t *testing.T) {
	wantSubreddit := SubredditInfo{"news", "Which do you want to hear first?"}

	// Stub and driver.
	db, teardown := newTestDB(t)
	defer teardown()

	m := SubredditModel{db, nil}

	// When.
	err := m.Insert(wantSubreddit.Name, wantSubreddit.Description)
	if err != nil {
		t.Fatal(err)
	}

	subreddit, err := m.Get(wantSubreddit.Name)

	// Want.
	if !reflect.DeepEqual(subreddit, wantSubreddit) {
		t.Errorf("want:\n%v\ngot:\n%v", wantSubreddit, subreddit)
	}
}

func TestSubredditGetTrending(t *testing.T) {
	// Stub and driver.
	db, teardown := newTestDB(t)
	defer teardown()

	rdb, rTeardown := newTestRedis(t)
	defer rTeardown()

	m := SubredditModel{db, rdb}

	// When.
	subreddits, err := m.GetTrending(3)
	if err != nil {
		t.Fatal(err)
	}

	// Want.
	wantSubreddits := []SubredditInfo{
		{Name: "dankmeme", Description: "The dankest meme of the world!"},
		{Name: "golang", Description: "Hey, ho! Let's go!"},
		{Name: "PHP", Description: "The best programming language in the world!"},
	}

	if !reflect.DeepEqual(subreddits, wantSubreddits) {
		t.Errorf("want:\n%v\ngot:\n%v", subreddits, wantSubreddits)
	}
}
