package usecases

import (
	"time"

	"github.com/intercom/intercom-go/domain"
)

type UserParams struct {
	Id         string
	UserId     string
	Email      string
	SignedUp   time.Time
	Name       string
	CustomData map[string]interface{}
	Unread     *bool
}

type UserRepository interface {
	Find(UserParams) (domain.User, error)
	List(PageParams) ([]domain.User, error)
	Save(domain.User) error
}

type User struct {
	domain.User
	UserRepository
}

func (u User) FindMe(ID string) (domain.User, error) {
	return u.UserRepository.Find(UserParams{Id: ID})
}

func (u User) Find(repository UserRepository, params UserParams) (domain.User, error) {
	return repository.Find(params)
}

func (u User) List(repository UserRepository, params PageParams) ([]domain.User, error) {
	return repository.List(params)
}

func (u User) SaveMe() error {
	return u.UserRepository.Save(u.User)
}

func (u User) Save(repository UserRepository, params UserParams) (domain.User, error) {
	user := domain.User{
		Id:       params.Id,
		UserId:   params.UserId,
		Email:    params.Email,
		SignedUp: params.SignedUp,
		Name:     params.Name,
	}
	err := repository.Save(user)
	return user, err
}
