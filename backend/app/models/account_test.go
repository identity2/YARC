package models

import (
	"reflect"
	"testing"
	"time"
)

func TestAccountInsert(t *testing.T) {
	// Testcases.
	tests := []struct {
		name      string
		username  string
		password  string
		email     string
		wantError error
	}{
		{"Valid", "islander", "password", "islander@maple.co.kr", nil},
		{"Duplicate Username", "Jonny", "myPassword", "jonny@gmail.com", ErrUsernameDup},
		{"Duplicate Email", "lubeDonuts", "my_password", "morty@thesmiths.com.uk", ErrEmailDup},
		{"Username Invalid", "", "the_password123", "test@ing.com", ErrUsernameInvalid},
		{"Password Invalid", "validUser", "pw", "t@c.com.tw", ErrPasswordInvalid},
		{"Email Invalid", "good_user_2D", "password", "feels_good_inc@malencholy.", ErrEmailInvalid},
	}

	// Perform Tests.
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Stub and driver.
			db, teardown := newTestDB(t)
			defer teardown()

			m := AccountModel{db}

			// When.
			err := m.Insert(tc.username, tc.email, tc.password)

			if err != tc.wantError {
				t.Errorf("want %v; got %v", tc.wantError, err)
			}
		})
	}
}

func TestAccountAuthenticate(t *testing.T) {
	// Testcases.
	tests := []struct {
		name      string
		username  string
		password  string
		wantError error
	}{
		{"Valid", "Jonny", "password", nil},
		{"Unregistered Username", "Dijkstra", "password", ErrUsernameNotExist},
		{"Wrong Password", "Morrissey", "only_weakness", ErrAccountInfoMismatch},
	}

	// Perform tests.
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Stub and driver.
			db, teardown := newTestDB(t)
			defer teardown()

			m := AccountModel{db}

			// When.
			err := m.Authenticate(tc.username, tc.password)

			// Want.
			if err != tc.wantError {
				t.Errorf("want %v; got %v", tc.wantError, err)
			}
		})
	}
}

func TestAccountModifyBio(t *testing.T) {
	// Testcases.
	tests := []struct {
		name      string
		username  string
		newBio    string
		wantError error
	}{
		{"Valid", "Morrissey", "I was bored before I even begun.", nil},
		{"Invalid Username", "Johnny_Rotten", "No feelings, no feelings, no feelings.", ErrUsernameNotExist},
		{"Invalid Bio", "Jonny", "I am very frightened when whenever we perform the song idioteque because master Thom Yorke always think I mess up with the beep beep boop in the song. However, I actually never mess up anything because I am officially one of the top five musicians in our entire band!", ErrBioInvalid},
	}

	// Perform tests.
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Stub and driver.
			db, teardown := newTestDB(t)
			defer teardown()

			m := AccountModel{db}

			// When.
			err := m.ModifyBio(tc.username, tc.newBio)

			// Want.
			if err != tc.wantError {
				t.Errorf("want %v; got %v", tc.wantError, err)
			}
		})
	}
}

func TestAccountGetSaveState(t *testing.T) {
	// Testcases.
	tests := []struct {
		name      string
		username  string
		articleID string
		wantRes   bool
	}{
		{"Valid", "Jonny", "246o1", true},
		{"Not Saved", "Morrissey", "246o1", false},
		{"Invalid", "Brandon", "Sanderson", false},
	}

	// Perform tests.
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Stub and driver.
			db, teardown := newTestDB(t)
			defer teardown()

			m := AccountModel{db}

			// When.
			res := m.GetSaveState(tc.username, tc.articleID)

			// Want.
			if res != tc.wantRes {
				t.Errorf("want %v; got %v", tc.wantRes, res)
			}
		})
	}
}

func TestAccountSaveArticle(t *testing.T) {
	// Testcases.
	tests := []struct {
		name      string
		articleID string
		username  string
		wantError error
	}{
		{"Valid", "246o1", "Morrissey", nil},
		{"Invalid Article", "77777", "Jonny", ErrArticleNotExist},
		{"Already Saved", "246o1", "Jonny", ErrArticleAlreadySaved},
		{"Invalid Username", "246o1", "Martin", ErrUsernameNotExist},
	}

	// Perform tests.
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Stub and driver.
			db, teardown := newTestDB(t)
			defer teardown()

			m := AccountModel{db}

			// When.
			err := m.SaveArticle(tc.articleID, tc.username)

			// Want.
			if err != tc.wantError {
				t.Errorf("want %v; got %v", tc.wantError, err)
			}
		})
	}
}

func TestAccountUnsaveArticle(t *testing.T) {
	// Testcases.
	tests := []struct {
		name      string
		articleID string
		username  string
		wantError error
	}{
		{"Valid", "IpX177", "Albarn", nil},
		{"Invalid Article", "77777", "Jonny", ErrArticleNotSaved},
		{"Invalid Username", "IpX177", "Morrissey", ErrArticleNotSaved},
	}

	// Perform tests.
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Stub and driver.
			db, teardown := newTestDB(t)
			defer teardown()

			m := AccountModel{db}

			// When.
			err := m.UnsaveArticle(tc.articleID, tc.username)

			// Want.
			if err != tc.wantError {
				t.Errorf("want %v; got %v", tc.wantError, err)
			}
		})
	}
}

func TestAccountGetJoinState(t *testing.T) {
	// Testcases.
	tests := []struct {
		name     string
		username string
		subName  string
		wantRes  bool
	}{
		{"Valid", "Jonny", "meirl", true},
		{"Not Joined", "Morrissey", "meirl", false},
		{"Invalid", "Brandon", "Sanderson", false},
	}

	// Perform tests.
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Stub and driver.
			db, teardown := newTestDB(t)
			defer teardown()

			m := AccountModel{db}

			// When.
			res := m.GetJoinState(tc.username, tc.subName)

			// Want.
			if res != tc.wantRes {
				t.Errorf("want %v; got %v", tc.wantRes, res)
			}
		})
	}
}

func TestAccountJoinSubreddit(t *testing.T) {
	// Testcases.
	tests := []struct {
		name      string
		subName   string
		username  string
		wantError error
	}{
		{"Valid", "PHP", "Jonny", nil},
		{"Already Joined", "meirl", "Jonny", ErrSubAlreadyJoined},
		{"Invalid Username", "meirl", "Morrison", ErrUsernameNotExist},
		{"Invalid Subreddit", "TC130", "Jonny", ErrSubredditNotExist},
	}

	// Perform tests.
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Stub and driver.
			db, teardown := newTestDB(t)
			defer teardown()

			m := AccountModel{db}

			// When.
			err := m.JoinSubreddit(tc.subName, tc.username)

			// Want.
			if err != tc.wantError {
				t.Errorf("want %v; got %v", tc.wantError, err)
			}
		})
	}
}

func TestAccountLeaveSubreddit(t *testing.T) {
	// Testcases.
	tests := []struct {
		name      string
		subName   string
		username  string
		wantError error
	}{
		{"Valid", "dankmeme", "Jonny", nil},
		{"Not Joined", "meirl", "Albarn", ErrSubNotJoined},
	}

	// Perform tests.
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Stub and driver.
			db, teardown := newTestDB(t)
			defer teardown()

			m := AccountModel{db}

			// When.
			err := m.LeaveSubreddit(tc.subName, tc.username)

			// Want.
			if err != tc.wantError {
				t.Errorf("want %v; got %v", tc.wantError, err)
			}
		})
	}
}

func TestAccountGet(t *testing.T) {
	// Testcases.
	tests := []struct {
		name         string
		username     string
		wantUserInfo UserInfo
		wantError    error
	}{
		{"Valid User 1", "Jonny", UserInfo{"Jonny", 2, "Pop is dead, long live pop.", time.Date(1971, 11, 5, 10, 33, 52, 0, time.UTC)}, nil},
		{"Valid User 2", "Morrissey", UserInfo{"Morrissey", 2, "Oh mother, I can feel, the soil falling.", time.Date(1959, 5, 22, 10, 23, 54, 0, time.UTC)}, nil},
		{"Invalid Username", "Bowie", UserInfo{}, ErrUsernameNotExist},
	}

	// Perform tests.
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Stub and driver.
			db, teardown := newTestDB(t)
			defer teardown()

			m := AccountModel{db}

			// When.
			userInfo, err := m.Get(tc.username)

			// Want.
			if !reflect.DeepEqual(userInfo, tc.wantUserInfo) {
				t.Errorf("want:\n%v\ngot:\n%v", tc.wantUserInfo, userInfo)
			}

			if err != tc.wantError {
				t.Errorf("want %v; got %v", tc.wantError, err)
			}
		})
	}
}

// Combination test.
func TestRegisterAndAuthenticate(t *testing.T) {
	username, password, email := "thomyorke", "loveoasis", "muse@circlejerk.com"

	// Stub and driver.
	db, teardown := newTestDB(t)
	defer teardown()

	m := AccountModel{db}

	// When.
	err := m.Insert(username, email, password)
	if err != nil {
		t.Fatal(err)
	}

	// Want.
	err = m.Authenticate(username, password)
	if err != nil {
		t.Errorf("want %v; got %v", nil, err)
	}
}
