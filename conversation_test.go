package intercom

import "testing"

func TestListAllConversations(t *testing.T) {
	conversationService := ConversationService{Repository: TestConversationAPI{t: t}}
	list, _ := conversationService.ListAll(PageParams{})
	if list.Conversations[0].ID != "123" {
		t.Errorf("did not receive conversation")
	}
}

func TestListUserConversationsUnread(t *testing.T) {
	testAPI := TestConversationAPI{t: t}
	testAPI.testFunc = func(t *testing.T, params conversationListParams) {
		if *params.Unread != true {
			t.Errorf("unread was %v, expected true", *params.Unread)
		}
	}
	conversationService := ConversationService{Repository: testAPI}
	user := User{}
	list, _ := conversationService.ListByUser(&user, SHOW_UNREAD, PageParams{})
	if list.Conversations[0].ID != "123" {
		t.Errorf("did not receive conversation")
	}
}

func TestListUserConversationsAll(t *testing.T) {
	testAPI := TestConversationAPI{t: t}
	testAPI.testFunc = func(t *testing.T, params conversationListParams) {
		if params.Unread != nil {
			t.Errorf("unread was not nil, was %v", *params.Unread)
		}
	}
	conversationService := ConversationService{Repository: testAPI}
	user := User{}
	list, _ := conversationService.ListByUser(&user, SHOW_ALL, PageParams{})
	if list.Conversations[0].ID != "123" {
		t.Errorf("did not receive conversation")
	}
}

func TestListAdminConversationsAll(t *testing.T) {
	testAPI := TestConversationAPI{t: t}
	testAPI.testFunc = func(t *testing.T, params conversationListParams) {
		if params.Open != nil {
			t.Errorf("open was not nil, was %v", *params.Open)
		}
	}
	conversationService := ConversationService{Repository: testAPI}
	admin := Admin{}
	list, _ := conversationService.ListByAdmin(&admin, SHOW_ALL, PageParams{})
	if list.Conversations[0].ID != "123" {
		t.Errorf("did not receive conversation")
	}
}

func TestListAdminConversationsOpen(t *testing.T) {
	testAPI := TestConversationAPI{t: t}
	testAPI.testFunc = func(t *testing.T, params conversationListParams) {
		if *params.Open != true {
			t.Errorf("open was not true, was %v", *params.Open)
		}
	}
	conversationService := ConversationService{Repository: testAPI}
	admin := Admin{}
	list, _ := conversationService.ListByAdmin(&admin, SHOW_OPEN, PageParams{})
	if list.Conversations[0].ID != "123" {
		t.Errorf("did not receive conversation")
	}
}

type TestConversationAPI struct {
	testFunc func(t *testing.T, params conversationListParams)
	t        *testing.T
}

func (t TestConversationAPI) list(params conversationListParams) (ConversationList, error) {
	if t.testFunc != nil {
		t.testFunc(t.t, params)
	}
	return ConversationList{Conversations: []Conversation{Conversation{ID: "123"}}, Pages: PageParams{Page: 1, PerPage: 20}}, nil
}
