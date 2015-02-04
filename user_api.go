package intercom

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/intercom/intercom-go/interfaces"
)

type UserRepository interface {
	find(UserIdentifiers) (User, error)
	list(userListParams) (UserList, error)
	save(*User) (User, error)
	delete(id string) (User, error)
}

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
}

func (api UserAPI) find(params UserIdentifiers) (User, error) {
	user := User{}
	data, err := api.getClientForFind(params)
	if err != nil {
		return user, err
	}
	err = json.Unmarshal(data, &user)
	return user, err
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
		RemoteCreatedAt:        user.RemoteCreatedAt,
		LastRequestAt:          user.LastRequestAt,
		LastSeenIP:             user.LastSeenIP,
		UnsubscribedFromEmails: user.UnsubscribedFromEmails,
		Companies:              api.getCompaniesToSendFromUser(user),
		CustomAttributes:       user.CustomAttributes,
		NewSession:             user.NewSession,
	}

	savedUser := User{}
	data, err := api.httpClient.Post("/users", &requestUser)
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
	companies := user.Companies.Companies
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
