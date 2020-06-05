package handlers

import (
	"net/http"
	"testing"
)

func TestListArticle(t *testing.T) {
	// TODO.
}

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
				t.Errorf("want:\n%s\ngot:\n%s\n\n", tc.wantBody, body)
			}
		})
	}
}

func TestNewArticle(t *testing.T) {
	// Stubs and Drivers.
	h := newTestHandler(t)
	r := newTestRouter(t, "POST", "/article", h.NewArticle)

	// Request bodies.
	validReq := `{"subreddit":"radiohead","type":"text","body":"The body.","title":"Title","postedBy":"user"}`
	notExistReq := `{"subreddit":"dankmeme","type":"text","body":"The body.","title":"Title","postedBy":"user"}`
	invalidReq := `{"subreddit:"okComputer"}`

	// Testset.
	tests := []struct {
		name     string
		body     string
		wantCode int
		wantBody string
	}{
		{"Valid Request Body", validReq, http.StatusCreated, `{"articleID":"WX-78"}`},
		{"Subreddit In Request Body Not Exist", notExistReq, http.StatusBadRequest, `{"error":"the subreddit does not exist"}`},
		{"Invalid Request Body", invalidReq, http.StatusBadRequest, ``},
	}

	// Perform tests.
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// When
			code, body := r.testHTTP(t, "POST", "/article", tc.body, "mediocre_nothingness")

			// Want
			if code != tc.wantCode {
				t.Errorf("want %d; got %d", tc.wantCode, code)
			}

			if body != tc.wantBody {
				t.Errorf("want %s; got %s", tc.wantBody, body)
			}
		})
	}
}

func TestModifyArticle(t *testing.T) {
	// Stubs and Drivers.
	h := newTestHandler(t)
	r := newTestRouter(t, "PUT", "/article/{id}", h.ModifyArticle)

	// Request Bodies.
	validReq := `{"body": "I'm a reasonable man."}`
	invalidReq := `{body: "get off my case."}`

	// Testcases.
	tests := []struct {
		name     string
		urlPath  string
		body     string
		wantCode int
		wantBody string
	}{
		{"Valid URL Valid body", "/article/WX-78", validReq, http.StatusNoContent, ``},
		{"Invalid URL Valid body", "/article/Dijkstra", validReq, http.StatusNotFound, `{"error":"the article does not exist"}`},
		{"Valid URL Invalid body", "/article/WX-78", invalidReq, http.StatusBadRequest, ``},
	}

	// Perform tests.
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// When
			code, body := r.testHTTP(t, "PUT", tc.urlPath, tc.body, "mediocre_nothingness")

			// Want
			if code != tc.wantCode {
				t.Errorf("want %d; got %d", tc.wantCode, code)
			}

			if body != tc.wantBody {
				t.Errorf("want %s; got %s", tc.wantBody, body)
			}
		})
	}
}

func TestDeleteArticle(t *testing.T) {
	// Stubs and Drivers
	h := newTestHandler(t)
	r := newTestRouter(t, "DELETE", "/article/{id}", h.DeleteArticle)

	// Testcases.
	tests := []struct {
		name     string
		urlPath  string
		wantCode int
		wantBody string
	}{
		{"Valid URL", "/article/WX-78", http.StatusNoContent, ``},
		{"Invalid URL", "/article/Dijkstra", http.StatusNotFound, `{"error":"the article does not exist"}`},
	}

	// Performing tests.
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// When.
			code, body := r.testHTTP(t, "DELETE", tc.urlPath, "", "mediocre_nothingness")

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
