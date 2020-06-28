package handlers

import (
	"net/http"
	"testing"
)

func TestGetArticleVote(t *testing.T) {
	// Stubs and drivers.
	h := newTestHandler(t)
	r := newTestRouter(t, "GET", "/me/karma/article/{id}", h.GetArticleVote)

	// Testcases.
	tests := []struct {
		name     string
		urlPath  string
		wantCode int
		wantBody string
	}{
		{"Valid", "/me/karma/article/88888", http.StatusOK, `{"action":"up"}`},
		{"Invalid", "/me/karma/article/77777", http.StatusOK, `{"action":"cancel"}`},
	}

	// Perform tests.
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// When.
			code, body := r.testHTTP(t, "GET", tc.urlPath, ``, "Jonny")

			// Want.
			if code != tc.wantCode {
				t.Errorf("want %v; got %v", tc.wantCode, code)
			}

			if body != tc.wantBody {
				t.Errorf("want %v; got %v", tc.wantBody, body)
			}
		})
	}
}

func TestGetCommentVote(t *testing.T) {
	// Stubs and drivers.
	h := newTestHandler(t)
	r := newTestRouter(t, "GET", "/me/karma/comment/{id}", h.GetCommentVote)

	// Testcases.
	tests := []struct {
		name     string
		urlPath  string
		wantCode int
		wantBody string
	}{
		{"Valid", "/me/karma/comment/88888", http.StatusOK, `{"action":"down"}`},
		{"Invalid", "/me/karma/comment/77777", http.StatusOK, `{"action":"cancel"}`},
	}

	// Perform tests.
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// When.
			code, body := r.testHTTP(t, "GET", tc.urlPath, ``, "Jonny")

			// Want.
			if code != tc.wantCode {
				t.Errorf("want %v; got %v", tc.wantCode, code)
			}

			if body != tc.wantBody {
				t.Errorf("want %v; got %v", tc.wantBody, body)
			}
		})
	}
}

func TestVoteArticle(t *testing.T) {
	// Stubs and drivers.
	h := newTestHandler(t)
	r := newTestRouter(t, "POST", "/karma/article/{id}", h.VoteArticle)

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
	r := newTestRouter(t, "POST", "/karma/comment/{id}", h.VoteComment)

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
