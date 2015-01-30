package interfaces

import (
	"io/ioutil"
	"testing"

	"github.com/intercom/intercom-go/client"
)

func TestFind(t *testing.T) {
	http := TestUserHTTPClient{fixtureFilename: "fixtures/user.json", expectedURI: "/users/54c42e7ea7a765fa7", t: t}
	api := UserAPI{httpClient: &http}
	user, _ := api.Find(client.UserIdentifiers{ID: "54c42e7ea7a765fa7"})
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

func TestFindByEmail(t *testing.T) {
	http := TestUserHTTPClient{fixtureFilename: "fixtures/user.json", expectedURI: "/users", t: t}
	api := UserAPI{httpClient: &http}
	user, _ := api.Find(client.UserIdentifiers{Email: "myuser@example.io"})
	if user.Email != "myuser@example.io" {
		t.Errorf("Email was %s, expected myuser@example.io", user.Email)
	}
}

func TestListDefault(t *testing.T) {
	http := TestUserHTTPClient{fixtureFilename: "fixtures/user_list.json", expectedURI: "/users", t: t}
	api := UserAPI{httpClient: &http}
	user_list, _ := api.List(client.PageParams{})
	users := user_list.Users
	if users[0].ID != "54c42e7ea7a765fa7" {
		t.Errorf("ID was %s, expected 54c42e7ea7a765fa7", users[0].ID)
	}
	pages := user_list.Pages
	if pages.Page != 1 {
		t.Errorf("Page was %d, expected 1", pages.Page)
	}
}

func TestListWithPageNumber(t *testing.T) {
	http := TestUserHTTPClient{fixtureFilename: "fixtures/user_list_page_2.json", expectedURI: "/users", t: t}
	api := UserAPI{httpClient: &http}
	user_list, _ := api.List(client.PageParams{Page: 2})
	pages := user_list.Pages
	if pages.Page != 2 {
		t.Errorf("Page was %d, expected 2", pages.Page)
	}
}

type TestUserHTTPClient struct {
	TestHTTPClient
	t               *testing.T
	fixtureFilename string
	expectedURI     string
}

func (h TestUserHTTPClient) Get(uri string, queryParams interface{}) ([]byte, error) {
	if h.expectedURI != uri {
		h.t.Errorf("URI was %s, expected %s", uri, h.expectedURI)
	}
	return ioutil.ReadFile(h.fixtureFilename)
}

func (h TestUserHTTPClient) Post(uri string, body interface{}) ([]byte, error) {
	return nil, nil
}
