package handlers

import (
	"net/http"
	"testing"
)

func TestSubreddit(t *testing.T) {
	// Stubs and drivers.
	h := newTestHandler(t)
	r := newTestRouter(t, "GET", "/subreddit/{name}", h.Subreddit)

	// Response body.
	validResp := `{"name":"radiohead","description":"Dedicated to all human beings."}`
	notFoundResp := `{"error":"the subreddit does not exist"}`

	// Testcases.
	tests := []struct {
		name     string
		urlPath  string
		wantCode int
		wantBody string
	}{
		{"Valid URL", "/subreddit/radiohead", http.StatusOK, validResp},
		{"Invalid URL", "/subreddit/muse", http.StatusNotFound, notFoundResp},
	}

	// Perform tests.
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// When.
			code, body := r.testGet(t, tc.urlPath)

			// Want.
			if code != tc.wantCode {
				t.Errorf("want %d; got %d", tc.wantCode, code)
			}

			if body != tc.wantBody {
				t.Errorf("want %s; got %s", tc.wantBody, body)
			}
		})
	}
}

func TestNewSubreddit(t *testing.T) {
	// Stubs and drivers.
	h := newTestHandler(t)
	r := newTestRouter(t, "POST", "/subreddit", h.NewSubreddit)

	// Request bodies.
	validBody := `{"name":"muse","description":"Crucifies my enemy."}`
	invalidBody := `bad body`

	// Testcases.
	tests := []struct {
		name     string
		body     string
		wantCode int
	}{
		{"Valid Request Body", validBody, http.StatusCreated},
		{"Invalid Request Body", invalidBody, http.StatusBadRequest},
	}

	// Perform testcases.
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// When.
			code, _ := r.testHTTP(t, "POST", "/subreddit", tc.body, "nothing")

			// Want.
			if code != tc.wantCode {
				t.Errorf("want %d; got %d", tc.wantCode, code)
			}
		})
	}
}

func TestTrending(t *testing.T) {
	// Stub and driver.
	h := newTestHandler(t)
	r := newTestRouter(t, "GET", "/trending", h.TrendingSubreddit)

	// When.
	code, body := r.testGet(t, "/trending?limit=2")

	// Want.
	wantCode := http.StatusOK
	wantBody := `{"subreddits":[{"name":"radiohead","description":"Dedicated to all human beings."},{"name":"underrated","description":"Let down and hanging around."}]}`
	if code != wantCode {
		t.Errorf("want %v; got %v", wantCode, code)
	}

	if body != wantBody {
		t.Errorf("want\n%v\ngot\n%v", wantBody, body)
	}
}
