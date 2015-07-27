package intercom

import (
	"io/ioutil"
	"testing"
)

func TestConversationListAll(t *testing.T) {
	http := TestConversationHTTPClient{t: t, expectedURI: "/conversations", fixtureFilename: "fixtures/conversations.json"}
	api := ConversationAPI{httpClient: &http}
	convos, _ := api.list(conversationListParams{})
	if convos.Conversations[0].ID != "147" {
		t.Errorf("Conversation not retrieved")
	}
	if convos.Conversations[0].User.ID != "536e564f316c83104c000020" {
		t.Errorf("Conversation user not retrieved")
	}
	if convos.Conversations[0].ConversationMessage.Author.ID != "25" {
		t.Errorf("Conversation Message Author not retrieved")
	}
	if convos.Conversations[0].ConversationParts.Parts[0].CreatedAt != 1400857494 {
		t.Errorf("Conversation Part CreatedAt not retrieved")
	}
}

func TestConversationListUserUnread(t *testing.T) {
	http := TestConversationHTTPClient{t: t, expectedURI: "/conversations", fixtureFilename: "fixtures/conversations.json"}
	http.testFunc = func(t *testing.T, queryParams interface{}) {
		ps := queryParams.(conversationListParams)
		if *ps.Unread != true {
			t.Errorf("Expect unread parameter, got %v", *ps.Unread)
		}
	}
	api := ConversationAPI{httpClient: &http}
	api.list(conversationListParams{Unread: Bool(true)})
}

func TestConversationListAdminOpen(t *testing.T) {
	http := TestConversationHTTPClient{t: t, expectedURI: "/conversations", fixtureFilename: "fixtures/conversations.json"}
	http.testFunc = func(t *testing.T, queryParams interface{}) {
		ps := queryParams.(conversationListParams)
		if *ps.Open != true {
			t.Errorf("Expect open parameter, got %v", *ps.Unread)
		}
	}
	api := ConversationAPI{httpClient: &http}
	api.list(conversationListParams{Open: Bool(true)})
}

type TestConversationHTTPClient struct {
	TestHTTPClient
	t               *testing.T
	testFunc        func(t *testing.T, queryParams interface{})
	fixtureFilename string
	expectedURI     string
	lastQueryParams interface{}
}

func (t *TestConversationHTTPClient) Get(uri string, queryParams interface{}) ([]byte, error) {
	if t.testFunc != nil {
		t.testFunc(t.t, queryParams)
	}
	if t.expectedURI != uri {
		t.t.Errorf("Wrong endpoint called")
	}
	return ioutil.ReadFile(t.fixtureFilename)
}
