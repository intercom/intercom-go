package usecases

import (
	"testing"

	"github.com/intercom/intercom-go/domain"
)

func TestFindUser(t *testing.T) {
	user, _ := User{}.Find(TestUserAPI{}, UserParams{Id: "46adad3f09126dca"})
	if user.Id != "46adad3f09126dca" {
		t.Errorf("User not found")
	}
}

func TestListUser(t *testing.T) {
	users, _ := User{}.List(TestUserAPI{}, PageParams{})
	if users[0].Id != "46adad3f09126dca" {
		t.Errorf("User not listed")
	}
}

//
// func TestSaveUser(t *testing.T) {
// 	user, _ := User{}.Save(TestUserAPI{}, UserParams{Id: "46adad3f09126dca"})
// 	if "46adad3f09126dca" != user.Id {
// 		t.Errorf("User not saved")
// 	}
// }

func TestUserSaveMe(t *testing.T) {
	user := User{UserRepository: TestUserAPI{t: t}}
	user.Id = "46adad3f09126dca"
	user.SaveMe()
}

type TestUserAPI struct {
	t *testing.T
}

func (t TestUserAPI) Find(params UserParams) (domain.User, error) {
	return domain.User{Id: params.Id, Email: params.Email, UserId: params.UserId}, nil
}

func (t TestUserAPI) List(params PageParams) ([]domain.User, error) {
	return []domain.User{domain.User{Id: "46adad3f09126dca", Email: "jamie@intercom.io", UserId: "aa123"}}, nil
}

func (t TestUserAPI) Save(user domain.User) error {
	if user.Id != "46adad3f09126dca" {
		t.t.Errorf("Nope %s", user.Id)
	}
	return nil
}
