package models

import (
	"reflect"
	"testing"
	"time"
)

func genStrWithLen(n int) string {
	longBody := []byte{}
	for i := 0; i < 600; i++ {
		longBody = append(longBody, 'a')
	}
	return string(longBody)
}

func TestCommentInsert(t *testing.T) {
	// Testcases.
	tests := []struct {
		name      string
		articleID string
		body      string
		postedBy  string
		wantError error
	}{
		{"Valid", "IpX177", "An epic post, lol!", "Jonny", nil},
		{"Invalid body", "246o19", genStrWithLen(600), "Morrissey", ErrCommentBodyInvalid},
		{"Invalid articleID", "benQ", "Good stuff.", "Jonny", ErrArticleNotExist},
		{"Invalid Username", "IpX1779", "Cruel.", "Kurt", ErrUsernameNotExist},
	}

	// Perform tests.
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Stub and driver.
			db, teardown := newTestDB(t)
			defer teardown()

			m := CommentModel{db}

			// When.
			_, err := m.Insert(tc.articleID, tc.body, tc.postedBy)

			// Want.
			if err != tc.wantError {
				t.Errorf("want %v; got %v", tc.wantError, err)
			}
		})
	}
}

func TestCommentModifyBody(t *testing.T) {
	// Testcases.
	tests := []struct {
		name      string
		commentID string
		username  string
		newBody   string
		wantError error
	}{
		{"Valid", "fewji3", "Albarn", "8 o'clock Kowloon emptiness.", nil},
		{"Invalid new body", "fewji3", "Albarn", genStrWithLen(689), ErrCommentBodyInvalid},
		{"Invalid commentID", "harder", "Jonny", "women and children first", ErrCommentNotExist},
		{"Invalid username", "01jf9", "Morrissey", "Heaven knows I'm mississippi now", ErrCommentNotExist},
	}

	// Perform tests.
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Stub and driver.
			db, teardown := newTestDB(t)
			defer teardown()

			m := CommentModel{db}

			// When.
			err := m.ModifyBody(tc.commentID, tc.username, tc.newBody)

			// Want.
			if err != tc.wantError {
				t.Errorf("want %v; got %v", tc.wantError, err)
			}
		})
	}
}

func TestCommentDelete(t *testing.T) {
	// Testcases.
	tests := []struct {
		name      string
		commentID string
		username  string
		wantError error
	}{
		{"Valid", "wjifo", "Morrissey", nil},
		{"Invalid Username", "wjifo", "Jonny", ErrCommentNotExist},
		{"Invalid CommentID", "77777", "KuoYu", ErrCommentNotExist},
	}

	// Perform tests.
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Stub and driver.
			db, teardown := newTestDB(t)
			defer teardown()

			m := CommentModel{db}

			// When.
			err := m.Delete(tc.commentID, tc.username)

			// Want.
			if err != tc.wantError {
				t.Errorf("want %v; got %v", tc.wantError, err)
			}
		})
	}
}

func TestCommentVote(t *testing.T) {
	// Testcases.
	tests := []struct {
		name      string
		username  string
		commentID string
		point     int
		wantError error
	}{
		{"Valid", "Morrissey", "007op", 1, nil},
		{"Change Vote", "Jonny", "WhipP", 0, nil},
		{"Invalid Comment", "Jonny", "77777", 0, ErrCommentNotExist},
		{"Invalid Username", "Kurt", "007op", 1, ErrUsernameNotExist},
	}

	// Perform Tests.
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Stub and driver.
			db, teardown := newTestDB(t)
			defer teardown()

			m := CommentModel{db}

			// When.
			err := m.Vote(tc.username, tc.commentID, tc.point)

			// Want.
			if err != tc.wantError {
				t.Errorf("want %v; got %v", tc.wantError, err)
			}
		})
	}
}

func TestCommentGet(t *testing.T) {
	// Testcases.
	tests := []struct {
		name        string
		commentID   string
		wantComment CommentInfo
		wantError   error
	}{
		{"Not Exist", "77777", CommentInfo{}, ErrCommentNotExist},
		{"Negative Points", "WhipP", CommentInfo{
			Subreddit:  "PHP",
			ArticleID:  "IpX177",
			CommentID:  "WhipP",
			Body:       "That is not true",
			Points:     -3,
			PostedBy:   "Albarn",
			PostedTime: time.Date(2020, 1, 20, 3, 2, 20, 0, time.UTC),
		}, nil},
		{"Zero Points", "fewji3", CommentInfo{
			Subreddit:  "dankmeme",
			ArticleID:  "246o1",
			CommentID:  "fewji3",
			Body:       "Nice.",
			Points:     0,
			PostedBy:   "Albarn",
			PostedTime: time.Date(2020, 1, 1, 1, 1, 1, 0, time.UTC),
		}, nil},
	}

	// Perform tests.
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Stub and driver.
			db, teardown := newTestDB(t)
			defer teardown()

			m := CommentModel{db}

			// When.
			comment, err := m.Get(tc.commentID)

			// Want.
			if err != tc.wantError {
				t.Errorf("want %v; got %v", tc.wantError, err)
			}

			if !reflect.DeepEqual(comment, tc.wantComment) {
				t.Errorf("want:\n%v\ngot:\n%v", tc.wantComment, comment)
			}
		})
	}
}

func TestCommentGetByArticle(t *testing.T) {
	// Testcases.
	tests := []struct {
		name           string
		articleID      string
		after          string
		limit          int
		wantCommentIDs []string
		wantError      error
	}{
		{"Invalid limit", "246o1", "t09o3", 101, []string{}, ErrLimitInvalid},
		{"From start", "246o1", "", 3, []string{"fewji3", "01jf9", "wjifo"}, nil},
		{"In between", "IpX177", "WhipPz", 5, []string{"WhipPx", "WhipP9", "n01bs", "n00bs", "WhipP"}, nil},
		{"At the end", "t09o3", "", 5, []string{"007op"}, nil},
	}

	// Perform tests.
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Stub and driver.
			db, teardown := newTestDB(t)
			defer teardown()

			m := CommentModel{db}

			// When.
			res, err := m.GetByArticle(tc.articleID, tc.after, tc.limit)

			// Want.
			if err != tc.wantError {
				t.Errorf("want %v; got %v", tc.wantError, err)
			}

			if len(res) != len(tc.wantCommentIDs) {
				t.Fatalf("want len %v; got len %v", len(tc.wantCommentIDs), len(res))
			}
			for i := range res {
				if res[i].CommentID != tc.wantCommentIDs[i] {
					t.Errorf("want %v; got %v", tc.wantCommentIDs[i], res[i].CommentID)
				}
			}
		})
	}
}

func TestCommentGetByUsername(t *testing.T) {
	// Testcases.
	tests := []struct {
		name           string
		username       string
		after          string
		limit          int
		wantCommentIDs []string
		wantError      error
	}{
		{"Invalid limit", "Jonny", "t09o3", 101, []string{}, ErrLimitInvalid},
		{"From start", "Jonny", "", 3, []string{"007op", "01jf9", "n01bs"}, nil},
		{"In between", "Albarn", "WhipPz", 4, []string{"WhipPx", "WhipP9", "fewji3", "WhipP"}, nil},
		{"At the end", "Morrissey", "wjifo", 5, []string{}, nil},
	}

	// Perform tests.
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Stub and driver.
			db, teardown := newTestDB(t)
			defer teardown()

			m := CommentModel{db}

			// When.
			res, err := m.GetByUsername(tc.username, tc.after, tc.limit)

			// Want.
			if err != tc.wantError {
				t.Errorf("want %v; got %v", tc.wantError, err)
			}

			if len(res) != len(tc.wantCommentIDs) {
				t.Fatalf("want len %v; got len %v", len(tc.wantCommentIDs), len(res))
			}
			for i := range res {
				if res[i].CommentID != tc.wantCommentIDs[i] {
					t.Errorf("want %v; got %v", tc.wantCommentIDs[i], res[i].CommentID)
				}
			}
		})
	}
}

func TestCommentVoteGet(t *testing.T) {
	// Testcases.
	tests := []struct {
		name      string
		username  string
		commentID string
		point     int
		wantPoint int
	}{
		{"New Vote", "Morrissey", "007op", -1, 0},
		{"Change Vote", "Jonny", "WhipP", 1, -1},
	}

	// Perform tests.
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Stub and driver.
			db, teardown := newTestDB(t)
			defer teardown()

			m := CommentModel{db}

			// When.
			err := m.Vote(tc.username, tc.commentID, tc.point)
			if err != nil {
				t.Fatal(err)
			}

			comment, err := m.Get(tc.commentID)
			if err != nil {
				t.Fatal(err)
			}

			// Want.
			if comment.Points != tc.wantPoint {
				t.Errorf("want %v; got %v", tc.wantPoint, comment.Points)
			}
		})
	}
}
