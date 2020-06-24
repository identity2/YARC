package handlers

import (
	"net/http"
	"testing"
)

func TestSearch(t *testing.T) {
	// Stubs and drivers.
	h := newTestHandler(t)
	r := newTestRouter(t, "GET", "/search", h.Search)

	// Testcases.
	tests := []struct {
		name     string
		urlPath  string
		wantCode int
		wantBody string
	}{
		{"Empty Query", "/search", http.StatusBadRequest, `{"error":"the query string cannot be empty"}`},
		{
			"Valid",
			"/search?q=ghost+town",
			http.StatusOK,
			`{"subreddits":[{"name":"PHP","description":"$_POST['cool']"},{"name":"golang","description":"if err != nil"}],` +
				`"articles":[{"subreddit":"Python","articleID":"1","type":"text","body":"","title":"title","points":3,"comments":1,"postedBy":"Collin",` +
				`"postedTime":"1995-11-24T00:00:00Z"}]}`,
		},
	}

	// Perform tests.
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// When.
			code, body := r.testGet(t, tc.urlPath)

			// Want.
			if code != tc.wantCode {
				t.Errorf("want %v; got %v", tc.wantCode, code)
			}

			if body != tc.wantBody {
				t.Errorf("want:\n%v\ngot:\n%v", tc.wantBody, body)
			}
		})
	}
}
