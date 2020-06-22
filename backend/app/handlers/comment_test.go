package handlers

import (
	"net/http"
	"testing"
)

func TestListComment(t *testing.T) {
	// Stubs and drivers.
	h := newTestHandler(t)
	r := newTestRouter(t, "GET", "/comment", h.ListComment)

	// Response bodies.
	getByArticleBody := `{"comments":[{"subreddit":"blur","articleID":"mAGiC","commentID":"wHiPp","body":"I was only 21 when I watched it on TV.","points":689,"postedBy":"GorillaZ","postedTime":"1999-09-09T00:00:00Z"},` +
		`{"subreddit":"thesmiths","articleID":"tENtoN","commentID":"tRUcC","body":"The pain was enough to make a shy bald buddist reflect.","points":333,"postedBy":"BernardSumner","postedTime":"2012-04-01T00:00:00Z"}]}`
	getByUserBody := `{"comments":[{"subreddit":"blur","articleID":"mAGiC","commentID":"wHiPp","body":"I was only 21 when I watched it on TV.","points":689,"postedBy":"GorillaZ","postedTime":"1999-09-09T00:00:00Z"}]}`

	// Testcases.
	tests := []struct {
		name     string
		urlPath  string
		wantCode int
		wantBody string
	}{
		{"By Article", "/comment?after=451542212&limit=3&ofArticle=124215545", http.StatusOK, getByArticleBody},
		{"By User", "/comment?postedBy=god", http.StatusOK, getByUserBody},
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
				t.Errorf("want %v; got %v", tc.wantBody, body)
			}
		})
	}
}

func TestComment(t *testing.T) {
	// Stubs and drivers.
	h := newTestHandler(t)
	r := newTestRouter(t, "GET", "/comment/{id}", h.Comment)

	// Testcases.
	tests := []struct {
		name     string
		urlPath  string
		wantCode int
		wantBody string
	}{
		{"Valid", "/comment/12345", http.StatusOK, `{"subreddit":"radiohead","articleID":"WX-78","commentID":"12345","body":"Flan","points":3,"postedBy":"Jonny","postedTime":"2003-11-24T00:00:00Z"}`},
		{"Not Exist", "/comment/24601", http.StatusNotFound, `{"error":"the comment does not exist"}`},
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
				t.Errorf("want %v; got %v", tc.wantBody, body)
			}
		})
	}
}

func TestNewComment(t *testing.T) {
	// Drivers and stubs.
	h := newTestHandler(t)
	r := newTestRouter(t, "POST", "/comment", h.NewComment)

	// Request bodies.
	validReq := `{"articleID":"sdf76","body":"good"}`
	invalidReq := `{articleID:88888,body:ok}`
	articleNotExist := `{"articleID":"88888","body":"ok"}`

	// Testcases.
	tests := []struct {
		name     string
		req      string
		wantCode int
		wantBody string
	}{
		{"Valid", validReq, http.StatusCreated, `{"commentID":"Maxwell"}`},
		{"Invalid", invalidReq, http.StatusBadRequest, ``},
		{"Not Exist Article", articleNotExist, http.StatusBadRequest, `{"error":"the article does not exist"}`},
	}

	// Perform tests.
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// When.
			code, body := r.testHTTP(t, "POST", "/comment", tc.req, "mediocre_nothingness")

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

func TestModifyComment(t *testing.T) {
	// Stubs and drivers.
	h := newTestHandler(t)
	r := newTestRouter(t, "PUT", "/comment/{id}", h.ModifyComment)

	// Testcases.
	tests := []struct {
		name     string
		urlPath  string
		req      string
		wantCode int
		wantBody string
	}{
		{"Valid Body", "/comment/Wendy", `{"body":"the new body"}`, http.StatusNoContent, ``},
		{"Invalid Body", "/comment/Wendy", `{body:ok}`, http.StatusBadRequest, ``},
		{"Comment Not Exist", "/comment/Willow", `{"body":"ok"}`, http.StatusNotFound, `{"error":"the comment does not exist"}`},
	}

	// Perform tests.
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// When.
			code, body := r.testHTTP(t, "PUT", tc.urlPath, tc.req, "mediocre_nothingness")

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

func TestDeleteComment(t *testing.T) {
	// Stubs and drivers.
	h := newTestHandler(t)
	r := newTestRouter(t, "DELETE", "/comment/{id}", h.DeleteComment)

	// Testcases.
	tests := []struct {
		name     string
		urlPath  string
		wantCode int
		wantBody string
	}{
		{"Valid", "/comment/Wickerbottom", http.StatusNoContent, ``},
		{"Comment Not Exist", "/comment/Willow", http.StatusNotFound, `{"error":"the comment does not exist"}`},
	}

	// Perform tests.
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// When.
			code, body := r.testHTTP(t, "DELETE", tc.urlPath, ``, "mediocre_nothingness")

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
