package intercom

import (
	"io/ioutil"
	"testing"
)

func TestConversationFind(t *testing.T) {
	http := TestConversationHTTPClient{t: t, expectedURI: "/conversations/147", fixtureFilename: "fixtures/conversation.json"}
	api := ConversationAPI{httpClient: &http}
	convo, _ := api.find("147")
	if convo.ID != "147" {
		t.Errorf("Conversation not retrieved, %s", convo.ID)
	}
	if convo.TagList == nil || convo.TagList.Tags[0].ID != "12345" {
		t.Errorf("Conversation tags not retrieved, %s", convo.ID)
	}
}

func TestConversationRead(t *testing.T) {
	http := TestConversationHTTPClient{t: t, expectedURI: "/conversations/147", fixtureFilename: "fixtures/conversation.json"}
	http.testFunc = func(t *testing.T, readRequest interface{}) {
		req := readRequest.(conversationReadRequest)
		if req.Read != true {
			t.Errorf("read was not marked true")
		}
	}
	api := ConversationAPI{httpClient: &http}
	convo, err := api.read("147")
	if err != nil {
		t.Errorf("%v", err)
	}
	if convo.ID != "147" {
		t.Errorf("Conversation not retrieved, %s", convo.ID)
	}
}

func TestConversationReply(t *testing.T) {
	http := TestConversationHTTPClient{t: t, expectedURI: "/conversations/147/reply", fixtureFilename: "fixtures/conversation.json"}
	http.testFunc = func(t *testing.T, replyRequest interface{}) {
		reply := replyRequest.(*Reply)
		if reply.ReplyType != CONVERSATION_NOTE.String() {
			t.Errorf("Reply was not note")
		}
	}
	api := ConversationAPI{httpClient: &http}
	convo, err := api.reply("147", &Reply{ReplyType: CONVERSATION_NOTE.String(), AdminID: "123"})
	if err != nil {
		t.Errorf("%v", err)
	}
	if convo.ID != "147" {
		t.Errorf("Conversation not retrieved, %s", convo.ID)
	}
}

func TestConversationReplyWithAttachment(t *testing.T) {
	http := TestConversationHTTPClient{t: t, expectedURI: "/conversations/147/reply", fixtureFilename: "fixtures/conversation.json"}
	http.testFunc = func(t *testing.T, replyRequest interface{}) {
		reply := replyRequest.(*Reply)
		if reply.ReplyType != CONVERSATION_COMMENT.String() {
			t.Errorf("Reply was not comment")
		}
	}
	api := ConversationAPI{httpClient: &http}
	convo, err := api.reply("147", &Reply{ReplyType: CONVERSATION_COMMENT.String(), AdminID: "123", AttachmentURLs: []string{"http://www.example.com/attachment.jpg"}})
	if err != nil {
		t.Errorf("%v", err)
	}
	if convo.ID != "147" {
		t.Errorf("Conversation not retrieved, %s", convo.ID)
	}
}

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
	if convos.Conversations[0].TagList != nil {
		t.Errorf("Conversation Tags should be nil")
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

func (t *TestConversationHTTPClient) Post(uri string, dataObject interface{}) ([]byte, error) {
	if t.testFunc != nil {
		t.testFunc(t.t, dataObject)
	}
	if t.expectedURI != uri {
		t.t.Errorf("Wrong endpoint called")
	}
	return ioutil.ReadFile(t.fixtureFilename)
}
