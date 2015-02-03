package intercom

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/intercom/intercom-go/interfaces"
)

type UserRepository interface {
	find(UserIdentifiers) (User, error)
	list(PageParams) (UserList, error)
	save(*User) (User, error)
}

type UserAPI struct {
	httpClient interfaces.HTTPClient
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

func (api UserAPI) list(params PageParams) (UserList, error) {
	userList := UserList{}
	data, err := api.httpClient.Get("/users", params)
	if err != nil {
		return userList, err
	}
	err = json.Unmarshal(data, &userList)
	return userList, err
}

func (api UserAPI) save(user *User) (User, error) {
	savedUser := User{}
	data, err := api.httpClient.Post("/users", user)
	if err != nil {
		return savedUser, err
	}
	err = json.Unmarshal(data, &savedUser)
	return savedUser, err
}
