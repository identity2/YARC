package models

import "testing"

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
	// TODO
}

func TestAccountSaveArticle(t *testing.T) {
	// TODO
}

func TestJoinSubreddit(t *testing.T) {
	// TODO
}

func TestAccountGet(t *testing.T) {
	// TODO
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
