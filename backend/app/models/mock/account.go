package mock

import (
	"time"

	"github.com/YuChaoGithub/YARC/backend/app/models"
)

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

// ModifyBio returns ErrBioInvalid if newBio is "Dirty Dan".
func (m *AccountModel) ModifyBio(username, newBio string) error {
	if newBio == "Dirty Dan" {
		return models.ErrBioInvalid
	}

	return nil
}

// SaveArticle returns ErrArticleAlreadySaved if articleID is "77777" and username is "PatrickStar".
func (m *AccountModel) SaveArticle(articleID, username string) error {
	if articleID == "77777" && username == "PatrickStar" {
		return models.ErrArticleAlreadySaved
	}

	return nil
}

// UnsaveArticle returns ErrArticleNotSaved if articleID is "5566" and username is "SpongeBob".
func (m *AccountModel) UnsaveArticle(articleID, username string) error {
	if articleID == "5566" && username == "SpongeBob" {
		return models.ErrArticleNotSaved
	}

	return nil
}

// GetJoinState returns true only when username is "Kurt" and subName is "Nirvana".
func (m *AccountModel) GetJoinState(username, subName string) bool {
	if username == "Kurt" && subName == "Nirvana" {
		return true
	}
	return false
}

// JoinSubreddit returns ErrSubAlreadyJoined when subName is "gamedev" and username is "nirvana".
func (m *AccountModel) JoinSubreddit(subName, username string) error {
	if subName == "gamedev" && username == "nirvana" {
		return models.ErrSubAlreadyJoined
	}

	return nil
}

// LeaveSubreddit returns ErrSubNotJoin when subName is "CSMajor" and username is "WarGod69".
func (m *AccountModel) LeaveSubreddit(subName, username string) error {
	if subName == "CSMajor" && username == "WarGod69" {
		return models.ErrSubNotJoined
	}

	return nil
}

// Get only has the account Jonny.
func (m *AccountModel) Get(username string) (models.UserInfo, error) {
	if username == "Jonny" {
		return models.UserInfo{
			Username: "Jonny",
			Karma:    -1,
			Bio:      "Who's in the bunker?",
			JoinTime: time.Date(1999, 9, 5, 10, 33, 52, 0, time.UTC),
		}, nil
	}
	return models.UserInfo{}, models.ErrUsernameNotExist
}
