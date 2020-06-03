package handlers

import (
	"net/http"
	"testing"
)

func TestArticle(t *testing.T) {
	// Stubs and drivers.
	h := newTestHandler(t)
	r := newTestRouter(t, "GET", "/article/{id}", h.Article)

	// Response bodies.
	validBody := `{"subreddit":"radiohead","articleID":"WX-78","type":"text",` +
		`"body":"I'm a reasonable man, get off my case.","title":"Tin can",` +
		`"points":42,"postedBy":"Bellamy","postedTime":"2019-11-13T00:00:00Z"}`
	notExistBody := `{"error":"the article does not exist"}`

	// Testset.
	tests := []struct {
		name     string
		urlPath  string
		wantCode int
		wantBody string
	}{
		{"Valid ID", "/article/WX-78", http.StatusOK, validBody},
		{"Non-existent ID", "/article/Wendy", http.StatusNotFound, notExistBody},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			code, body := r.testGet(t, tc.urlPath)

			if code != tc.wantCode {
				t.Errorf("want %d; got %d", tc.wantCode, code)
			}

			if body != tc.wantBody {
				t.Errorf("want:\n%s\ngot:\n%s\n\n", tc.wantBody, body)
			}
		})
	}
}
