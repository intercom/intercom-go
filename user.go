package intercom

import "fmt"

type UserService struct {
	User
	Repository UserRepository
}

type UserList struct {
	Pages PageParams
	Users []User
}

type User struct {
	ID               string                 `json:"id,omitempty"`
	Email            string                 `json:"email,omitempty"`
	UserID           string                 `json:"user_id,omitempty"`
	SignedUpAt       int32                  `json:"signed_up_at,omitempty"`
	Name             string                 `json:"name,omitempty"`
	CustomAttributes map[string]interface{} `json:"custom_attributes,omitempty"`
}

type UserIdentifiers struct {
	ID     string `url:"-"`
	UserID string `url:"user_id,omitempty"`
	Email  string `url:"email,omitempty"`
}

func (u UserService) FindByID(id string) (User, error) {
	return u.findWithIdentifiers(UserIdentifiers{ID: id})
}

func (u UserService) FindByUserID(userID string) (User, error) {
	return u.findWithIdentifiers(UserIdentifiers{UserID: userID})
}

func (u UserService) FindByEmail(email string) (User, error) {
	return u.findWithIdentifiers(UserIdentifiers{Email: email})
}

func (u UserService) findWithIdentifiers(identifiers UserIdentifiers) (User, error) {
	var err error
	u.User, err = u.Repository.find(identifiers)
	return u.User, err
}

func (u UserService) List(params PageParams) (UserList, error) {
	return u.Repository.list(params)
}

func (u UserService) Save(user *User) (User, error) {
	return u.Repository.save(user)
}

func (u User) String() string {
	return fmt.Sprintf("[intercom] user { id: %s name: %s, user_id: %s, email: %s }", u.ID, u.Name, u.UserID, u.Email)
}

func (u User) authorType() string {
	return "user"
}
