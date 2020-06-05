package mock

import "github.com/YuChaoGithub/YARC/backend/app/models"

// AccountModel is a mock account model.
type AccountModel struct{}

// Insert only returns ErrEmailDup if the email is "mobile@blizzard.com.hk".
func (m *AccountModel) Insert(username, email, password string) error {
	if email == "mobile@blizzard.com.hk" {
		return models.ErrEmailDup
	}
	return nil
}

// Authenticate only returns nil when username is "mediocre_nothingness" and password "pasword".
func (m *AccountModel) Authenticate(username, password string) error {
	if username == "mediocre_nothingness" {
		if password == "password" {
			return nil
		}
		return models.ErrAccountInfoMismatch
	}

	return models.ErrUsernameNotExist
}

// ModifyBio TODO.
func (m *AccountModel) ModifyBio(username, newBio string) error {
	return nil
}

// SaveArticle TODO.
func (m *AccountModel) SaveArticle(articleID, username string) error {
	return nil
}

// JoinSubreddit TODO.
func (m *AccountModel) JoinSubreddit(subName, username string) error {
	return nil
}

// Get TODO.
func (m *AccountModel) Get(username string) (models.UserInfo, error) {
	return models.UserInfo{Username: "user", Karma: 0, Bio: "Bio"}, nil
}
