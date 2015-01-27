package client

import "github.com/intercom/intercom-go/domain"

type UserIdentifiers struct {
	ID     string `url:"-"`
	UserID string `url:"user_id,omitempty"`
	Email  string `url:"email,omitempty"`
}

type UserRepository interface {
	Find(UserIdentifiers) (domain.User, error)
	List(PageParams) (UserList, error)
	Save(domain.User) (domain.User, error)
}

type UserList struct {
	Pages PageParams
	Users []domain.User
}

type User struct {
	domain.User
	Repository UserRepository
}

func (u User) FindByID(id string) (User, error) {
	return u.findWithIdentifiers(UserIdentifiers{ID: id})
}

func (u User) FindByUserID(userID string) (User, error) {
	return u.findWithIdentifiers(UserIdentifiers{UserID: userID})
}

func (u User) FindByEmail(email string) (User, error) {
	return u.findWithIdentifiers(UserIdentifiers{Email: email})
}

func (u User) findWithIdentifiers(identifiers UserIdentifiers) (User, error) {
	var err error
	u.User, err = u.Repository.Find(identifiers)
	return u, err
}

func (u User) List(params PageParams) (UserList, error) {
	return u.Repository.List(params)
}

func (u User) Save() (User, error) {
	var err error
	u.User, err = u.Repository.Save(u.User)
	return u, err
}
