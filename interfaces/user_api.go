package interfaces

import (
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
	api.httpClient.Get("")
	return domain.User{ID: params.ID, Email: params.Email, UserID: params.UserID}, nil
}

func (api UserAPI) List(params usecases.PageParams) ([]domain.User, error) {
	api.httpClient.Get("")
	return []domain.User{domain.User{}}, nil
}

func (api UserAPI) Save(user domain.User) error {
	api.httpClient.Post("", user)
	return nil
}
