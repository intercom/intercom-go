package interfaces

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/intercom/intercom-go/client"
	"github.com/intercom/intercom-go/domain"
)

type UserAPI struct {
	httpClient HTTPClient
}

func NewUserAPI(httpClient HTTPClient) UserAPI {
	return UserAPI{httpClient: httpClient}
}

func (api UserAPI) Find(params client.UserIdentifiers) (domain.User, error) {
	user := domain.User{}
	data, err := api.getClientForFind(params)
	if err != nil {
		return user, err
	}
	err = json.Unmarshal(data, &user)
	return user, err
}

func (api UserAPI) getClientForFind(params client.UserIdentifiers) ([]byte, error) {
	switch {
	case params.ID != "":
		return api.httpClient.Get(fmt.Sprintf("/users/%s", params.ID), nil)
	case params.UserID != "", params.Email != "":
		return api.httpClient.Get("/users", params)
	}
	return nil, errors.New("Missing User Identifier")
}

func (api UserAPI) List(params client.PageParams) (client.UserList, error) {
	user_list := client.UserList{}
	data, err := api.httpClient.Get("/users", params)
	if err != nil {
		return user_list, err
	}
	err = json.Unmarshal(data, &user_list)
	return user_list, err
}

func (api UserAPI) Save(user domain.User) (domain.User, error) {
	saved_user := domain.User{}
	data, err := api.httpClient.Post("/users", user)
	if err != nil {
		return user, err
	}
	err = json.Unmarshal(data, &saved_user)
	return saved_user, err
}
