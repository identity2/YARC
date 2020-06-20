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
		{"Valid", "246o1", ArticleInfo{"dankmeme", "246o1", "link", "http://www.google.com.tw/", "This is a nice website I discovered.", -1, 3, "Jonny", time.Date(2012, 12, 12, 12, 12, 12, 0, time.UTC)}, nil},
		{"Valid No Comments No points", "RgMG_RTSvkQa", ArticleInfo{"PHP", "RgMG_RTSvkQa", "text", "Abominations, I think Visual Basic is just too good to be real!", "One of my dominations", 0, 0, "Jonny", time.Date(2020, 11, 21, 19, 20, 2, 0, time.UTC)}, nil},
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

func TestArticleGetAll(t *testing.T) {
	// Testcases.
	tests := []struct {
		name           string
		afterArticleID string
		sortedBy       string
		limit          int
		wantArticleIDs []string
		wantError      error
	}{
		{"Exceed Max Limit", "RgMG_RTSvkQ6", "hot", 101, []string{}, ErrLimitInvalid},
		{"Exceed Min Limit", "RgMG_RTSvkQ6", "new", -1, []string{}, ErrLimitInvalid},
		{"Invalid sortedBy", "RgMG_RTSvkQ6", "hey", 20, []string{}, ErrSortedByInvalid},
		{"Sort by hot in range", "RgMG_RTSvkQ7", "hot", 3, []string{"RgMG_RTSvkQ8", "IpX1779", "t09o39"}, nil},
		{"Sort by hot from start", "", "hot", 2, []string{"IpX177", "WX-78"}, nil},
		{"Sort by hot pass end", "RgMG_RTSvkQ", "hot", 5, []string{"246o1"}, nil},
		{"Sort by new in range", "t09o39", "new", 3, []string{"t09o3", "246o19", "246o1"}, nil},
		{"Sort by new from start", "", "new", 3, []string{"RgMG_RTSvkQa", "RgMG_RTSvkQ6", "RgMG_RTSvkQ7"}, nil},
		{"Sort by new pass end", "WX-789", "new", 3, []string{"WX-78"}, nil},
		{"Sort by old in range", "t09o3", "old", 4, []string{"t09o39", "IpX177", "IpX1779", "RgMG_RTSvkQ"}, nil},
		{"Sort by old from start", "", "old", 4, []string{"WX-78", "WX-789", "246o1", "246o19"}, nil},
		{"Sort by old pass end", "RgMG_RTSvkQ6", "old", 4, []string{"RgMG_RTSvkQa"}, nil},
	}

	// Perform tests.
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Stub and driver.
			db, teardown := newTestDB(t)
			defer teardown()

			m := ArticleModel{db}

			// When.
			res, err := m.GetAll(tc.afterArticleID, tc.sortedBy, tc.limit)

			// Want.
			if tc.wantError != err {
				t.Errorf("want %v; got %v", tc.wantError, err)
			}

			if len(tc.wantArticleIDs) != len(res) {
				t.Fatalf("want len %v; got len %v", len(tc.wantArticleIDs), len(res))
			}
			for i := range tc.wantArticleIDs {
				if tc.wantArticleIDs[i] != res[i].ArticleID {
					t.Errorf("want %v; got %v", tc.wantArticleIDs[i], res[i].ArticleID)
				}
			}
		})
	}
}

func TestArticleGetBySubscribed(t *testing.T) {
	// Testcases.
	tests := []struct {
		name           string
		username       string
		afterArticleID string
		sortedBy       string
		limit          int
		wantArticleIDs []string
		wantError      error
	}{
		{"Exceed Max Limit", "Morrissey", "RgMG_RTSvkQ6", "hot", 101, []string{}, ErrLimitInvalid},
		{"Exceed Min Limit", "Morrissey", "RgMG_RTSvkQ6", "new", -1, []string{}, ErrLimitInvalid},
		{"Invalid sortedBy", "Morrissey", "RgMG_RTSvkQ6", "hey", 20, []string{}, ErrSortedByInvalid},
		{"Sort by hot in range", "Jonny", "t09o3", "hot", 3, []string{"246o19", "WX-789", "246o1"}, nil},
		{"Sort by hot from start", "Jonny", "", "hot", 2, []string{"WX-78", "t09o39"}, nil},
		{"Sort by hot pass end", "Jonny", "246o19", "hot", 5, []string{"WX-789", "246o1"}, nil},
		{"Sort by new in range", "Albarn", "RgMG_RTSvkQa", "new", 2, []string{"RgMG_RTSvkQ6", "RgMG_RTSvkQ7"}, nil},
		{"Sort by new from start", "Albarn", "", "new", 3, []string{"RgMG_RTSvkQa", "RgMG_RTSvkQ6", "RgMG_RTSvkQ7"}, nil},
		{"Sort by new pass end", "Albarn", "IpX1779", "new", 3, []string{"IpX177"}, nil},
		{"Sort by old in range", "Albarn", "IpX177", "old", 4, []string{"IpX1779", "RgMG_RTSvkQ", "RgMG_RTSvkQ9", "RgMG_RTSvkQ8"}, nil},
		{"Sort by old from start", "Albarn", "", "old", 4, []string{"IpX177", "IpX1779", "RgMG_RTSvkQ", "RgMG_RTSvkQ9"}, nil},
		{"Sort by old pass end", "Albarn", "RgMG_RTSvkQ6", "old", 4, []string{"RgMG_RTSvkQa"}, nil},
	}

	// Perform tests.
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Stub and driver.
			db, teardown := newTestDB(t)
			defer teardown()

			m := ArticleModel{db}

			// When.
			res, err := m.GetBySubscribed(tc.username, tc.afterArticleID, tc.sortedBy, tc.limit)

			// Want.
			if tc.wantError != err {
				t.Errorf("want %v; got %v", tc.wantError, err)
			}

			if len(tc.wantArticleIDs) != len(res) {
				t.Fatalf("want len %v; got len %v", len(tc.wantArticleIDs), len(res))
			}
			for i := range tc.wantArticleIDs {
				if tc.wantArticleIDs[i] != res[i].ArticleID {
					t.Errorf("want %v; got %v", tc.wantArticleIDs[i], res[i].ArticleID)
				}
			}
		})
	}
}

func TestArticleGetBySubreddit(t *testing.T) {
	// Testcases.
	tests := []struct {
		name           string
		subName        string
		afterArticleID string
		sortedBy       string
		limit          int
		wantArticleIDs []string
		wantError      error
	}{
		{"Exceed Max Limit", "PHP", "RgMG_RTSvkQ6", "hot", 101, []string{}, ErrLimitInvalid},
		{"Exceed Min Limit", "PHP", "RgMG_RTSvkQ6", "new", -1, []string{}, ErrLimitInvalid},
		{"Invalid sortedBy", "PHP", "RgMG_RTSvkQ6", "hey", 20, []string{}, ErrSortedByInvalid},
		{"Sort by hot in range", "PHP", "IpX177", "hot", 3, []string{"RgMG_RTSvkQ9", "RgMG_RTSvkQa", "RgMG_RTSvkQ6"}, nil},
		{"Sort by hot from start", "PHP", "", "hot", 2, []string{"IpX177", "RgMG_RTSvkQ9"}, nil},
		{"Sort by hot pass end", "PHP", "RgMG_RTSvkQ8", "hot", 5, []string{"IpX1779", "RgMG_RTSvkQ"}, nil},
		{"Sort by new in range", "dankmeme", "246o1", "new", 2, []string{"WX-789", "WX-78"}, nil},
		{"Sort by new from start", "dankmeme", "", "new", 3, []string{"246o19", "246o1", "WX-789"}, nil},
		{"Sort by new pass end", "dankmeme", "WX-78", "new", 3, []string{}, nil},
		{"Sort by old in range", "PHP", "RgMG_RTSvkQ", "old", 2, []string{"RgMG_RTSvkQ9", "RgMG_RTSvkQ8"}, nil},
		{"Sort by old from start", "PHP", "77777", "old", 4, []string{"IpX177", "IpX1779", "RgMG_RTSvkQ", "RgMG_RTSvkQ9"}, nil},
		{"Sort by old pass end", "PHP", "RgMG_RTSvkQ6", "old", 4, []string{"RgMG_RTSvkQa"}, nil},
	}

	// Perform tests.
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Stub and driver.
			db, teardown := newTestDB(t)
			defer teardown()

			m := ArticleModel{db}

			// When.
			res, err := m.GetBySubreddit(tc.subName, tc.afterArticleID, tc.sortedBy, tc.limit)

			// Want.
			if tc.wantError != err {
				t.Errorf("want %v; got %v", tc.wantError, err)
			}

			if len(tc.wantArticleIDs) != len(res) {
				t.Fatalf("want len %v; got len %v", len(tc.wantArticleIDs), len(res))
			}
			for i := range tc.wantArticleIDs {
				if tc.wantArticleIDs[i] != res[i].ArticleID {
					t.Errorf("want %v; got %v", tc.wantArticleIDs[i], res[i].ArticleID)
				}
			}
		})
	}
}

func TestArticleGetByUser(t *testing.T) {
	// Testcases.
	tests := []struct {
		name           string
		username       string
		afterArticleID string
		sortedBy       string
		limit          int
		wantArticleIDs []string
		wantError      error
	}{
		{"Exceed Max Limit", "Jonny", "RgMG_RTSvkQ6", "hot", 101, []string{}, ErrLimitInvalid},
		{"Exceed Min Limit", "Jonny", "RgMG_RTSvkQ6", "new", -1, []string{}, ErrLimitInvalid},
		{"Invalid sortedBy", "Jonny", "RgMG_RTSvkQ6", "hey", 20, []string{}, ErrSortedByInvalid},
		{"Sort by hot in range", "Jonny", "IpX1779", "hot", 3, []string{"246o19", "RgMG_RTSvkQ", "246o1"}, nil},
		{"Sort by hot from start", "Jonny", "", "hot", 2, []string{"IpX177", "RgMG_RTSvkQ9"}, nil},
		{"Sort by hot pass end", "Jonny", "246o19", "hot", 5, []string{"RgMG_RTSvkQ", "246o1"}, nil},
		{"Sort by new in range", "Jonny", "IpX1779", "new", 2, []string{"IpX177", "246o19"}, nil},
		{"Sort by new from start", "Jonny", "", "new", 3, []string{"RgMG_RTSvkQa", "RgMG_RTSvkQ6", "RgMG_RTSvkQ7"}, nil},
		{"Sort by new pass end", "Jonny", "246o19", "new", 3, []string{"246o1"}, nil},
		{"Sort by old in range", "Morrissey", "WX-78", "old", 1, []string{"WX-789"}, nil},
		{"Sort by old from start", "Morrissey", "", "old", 2, []string{"WX-78", "WX-789"}, nil},
		{"Sort by old pass end", "Morrissey", "WX-78", "old", 4, []string{"WX-789"}, nil},
	}

	// Perform tests.
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Stub and driver.
			db, teardown := newTestDB(t)
			defer teardown()

			m := ArticleModel{db}

			// When.
			res, err := m.GetByUser(tc.username, tc.afterArticleID, tc.sortedBy, tc.limit)

			// Want.
			if tc.wantError != err {
				t.Errorf("want %v; got %v", tc.wantError, err)
			}

			if len(tc.wantArticleIDs) != len(res) {
				t.Fatalf("want len %v; got len %v", len(tc.wantArticleIDs), len(res))
			}
			for i := range tc.wantArticleIDs {
				if tc.wantArticleIDs[i] != res[i].ArticleID {
					t.Errorf("want %v; got %v", tc.wantArticleIDs[i], res[i].ArticleID)
				}
			}
		})
	}
}

func TestArticleGetSavedByUser(t *testing.T) {
	// Testcases.
	tests := []struct {
		name           string
		username       string
		afterArticleID string
		sortedBy       string
		limit          int
		wantArticleIDs []string
		wantError      error
	}{
		{"Exceed Max Limit", "Morrissey", "RgMG_RTSvkQ6", "hot", 101, []string{}, ErrLimitInvalid},
		{"Exceed Min Limit", "Morrissey", "RgMG_RTSvkQ6", "new", -1, []string{}, ErrLimitInvalid},
		{"Invalid sortedBy", "Morrissey", "RgMG_RTSvkQ6", "hey", 20, []string{}, ErrSortedByInvalid},
		{"Sort by hot in range", "Albarn", "t09o3", "hot", 2, []string{"246o19", "RgMG_RTSvkQ"}, nil},
		{"Sort by hot from start", "Albarn", "", "hot", 3, []string{"IpX177", "RgMG_RTSvkQ9", "t09o3"}, nil},
		{"Sort by hot pass end", "Albarn", "246o19", "hot", 5, []string{"RgMG_RTSvkQ"}, nil},
		{"Sort by new in range", "Jonny", "t09o3", "new", 2, []string{"246o19", "246o1"}, nil},
		{"Sort by new from start", "Jonny", "", "new", 2, []string{"IpX177", "t09o3"}, nil},
		{"Sort by new pass end", "Jonny", "246o19", "new", 3, []string{"246o1"}, nil},
		{"Sort by old in range", "Albarn", "IpX177", "old", 1, []string{"RgMG_RTSvkQ"}, nil},
		{"Sort by old from start", "Albarn", "", "old", 3, []string{"246o19", "t09o3", "IpX177"}, nil},
		{"Sort by old pass end", "Albarn", "RgMG_RTSvkQ", "old", 4, []string{"RgMG_RTSvkQ9"}, nil},
	}

	// Perform tests.
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Stub and driver.
			db, teardown := newTestDB(t)
			defer teardown()

			m := ArticleModel{db}

			// When.
			res, err := m.GetSavedByUser(tc.username, tc.afterArticleID, tc.sortedBy, tc.limit)

			// Want.
			if tc.wantError != err {
				t.Errorf("want %v; got %v", tc.wantError, err)
			}

			if len(tc.wantArticleIDs) != len(res) {
				t.Fatalf("want len %v; got len %v", len(tc.wantArticleIDs), len(res))
			}
			for i := range tc.wantArticleIDs {
				if tc.wantArticleIDs[i] != res[i].ArticleID {
					t.Errorf("want %v; got %v", tc.wantArticleIDs[i], res[i].ArticleID)
				}
			}
		})
	}
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
