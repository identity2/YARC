package models

import (
	"reflect"
	"testing"
	"time"
)

func TestArticleInsert(t *testing.T) {
	// Testcases.
	tests := []struct {
		name      string
		subName   string
		postType  string
		body      string
		title     string
		postedBy  string
		wantError error
	}{
		{"Valid Text Article", "golang", "text", "Body of a valid article.", "Title of a valid article", "Jonny", nil},
		{"Valid Link Article", "golang", "link", "http://www.github.com/", "A nice Microsoft website", "Albarn", nil},
		{"Valid Image Article", "dankmeme", "image", "https://i.imgur.com/XXgVUqa.jpg", "Is good image", "Morrissey", nil},
		{"Invalid Title", "PHP", "link", "http://www.google.com/", "Come on, come on You think you drive me crazy Come on, come on You and whose army? You and your cronies Come on, come on Holy Roman Empire Come on if you think Come on if you think You can take us on You can take us on You and whose army? You and your cronies You forget so easily We ride tonight Ghost horses", "Morrissey", ErrTitleInvalid},
		{"Invalid Subreddit", "existentialism", "text", "Sick content right here.", "Warning, sick content:", "Jonny", ErrSubredditNotExist},
		{"Invalid Username", "PHP", "image", "https://i.imgur.com/XXgVUqa.jpg", "Below is an image", "JimMorrison", ErrUsernameNotExist},
	}

	// Perform tests.
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Stub and driver.
			db, teardown := newTestDB(t)
			defer teardown()

			m := ArticleModel{db}

			// When.
			_, err := m.Insert(tc.subName, tc.postType, tc.body, tc.title, tc.postedBy)

			// Want.
			if err != tc.wantError {
				t.Errorf("want %v; got %v", tc.wantError, err)
			}
		})
	}
}

func TestArticleModifyBody(t *testing.T) {
	// Testcases.
	tests := []struct {
		name      string
		articleID string
		username  string
		body      string
		wantError error
	}{
		{"Valid", "WX-78", "Morrissey", "The new content of the article.", nil},
		{"Not Text Post", "246o1", "Jonny", "http://www.google.com/", ErrArticleNotExist},
		{"Invalid Article", "zzzzz", "Jonny", "New content.", ErrArticleNotExist},
		{"Invalid username", "WX-78", "Jonny", "The new content of the article", ErrArticleNotExist},
	}

	// Perform tests.
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Stub and driver.
			db, teardown := newTestDB(t)
			defer teardown()

			m := ArticleModel{db}

			// When.
			err := m.ModifyBody(tc.articleID, tc.username, tc.body)

			// Want.
			if err != tc.wantError {
				t.Errorf("want %v; got %v", tc.wantError, err)
			}
		})
	}
}

func TestArticleDelete(t *testing.T) {
	// Testcases.
	tests := []struct {
		name      string
		articleID string
		username  string
		wantError error
	}{
		{"Valid", "WX-78", "Morrissey", nil},
		{"Invalid User", "WX-78", "Albarn", ErrArticleNotExist},
		{"Invalid Article", "heHeXd", "Jonny", ErrArticleNotExist},
	}

	// Perform tests.
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Stub and driver.
			db, teardown := newTestDB(t)
			defer teardown()

			m := ArticleModel{db}

			// When.
			err := m.Delete(tc.articleID, tc.username)

			// Want.
			if err != tc.wantError {
				t.Errorf("want %v; got %v", tc.wantError, err)
			}
		})
	}
}

func TestArticleGet(t *testing.T) {
	// Testcases.
	tests := []struct {
		name        string
		articleID   string
		wantArticle ArticleInfo
		wantError   error
	}{
		{"Invalid ArticleID", "abcde", ArticleInfo{}, ErrArticleNotExist},
		{"Valid", "246o1", ArticleInfo{"dankmeme", "246o1", "link", "http://www.google.com.tw/", "This is a nice website I discovered.", -1, "Jonny", time.Date(2012, 12, 12, 12, 12, 12, 0, time.UTC)}, nil},
	}

	// Perform tests.
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Stub and driver.
			db, teardown := newTestDB(t)
			defer teardown()

			m := ArticleModel{db}

			// When.
			article, err := m.Get(tc.articleID)

			// Want.
			if !reflect.DeepEqual(article, tc.wantArticle) {
				t.Errorf("want:\n %v\n got:\n %v", tc.wantArticle, article)
			}

			if err != tc.wantError {
				t.Errorf("want %v; got %v", tc.wantError, err)
			}
		})
	}
}

func TestArticleGetBySubreddit(t *testing.T) {
	// TODO
}

func TestArticleGetByUser(t *testing.T) {
	// TODO
}

func TestArticleGetSavedByUser(t *testing.T) {
	// TODO
}

// Combination test.
func TestInsertGetModifyDelete(t *testing.T) {
	subName, postType, body, title, postedBy := "dankmeme", "text", "is good body", "interesting title", "Jonny"
	newBody := "A new body!"
	article := ArticleInfo{}

	// Stub and driver.
	db, teardown := newTestDB(t)
	defer teardown()

	m := ArticleModel{db}

	// Perform tests.
	// Insert.
	articleID, err := m.Insert(subName, postType, body, title, postedBy)
	if err != nil {
		t.Fatal(err)
	}

	// Get.
	article, err = m.Get(articleID)
	if err != nil {
		t.Fatal(err)
	}

	if article.Subreddit != subName || article.Type != postType || article.Body != body || article.Title != title || article.PostedBy != postedBy {
		t.Fatal("The article inserted is different from the article gotten.")
	}

	// Modify.
	err = m.ModifyBody(articleID, postedBy, newBody)
	if err != nil {
		t.Fatal(err)
	}

	// Get again.
	article, err = m.Get(articleID)
	if err != nil {
		t.Fatal(err)
	}

	if article.Subreddit != subName || article.Type != postType || article.Body != newBody || article.Title != title || article.PostedBy != postedBy {
		t.Fatal("The article inserted is different from the article gotten.")
	}

	// Delete.
	err = m.Delete(articleID, postedBy)
	if err != nil {
		t.Fatal(err)
	}

	// Get again.
	_, err = m.Get(articleID)
	if err != ErrArticleNotExist {
		t.Errorf("want %v; got %v", ErrArticleNotExist, err)
	}

}
