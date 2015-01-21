package intercom

import (
	"github.com/intercom/intercom-go/interfaces"
	"github.com/intercom/intercom-go/usecases"
)

type SimpleClient struct {
	User usecases.User
}

func (c SimpleClient) NewUser() *usecases.User {
	return &usecases.User{UserRepository: interfaces.NewUserAPI(interfaces.NewIntercomHttpClient("tx2p130c", "28d8ae4c0868c5f7deb75eda9a8a7c6cc9f435b3"))}
}

func NewSimpleClient() SimpleClient {
	httpClient := interfaces.NewIntercomHttpClient("tx2p130c", "28d8ae4c0868c5f7deb75eda9a8a7c6cc9f435b3")
	userAPI := interfaces.NewUserAPI(httpClient)
	return SimpleClient{User: usecases.User{UserRepository: userAPI}}
}
