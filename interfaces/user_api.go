package interfaces

import (
	"github.com/intercom/intercom-go/domain"
	"github.com/intercom/intercom-go/usecases"
)

type UserAPI struct {
	httpClient HttpClient
}

func NewUserAPI(httpClient HttpClient) UserAPI {
	return UserAPI{httpClient: httpClient}
}

func (api UserAPI) Find(params usecases.UserParams) (domain.User, error) {
	api.httpClient.Get("")
	return domain.User{Id: params.Id, Email: params.Email, UserId: params.UserId}, nil
}

func (api UserAPI) List(params usecases.PageParams) ([]domain.User, error) {
	api.httpClient.Get("")
	return []domain.User{domain.User{}}, nil
}

func (api UserAPI) Save(user domain.User) error {
	api.httpClient.Post("", user)
	return nil
}
