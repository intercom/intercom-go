package interfaces

import (
	"encoding/json"
	"fmt"

	"github.com/intercom/intercom-go/domain"
	"github.com/intercom/intercom-go/usecases"
)

type UserAPI struct {
	httpClient HTTPClient
}

func NewUserAPI(httpClient HTTPClient) UserAPI {
	return UserAPI{httpClient: httpClient}
}

func (api UserAPI) Find(params usecases.UserIdentifiers) (domain.User, error) {
	user := domain.User{}
	data, err := api.getClientForFind(params)
	if err != nil {
		return user, err
	}
	err = json.Unmarshal(data, &user)
	return user, err
}

func (api UserAPI) getClientForFind(params usecases.UserIdentifiers) ([]byte, error) {
	switch {
	case params.ID != "":
		return api.httpClient.Get(fmt.Sprintf("/users/%s", params.ID), nil)
	case params.UserID != "", params.Email != "":
		return api.httpClient.Get("/users", params)
	}
	return nil, nil
}

func (api UserAPI) List(params usecases.PageParams) (usecases.UserList, error) {
	user_list := usecases.UserList{}
	data, err := api.httpClient.Get("/users", params)
	if err != nil {
		return user_list, err
	}
	err = json.Unmarshal(data, &user_list)
	return user_list, err
}

func (api UserAPI) Save(user domain.User) error {
	api.httpClient.Post("", user)
	return nil
}
