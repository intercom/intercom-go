package usecases

import "github.com/intercom/intercom-go/domain"

type UserIdentifiers struct {
	ID     string
	UserID string
	Email  string
}

type UserRepository interface {
	Find(UserIdentifiers) (domain.User, error)
	List(PageParams) ([]domain.User, error)
	Save(domain.User) error
}

type User struct {
	domain.User
	Repository UserRepository
}

func (u User) FindByID(id string) (User, error) {
	return u.FindWithIdentifiers(UserIdentifiers{ID: id})
}

func (u User) FindByUserID(userID string) (User, error) {
	return u.FindWithIdentifiers(UserIdentifiers{UserID: userID})
}

func (u User) FindByEmail(email string) (User, error) {
	return u.FindWithIdentifiers(UserIdentifiers{Email: email})
}

func (u User) FindWithIdentifiers(identifiers UserIdentifiers) (User, error) {
	var err error
	u.User, err = u.Repository.Find(identifiers)
	return u, err
}

func (u User) List(repository UserRepository, params PageParams) ([]domain.User, error) {
	return repository.List(params)
}

func (u User) Save() error {
	return u.Repository.Save(u.User)
}
