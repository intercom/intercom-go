package intercom

import (
	"io/ioutil"
	"testing"
)

func TestConversationAPIFind(t *testing.T) {
	http := TestUserHTTPClient{fixtureFilename: "fixtures/conversation.json", expectedURI: "/conversations/171", t: t}
	api := ConversationAPI{httpClient: &http}
	convo, _ := api.find("171")
	if convo.ID != "171" {
		t.Errorf("ID was %s, expected 171", convo.ID)
	}
}

type TestConversationHTTPClient struct {
	TestHTTPClient
	t               *testing.T
	fixtureFilename string
	expectedURI     string
}

func (t TestConversationHTTPClient) Get(uri string, queryParams interface{}) ([]byte, error) {
	if t.expectedURI != uri {
		t.t.Errorf("URI was %s, expected %s", uri, t.expectedURI)
	}
	return ioutil.ReadFile(t.fixtureFilename)
}
