package mock

import "github.com/YuChaoGithub/YARC/backend/app/models"

// AccountModel is a mock account model.
type AccountModel struct{}

// Insert TODO.
func (m *AccountModel) Insert(username, email, password string) error {
	return nil
}

// Authenticate TODO.
func (m *AccountModel) Authenticate(username, password string) error {
	return nil
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
