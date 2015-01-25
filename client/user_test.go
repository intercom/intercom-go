package client

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

func TestUserFindByEmail(t *testing.T) {
	user, _ := User{Repository: TestUserAPI{t: t}}.FindByEmail("jamie@example.io")
	if user.Email != "jamie@example.io" {
		t.Errorf("User not found")
	}
}

func TestUserFindByUserID(t *testing.T) {
	user, _ := User{Repository: TestUserAPI{t: t}}.FindByUserID("134d")
	if user.UserID != "134d" {
		t.Errorf("User not found")
	}
}

func TestUserList(t *testing.T) {
	user_list, _ := User{Repository: TestUserAPI{t: t}}.List(PageParams{})
	users := user_list.Users
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

func (t TestUserAPI) List(params PageParams) (UserList, error) {
	return UserList{Users: []domain.User{domain.User{ID: "46adad3f09126dca", Email: "jamie@example.io", UserID: "aa123"}}}, nil
}

func (t TestUserAPI) Save(user domain.User) error {
	if user.ID != "46adad3f09126dca" {
		t.t.Errorf("Nope %s", user.ID)
	}
	return nil
}
