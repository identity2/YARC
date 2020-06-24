package models

import "testing"

func TestSearchGetResult(t *testing.T) {
	// Testcase.
	query := "t"

	// Want results.
	wantSubreddits := []string{"PHP", "dankmeme", "golang"}
	wantArticles := []string{"IpX177", "WX-78", "RgMG_RTSvkQa", "RgMG_RTSvkQ6", "RgMG_RTSvkQ7", "IpX1779", "246o19", "WX-789", "246o1"}

	// Stub and driver.
	db, teardown := newTestDB(t)
	defer teardown()

	m := SearchModel{db}

	// When.
	res, err := m.GetResult(query)
	if err != nil {
		t.Fatal(err)
	}

	// Want.
	if len(res.Subreddits) != len(wantSubreddits) {
		t.Fatalf("want len %v; got len %v", len(wantSubreddits), len(res.Subreddits))
	}

	for i := range wantSubreddits {
		if res.Subreddits[i].Name != wantSubreddits[i] {
			t.Errorf("want %v; got %v", wantSubreddits[i], res.Subreddits[i].Name)
		}
	}

	if len(res.Articles) != len(wantArticles) {
		t.Fatalf("want len %v; got len %v", len(wantArticles), len(res.Articles))
	}

	for i := range wantArticles {
		if res.Articles[i].ArticleID != wantArticles[i] {
			t.Errorf("want %v; got %v", wantArticles[i], res.Articles[i].ArticleID)
		}
	}
}
