package intercom

import (
	"io/ioutil"
	"testing"
)

func TestAdminAPIList(t *testing.T) {
	http := TestAdminHTTPClient{fixtureFilename: "fixtures/admins.json", expectedURI: "/admins", t: t}
	api := AdminAPI{httpClient: &http}
	adminList, _ := api.list()
	if adminList.Admins[0].ID != "1" {
		t.Errorf("ID was %s, expected 1", adminList.Admins[0].ID)
	}
	if adminList.Admins[0].Avatar.ImageURL != "https://intercom.io/testA.png" {
		t.Errorf("Avatar was %s, expected https://intercom.io/testA.png", adminList.Admins[0].Avatar)
	}
}

type TestAdminHTTPClient struct {
	TestHTTPClient
	t               *testing.T
	fixtureFilename string
	expectedURI     string
}

func (t TestAdminHTTPClient) Get(uri string, queryParams interface{}) ([]byte, error) {
	if t.expectedURI != uri {
		t.t.Errorf("URI was %s, expected %s", uri, t.expectedURI)
	}
	return ioutil.ReadFile(t.fixtureFilename)
}
