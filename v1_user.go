package intercom

import "github.com/franela/goreq"

type User struct {
	*Resource
	UserId     string       `json:"user_id,omitempty"`
	CustomData AttributeMap `json:"custom_data,omitempty"`
}

type UserParams struct {
	UserId     string
	CustomData AttributeMap
}

func (u User) New(params *UserParams) (*goreq.Response, error) {
	user := User{
		UserId:     params.UserId,
		CustomData: params.CustomData,
	}
	res, err := u.client.Post("/v1/users", user)
	return res, err
}
