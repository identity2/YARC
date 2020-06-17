package handlers

import (
	"net/http"
	"testing"
)

func TestUser(t *testing.T) {
	// Stubs and drivers.
	h := newTestHandler(t)
	r := newTestRouter(t, "GET", "/user/{username}", h.User)

	// Response bodies.
	validBody := `{"username":"Jonny","karma":-1,"bio":"Who's in the bunker?","joinTime":"1999-09-05T10:33:52Z"}`
	notExistBody := `{"error":"the username does not exist"}`

	// Testcases.
	tests := []struct {
		name     string
		urlPath  string
		wantCode int
		wantBody string
	}{
		{"Valid ID", "/user/Jonny", http.StatusOK, validBody},
		{"Non-existent Username", "/user/Martin", http.StatusNotFound, notExistBody},
	}

	// Perform Tests.
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// When.
			code, body := r.testGet(t, tc.urlPath)

			// Want.
			if code != tc.wantCode {
				t.Errorf("want %d; got %d", tc.wantCode, code)
			}

			if body != tc.wantBody {
				t.Errorf("want:\n%s\ngot:\n%s", tc.wantBody, body)
			}
		})
	}
}

func TestModifyBio(t *testing.T) {
	// Stub and driver.
	h := newTestHandler(t)
	r := newTestRouter(t, "PUT", "/me/bio", h.ModifyBio)

	// Testcases.
	tests := []struct {
		name     string
		urlPath  string
		body     string
		wantCode int
		wantBody string
	}{
		{"Valid", "/me/bio", `{"bio":"About me!"}`, http.StatusNoContent, ``},
		{"Invalid", "/me/bio", `{"bio":"Dirty Dan"}`, http.StatusBadRequest, `{"error":"the short bio should contain at most 256 characters"}`},
	}

	// Perform tests.
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// When.
			code, body := r.testHTTP(t, "PUT", "/me/bio", tc.body, "mediocre_nothingness")

			// Want.
			if code != tc.wantCode {
				t.Errorf("want %d; got %d", tc.wantCode, code)
			}

			if body != tc.wantBody {
				t.Errorf("want:\n%v\ngot:\n%v", tc.wantBody, body)
			}
		})
	}
}

func TestSaveArticle(t *testing.T) {
	// Stub and driver.
	h := newTestHandler(t)
	r := newTestRouter(t, "POST", "/me/save/{articleID}", h.SaveArticle)

	// Testcases.
	tests := []struct {
		name     string
		urlPath  string
		wantCode int
	}{
		{"Valid", "/me/save/55688", http.StatusCreated},
		{"Invalid", "/me/save/77777", http.StatusNotFound},
	}

	// Perform tests.
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// When.
			code, _ := r.testHTTP(t, "POST", tc.urlPath, "", "PatrickStar")

			// Want.
			if code != tc.wantCode {
				t.Errorf("want %v; got %v", tc.wantCode, code)
			}
		})
	}
}

func TestUnsaveArticle(t *testing.T) {
	// Stub and driver.
	h := newTestHandler(t)
	r := newTestRouter(t, "DELETE", "/me/save/{articleID}", h.UnsaveArticle)

	// Testcases.
	tests := []struct {
		name     string
		urlPath  string
		wantCode int
	}{
		{"Valid", "/me/save/cp3cl3", http.StatusNoContent},
		{"Invalid", "/me/save/5566", http.StatusNotFound},
	}

	// Perform tests.
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// When.
			code, _ := r.testHTTP(t, "DELETE", tc.urlPath, "", "SpongeBob")

			// Want.
			if code != tc.wantCode {
				t.Errorf("want %v; got %v", tc.wantCode, code)
			}
		})
	}
}

func TestJoinSubreddit(t *testing.T) {
	// Stub and driver.
	h := newTestHandler(t)
	r := newTestRouter(t, "POST", "/me/join/{subreddit}", h.JoinSubreddit)

	// Testcases.
	tests := []struct {
		name     string
		urlPath  string
		wantCode int
	}{
		{"Valid", "/me/join/game_design", http.StatusCreated},
		{"Invalid", "/me/join/gamedev", http.StatusNotFound},
	}

	// Perform tests.
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// When.
			code, _ := r.testHTTP(t, "POST", tc.urlPath, "", "nirvana")

			// Want.
			if code != tc.wantCode {
				t.Errorf("want %v; got %v", tc.wantCode, code)
			}
		})
	}
}

func TestLeaveSubreddit(t *testing.T) {
	// Stub and driver.
	h := newTestHandler(t)
	r := newTestRouter(t, "DELETE", "/me/join/{subreddit}", h.LeaveSubreddit)

	// Testcases.
	tests := []struct {
		name     string
		urlPath  string
		wantCode int
	}{
		{"Valid", "/me/join/coldplaycirclejerk", http.StatusNoContent},
		{"Invalid", "/me/join/CSMajor", http.StatusNotFound},
	}

	// Perform tests.
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// When.
			code, _ := r.testHTTP(t, "DELETE", tc.urlPath, "", "WarGod69")

			// Want.
			if code != tc.wantCode {
				t.Errorf("want %v; got %v", tc.wantCode, code)
			}
		})
	}
}
