package intercom

import (
	"io/ioutil"
	"testing"
)

func TestUserAPIFind(t *testing.T) {
	http := TestUserHTTPClient{fixtureFilename: "fixtures/user.json", expectedURI: "/users/54c42e7ea7a765fa7", t: t}
	api := UserAPI{httpClient: &http}
	user, err := api.find(UserIdentifiers{ID: "54c42e7ea7a765fa7"})
	if err != nil {
		t.Errorf("Error parsing fixture %s", err)
	}
	if user.ID != "54c42e7ea7a765fa7" {
		t.Errorf("ID was %s, expected 54c42e7ea7a765fa7", user.ID)
	}
	if user.UserID != "123" {
		t.Errorf("UserID was %s, expected 123", user.UserID)
	}
	if user.SignedUpAt != 1422143117 {
		t.Errorf("SignedUpAt was %d, expected %d", user.SignedUpAt, 1422143117)
	}
	if user.CustomAttributes["is_awesome"] != true {
		t.Errorf("CustomAttributes was %v, expected %v", user.CustomAttributes, map[string]interface{}{"is_awesome": true})
	}
}

func TestUserAPIFindByEmail(t *testing.T) {
	http := TestUserHTTPClient{fixtureFilename: "fixtures/user.json", expectedURI: "/users", t: t}
	api := UserAPI{httpClient: &http}
	user, _ := api.find(UserIdentifiers{Email: "myuser@example.io"})
	if user.Email != "myuser@example.io" {
		t.Errorf("Email was %s, expected myuser@example.io", user.Email)
	}
}

func TestUserAPIListDefault(t *testing.T) {
	http := TestUserHTTPClient{fixtureFilename: "fixtures/users.json", expectedURI: "/users", t: t}
	api := UserAPI{httpClient: &http}
	userList, _ := api.list(PageParams{})
	users := userList.Users
	if users[0].ID != "54c42e7ea7a765fa7" {
		t.Errorf("ID was %s, expected 54c42e7ea7a765fa7", users[0].ID)
	}
	pages := userList.Pages
	if pages.Page != 1 {
		t.Errorf("Page was %d, expected 1", pages.Page)
	}
}

func TestUserAPIListWithPageNumber(t *testing.T) {
	http := TestUserHTTPClient{fixtureFilename: "fixtures/users_page_2.json", expectedURI: "/users", t: t}
	api := UserAPI{httpClient: &http}
	userList, _ := api.list(PageParams{Page: 2})
	pages := userList.Pages
	if pages.Page != 2 {
		t.Errorf("Page was %d, expected 2", pages.Page)
	}
}

func TestUserAPISave(t *testing.T) {
	http := TestUserHTTPClient{t: t, expectedURI: "/users"}
	api := UserAPI{httpClient: &http}
	user := User{UserID: "27"}
	api.save(&user)
}

type TestUserHTTPClient struct {
	TestHTTPClient
	t               *testing.T
	fixtureFilename string
	expectedURI     string
}

func (t TestUserHTTPClient) Get(uri string, queryParams interface{}) ([]byte, error) {
	if t.expectedURI != uri {
		t.t.Errorf("URI was %s, expected %s", uri, t.expectedURI)
	}
	return ioutil.ReadFile(t.fixtureFilename)
}

func (t TestUserHTTPClient) Post(uri string, body interface{}) ([]byte, error) {
	if uri != "/users" {
		t.t.Errorf("Wrong endpoint called")
	}
	return nil, nil
}
