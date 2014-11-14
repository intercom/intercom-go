package intercom

import "github.com/franela/goreq"

type User struct {
	*Resource
	UserId     string                 `json:"user_id,omitempty"`
	CustomData map[string]interface{} `json:"custom_data,omitempty"`
}

type UserParams struct {
	UserId     string
	CustomData CustomData
}

func (u User) New(params *UserParams) (*goreq.Response, error) {
	user := User{
		UserId:     params.UserId,
		CustomData: params.CustomData,
	}
	res, err := u.client.Post("/v1/users", user)
	return res, err
}
