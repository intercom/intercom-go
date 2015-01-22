package intercom

import (
	"github.com/intercom/intercom-go/interfaces"
	"github.com/intercom/intercom-go/usecases"
)

type Client struct {
	User           usecases.User
	userRepository usecases.UserRepository
}

func NewClient(appID, apiKey string) Client {
	httpClient := interfaces.NewIntercomHTTPClient(appID, apiKey)
	userAPI := interfaces.NewUserAPI(httpClient)
	return Client{User: usecases.User{Repository: userAPI}}
}
