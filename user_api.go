package intercom

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/intercom/intercom-go/interfaces"
)

// UserRepository defines the interface for working with Users through the API.
type UserRepository interface {
	find(UserIdentifiers) (User, error)
	list(userListParams) (UserList, error)
	scroll(scrollParam string) (UserScroll, error)
	save(*User) (User, error)
	delete(id string) (User, error)
}

// UserAPI implements UserRepository
type UserAPI struct {
	httpClient interfaces.HTTPClient
}

type requestUser struct {
	ID                     string                 `json:"id,omitempty"`
	Email                  string                 `json:"email,omitempty"`
	UserID                 string                 `json:"user_id,omitempty"`
	Name                   string                 `json:"name,omitempty"`
	SignedUpAt             int32                  `json:"signed_up_at,omitempty"`
	RemoteCreatedAt        int32                  `json:"remote_created_at,omitempty"`
	LastRequestAt          int32                  `json:"last_request_at,omitempty"`
	LastSeenIP             string                 `json:"last_seen_ip,omitempty"`
	UnsubscribedFromEmails *bool                  `json:"unsubscribed_from_emails,omitempty"`
	Companies              []UserCompany          `json:"companies,omitempty"`
	CustomAttributes       map[string]interface{} `json:"custom_attributes,omitempty"`
	UpdateLastRequestAt    *bool                  `json:"update_last_request_at,omitempty"`
	NewSession             *bool                  `json:"new_session,omitempty"`
	LastSeenUserAgent      string                 `json:"last_seen_user_agent,omitempty"`
}

func (api UserAPI) find(params UserIdentifiers) (User, error) {
	return unmarshalToUser(api.getClientForFind(params))
}

func (api UserAPI) getClientForFind(params UserIdentifiers) ([]byte, error) {
	switch {
	case params.ID != "":
		return api.httpClient.Get(fmt.Sprintf("/users/%s", params.ID), nil)
	case params.UserID != "", params.Email != "":
		return api.httpClient.Get("/users", params)
	}
	return nil, errors.New("Missing User Identifier")
}

func (api UserAPI) list(params userListParams) (UserList, error) {
	userList := UserList{}
	data, err := api.httpClient.Get("/users", params)
	if err != nil {
		return userList, err
	}
	err = json.Unmarshal(data, &userList)
	return userList, err
}

func (api UserAPI) scroll(scrollParam string) (UserScroll, error) {
	userScroll := UserScroll{}

	data, err := api.httpClient.Get("/users/scroll", map[string]string{
		"scroll_param": scrollParam,
	})
	if err != nil {
		return userScroll, err
	}

	if err = json.Unmarshal(data, &userScroll); err != nil {
		return userScroll, err
	}

	return userScroll, err
}

// A Company the User belongs to
// used to update Companies on a User.
type UserCompany struct {
	ID     string `json:"id,omitempty"`
	Name   string `json:"name,omitempty"`
	Remove *bool  `json:"remove,omitempty"`
}

func (api UserAPI) save(user *User) (User, error) {
	requestUser := requestUser{
		ID:                     user.ID,
		Email:                  user.Email,
		UserID:                 user.UserID,
		Name:                   user.Name,
		SignedUpAt:             user.SignedUpAt,
		RemoteCreatedAt:        user.RemoteCreatedAt,
		LastRequestAt:          user.LastRequestAt,
		LastSeenIP:             user.LastSeenIP,
		UnsubscribedFromEmails: user.UnsubscribedFromEmails,
		Companies:              api.getCompaniesToSendFromUser(user),
		CustomAttributes:       user.CustomAttributes,
		UpdateLastRequestAt:    user.UpdateLastRequestAt,
		NewSession:             user.NewSession,
		LastSeenUserAgent:      user.LastSeenUserAgent,
	}
	return unmarshalToUser(api.httpClient.Post("/users", &requestUser))
}

func unmarshalToUser(data []byte, err error) (User, error) {
	savedUser := User{}
	if err != nil {
		return savedUser, err
	}
	err = json.Unmarshal(data, &savedUser)
	return savedUser, err
}

func (api UserAPI) getCompaniesToSendFromUser(user *User) []UserCompany {
	if user.Companies == nil {
		return []UserCompany{}
	}
	return makeUserCompaniesFromCompanies(user.Companies.Companies)
}

func makeUserCompaniesFromCompanies(companies []Company) []UserCompany {
	userCompanies := make([]UserCompany, len(companies))
	for i := 0; i < len(companies); i++ {
		userCompanies[i] = UserCompany{
			ID:     companies[i].ID,
			Name:   companies[i].Name,
			Remove: companies[i].Remove,
		}
	}
	return userCompanies
}

func (api UserAPI) delete(id string) (User, error) {
	user := User{}
	data, err := api.httpClient.Delete(fmt.Sprintf("/users/%s", id), nil)
	if err != nil {
		return user, err
	}
	err = json.Unmarshal(data, &user)
	return user, err
}
