package usecases

import (
	"testing"

	"github.com/intercom/intercom-go/domain"
)

func TestUserFindByID(t *testing.T) {
	user, _ := User{Repository: TestUserAPI{t: t}}.FindByID("46adad3f09126dca")
	if user.ID != "46adad3f09126dca" {
		t.Errorf("User not found")
	}
}

func TestUserList(t *testing.T) {
	users, _ := User{}.List(TestUserAPI{}, PageParams{})
	if users[0].ID != "46adad3f09126dca" {
		t.Errorf("User not listed")
	}
}

func TestUserSave(t *testing.T) {
	user := User{Repository: TestUserAPI{t: t}}
	user.ID = "46adad3f09126dca"
	user.Save()
}

type TestUserAPI struct {
	t *testing.T
}

func (t TestUserAPI) Find(params UserIdentifiers) (domain.User, error) {
	return domain.User{ID: params.ID, Email: params.Email, UserID: params.UserID}, nil
}

func (t TestUserAPI) List(params PageParams) ([]domain.User, error) {
	return []domain.User{domain.User{ID: "46adad3f09126dca", Email: "jamie@intercom.io", UserID: "aa123"}}, nil
}

func (t TestUserAPI) Save(user domain.User) error {
	if user.ID != "46adad3f09126dca" {
		t.t.Errorf("Nope %s", user.ID)
	}
	return nil
}
