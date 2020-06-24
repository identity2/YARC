package handlers

import (
	"net/http"
	"testing"
)

func TestVoteArticle(t *testing.T) {
	// Stubs and drivers.
	h := newTestHandler(t)
	r := newTestRouter(t, "POST", "/karma/article/{articleID}", h.VoteArticle)

	// Testcases.
	tests := []struct {
		name     string
		urlPath  string
		wantCode int
	}{
		{"Valid", "/karma/article/88888?action=up", http.StatusNoContent},
		{"No Action", "/karma/article/88888", http.StatusBadRequest},
		{"Invalid", "/karma/article/0?action=down", http.StatusBadRequest},
	}

	// Perform tests.
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// When.
			code, _ := r.testHTTP(t, "POST", tc.urlPath, ``, "Jonny")

			// Want.
			if code != tc.wantCode {
				t.Errorf("want %v; got %v", tc.wantCode, code)
			}
		})
	}
}

func TestVoteComment(t *testing.T) {
	// Stubs and drivers.
	h := newTestHandler(t)
	r := newTestRouter(t, "POST", "/karma/comment/{commentID}", h.VoteComment)

	// Testcases.
	tests := []struct {
		name     string
		urlPath  string
		wantCode int
	}{
		{"Valid", "/karma/comment/88888?action=down", http.StatusNoContent},
		{"No Action", "/karma/comment/88888", http.StatusBadRequest},
		{"Invalid", "/karma/comment/0?action=up", http.StatusBadRequest},
	}

	// Perform tests.
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// When.
			code, _ := r.testHTTP(t, "POST", tc.urlPath, ``, "Jonny")

			// Want.
			if code != tc.wantCode {
				t.Errorf("want %v; got %v", tc.wantCode, code)
			}
		})
	}
}
