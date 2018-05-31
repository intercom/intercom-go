package intercom

import (
	"testing"
)

func TestUserFindByID(t *testing.T) {
	user, _ := (&UserService{Repository: TestUserAPI{t: t}}).FindByID("46adad3f09126dca")
	if user.ID != "46adad3f09126dca" {
		t.Errorf("User not found")
	}
}

func TestUserFindByEmail(t *testing.T) {
	user, _ := (&UserService{Repository: TestUserAPI{t: t}}).FindByEmail("jamie@example.io")
	if user.Email != "jamie@example.io" {
		t.Errorf("User not found")
	}
}

func TestUserFindByUserID(t *testing.T) {
	user, _ := (&UserService{Repository: TestUserAPI{t: t}}).FindByUserID("134d")
	if user.UserID != "134d" {
		t.Errorf("User not found")
	}
}

func TestUserList(t *testing.T) {
	userList, _ := (&UserService{Repository: TestUserAPI{t: t}}).List(PageParams{})
	users := userList.Users
	if users[0].ID != "46adad3f09126dca" {
		t.Errorf("User not listed")
	}
}

func TestUserSave(t *testing.T) {
	userService := UserService{Repository: TestUserAPI{t: t}}
	user := User{ID: "46adad3f09126dca", CustomAttributes: map[string]interface{}{"is_cool": true}}
	userService.Save(&user)
}

func TestUserDelete(t *testing.T) {
	(&UserService{Repository: TestUserAPI{t: t}}).Delete("46adad3f09126dca")
}

func TestUserMessageAddress(t *testing.T) {
	contact := User{ID: "46adad3f09126dca", UserID: "aaaa", Email: "some@email.com"}
	address := contact.MessageAddress()
	if address.ID != "46adad3f09126dca" {
		t.Errorf("User address had wrong ID")
	}
	if address.Type != "user" {
		t.Errorf("User address was not of type user, was %s", address.Type)
	}
	if address.Email != "some@email.com" {
		t.Errorf("User address had wrong Email")
	}
	if address.UserID != "aaaa" {
		t.Errorf("User address had wrong UserID")
	}
}

type TestUserAPI struct {
	t *testing.T
}

func (t TestUserAPI) find(params UserIdentifiers) (User, error) {
	return User{ID: params.ID, Email: params.Email, UserID: params.UserID}, nil
}

func (t TestUserAPI) list(params userListParams) (UserList, error) {
	return UserList{Users: []User{User{ID: "46adad3f09126dca", Email: "jamie@example.io", UserID: "aa123"}}}, nil
}

func (t TestUserAPI) scroll(scrollParam string) (UserList, error) {
	return UserList{Users: []User{User{ID: "46adad3f09126dca", Email: "jamie@example.io", UserID: "aa123"}}}, nil
}

func (t TestUserAPI) save(user *User) (User, error) {
	if user.ID != "46adad3f09126dca" {
		t.t.Errorf("User ID was %s, expected 46adad3f09126dca", user.ID)
	}
	expectedCAs := map[string]interface{}{"is_cool": true}
	if user.CustomAttributes["is_cool"] != expectedCAs["is_cool"] {
		t.t.Errorf("Custom attributes was %v, expected %v", user.CustomAttributes, expectedCAs)
	}
	return User{}, nil
}

func (t TestUserAPI) delete(id string) (User, error) {
	if id != "46adad3f09126dca" {
		t.t.Errorf("id was %s, expected 46adad3f09126dca", id)
	}
	return User{}, nil
}
