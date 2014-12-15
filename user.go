package intercom

type User struct {
	*Resource
	*userIdentifiers              // inline identifiers
	RemoteCreatedAt  int32        `json:"remote_created_at,omitempty"`
	Name             string       `json:"name,omitempty"`
	CustomData       AttributeMap `json:"custom_data,omitempty"`
}

type UserList struct {
	Pages Pages  `json:"pages,omitempty"`
	Users []User `json:"users"`
}

type UserParams struct {
	Id              string
	Email           string
	UserId          string
	RemoteCreatedAt int32
	Name            string
	CustomData      AttributeMap
}

type userIdentifiers struct {
	Id     string `json:"id,omitempty" url:"id,omitempty"`
	UserId string `json:"user_id,omitempty" url:"user_id,omitempty"`
	Email  string `json:"email,omitempty" url:"email,omitempty"`
}

type queryUser struct {
	Id     string
	UserId string `url:"user_id"`
}

func (u User) Find(params *UserParams) (*User, error) {
	userIdentifiers := userIdentifiers{
		Id:     params.Id,
		UserId: params.UserId,
		Email:  params.Email,
	}
	if responseBody, err := u.client.Get("/users", userIdentifiers); err != nil {
		return nil, err
	} else {
		return u.unmarshalIntoNewUser(responseBody)
	}
}

func (u User) List(params *PageParams) (*UserList, error) {
	if responseBody, err := u.client.Get("/users", params); err != nil {
		return nil, err
	} else {
		userList := UserList{}
		return &userList, u.Unmarshal(&userList, responseBody.([]byte))
	}
}

func (u User) New(params *UserParams) (*User, error) {
	user := User{
		userIdentifiers: &userIdentifiers{
			UserId: params.UserId,
		},
		RemoteCreatedAt: params.RemoteCreatedAt,
		Name:            params.Name,
		CustomData:      params.CustomData,
	}
	if responseBody, err := u.client.Post("/users", user); err != nil {
		return nil, err
	} else {
		return u.unmarshalIntoNewUser(responseBody)
	}
}

func (u User) unmarshalIntoNewUser(responseBody interface{}) (*User, error) {
	user := User{}
	return &user, u.Unmarshal(&user, responseBody.([]byte))
}
