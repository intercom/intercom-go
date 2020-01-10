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
	if user.Phone != "+12345678910" {
		t.Errorf("Phone was %s, expected +12345678910", user.Phone)
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
	userList, _ := api.list(userListParams{})
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
	userList, _ := api.list(userListParams{PageParams: PageParams{Page: 2}})
	pages := userList.Pages
	if pages.Page != 2 {
		t.Errorf("Page was %d, expected 2", pages.Page)
	}
}

func TestUserAPIListWithSegment(t *testing.T) {
	http := TestUserHTTPClient{fixtureFilename: "fixtures/users.json", expectedURI: "/users", t: t}
	api := UserAPI{httpClient: &http}
	api.list(userListParams{SegmentID: "abc123"})
	if ulParams, ok := http.lastQueryParams.(userListParams); !ok || ulParams.SegmentID != "abc123" {
		t.Errorf("SegmentID expected to be abc123, but was %s", ulParams.SegmentID)
	}
}

func TestUserAPIListWithTag(t *testing.T) {
	http := TestUserHTTPClient{fixtureFilename: "fixtures/users.json", expectedURI: "/users", t: t}
	api := UserAPI{httpClient: &http}
	api.list(userListParams{TagID: "123"})
	if ulParams, ok := http.lastQueryParams.(userListParams); !ok || ulParams.TagID != "123" {
		t.Errorf("SegmentID expected to be 123, but was %s", ulParams.TagID)
	}
}

func TestUserAPISave(t *testing.T) {
	http := TestUserHTTPClient{t: t, expectedURI: "/users"}
	api := UserAPI{httpClient: &http}
	companyList := CompanyList{
		Companies: []Company{
			Company{ID: "5"},
		},
	}
	user := User{UserID: "27", Companies: &companyList}
	api.save(&user)
}

func TestUserAPIDelete(t *testing.T) {
	http := TestUserHTTPClient{t: t, expectedURI: "/users/1234"}
	api := UserAPI{httpClient: &http}
	api.delete("1234")
}

type TestUserHTTPClient struct {
	TestHTTPClient
	t               *testing.T
	fixtureFilename string
	expectedURI     string
	lastQueryParams interface{}
}

func (t *TestUserHTTPClient) Get(uri string, queryParams interface{}) ([]byte, error) {
	if t.expectedURI != uri {
		t.t.Errorf("URI was %s, expected %s", uri, t.expectedURI)
	}
	t.lastQueryParams = queryParams
	return ioutil.ReadFile(t.fixtureFilename)
}

func (t *TestUserHTTPClient) Post(uri string, body interface{}) ([]byte, error) {
	if t.expectedURI != uri {
		t.t.Errorf("Wrong endpoint called")
	}
	return ioutil.ReadFile(t.fixtureFilename)
}

func (t *TestUserHTTPClient) Delete(uri string, queryParams interface{}) ([]byte, error) {
	if t.expectedURI != uri {
		t.t.Errorf("URI was %s, expected %s", uri, t.expectedURI)
	}
	return ioutil.ReadFile(t.fixtureFilename)
}
