package intercom

import "testing"

func TestFindConversation(t *testing.T) {
	conversationService := ConversationService{Repository: TestConversationAPI{t: t}}
	convo, _ := conversationService.Find("123")
	if convo.ID != "123" {
		t.Errorf("Did not receive conversation")
	}
}

func TestReadConversation(t *testing.T) {
	conversationService := ConversationService{Repository: TestConversationAPI{t: t}}
	convo, _ := conversationService.MarkRead("123")
	if convo.ID != "123" {
		t.Errorf("Did not receive conversation")
	}
}

func TestReplyConversationComment(t *testing.T) {
	testAPI := TestConversationAPI{t: t}
	testAPI.testFunc = func(t *testing.T, reply interface{}) {
		if reply.(*Reply).IntercomID != "abc123" {
			t.Errorf("user id not supplied")
		}
		if reply.(*Reply).ReplyType != "comment" {
			t.Errorf("part was not comment, was %s", reply.(*Reply).ReplyType)
		}
	}
	conversationService := ConversationService{Repository: testAPI}
	_, err := conversationService.Reply("123", &Contact{ID: "abc123"}, CONVERSATION_COMMENT, "Body")
	if err != nil {
		t.Errorf("Failed to add conversation reply: %s", err)
	}
}

func TestReplyConversationCommentWithAttachment(t *testing.T) {
	testAPI := TestConversationAPI{t: t}
	testAPI.testFunc = func(t *testing.T, reply interface{}) {
		if reply.(*Reply).IntercomID != "abc123" {
			t.Errorf("user id not supplied")
		}
		if reply.(*Reply).ReplyType != "comment" {
			t.Errorf("part was not comment, was %s", reply.(*Reply).ReplyType)
		}
	}
	conversationService := ConversationService{Repository: testAPI}
	_, err := conversationService.ReplyWithAttachmentURLs("123", &Contact{ID: "abc123"}, CONVERSATION_COMMENT, "Body", []string{"https://www.example.com/attachment.jpg"})
	if err != nil {
		t.Errorf("Failed to add conversation reply with attachment: %s", err)
	}
}

func TestReplyConversationOpen(t *testing.T) {
	testAPI := TestConversationAPI{t: t}
	testAPI.testFunc = func(t *testing.T, reply interface{}) {
		if reply.(*Reply).IntercomID != "abc123" {
			t.Errorf("user id not supplied")
		}
		if reply.(*Reply).ReplyType != "open" {
			t.Errorf("part was not open, was %s", reply.(*Reply).ReplyType)
		}
	}
	conversationService := ConversationService{Repository: testAPI}
	_, err := conversationService.Reply("123", &Contact{ID: "abc123"}, CONVERSATION_OPEN, "Body")
	if err != nil {
		t.Errorf("Failed to add conversation reply for open: %s", err)
	}
}

func TestReplyConversationNote(t *testing.T) {
	testAPI := TestConversationAPI{t: t}
	testAPI.testFunc = func(t *testing.T, reply interface{}) {
		if reply.(*Reply).AdminID != "abc123" {
			t.Errorf("admin id not supplied")
		}
		if reply.(*Reply).ReplyType != "note" {
			t.Errorf("part was not note, was %s", reply.(*Reply).ReplyType)
		}
	}
	conversationService := ConversationService{Repository: testAPI}
	_, err := conversationService.Reply("123", &Admin{ID: "abc123"}, CONVERSATION_NOTE, "Body")
	if err != nil {
		t.Errorf("Failed to add conversation reply with note: %s", err)
	}
}

func TestAssignConversation(t *testing.T) {
	testAPI := TestConversationAPI{t: t}
	testAPI.testFunc = func(t *testing.T, reply interface{}) {
		if reply.(*Reply).AssigneeID != "def789" {
			t.Errorf("assignee_id not supplied")
		}
		if reply.(*Reply).AdminID != "abc123" {
			t.Errorf("admin id was not supplied")
		}
		if reply.(*Reply).ReplyType != "assignment" {
			t.Errorf("part was not assignment, was %s", reply.(*Reply).ReplyType)
		}
	}
	conversationService := ConversationService{Repository: testAPI}
	_, err := conversationService.Assign("123", &Admin{ID: "abc123"}, &Admin{ID: "def789"})
	if err != nil {
		t.Errorf("Failed to assign conversation: %s", err)
	}
}

func TestListAllConversations(t *testing.T) {
	conversationService := ConversationService{Repository: TestConversationAPI{t: t}}
	list, _ := conversationService.ListAll(PageParams{})
	if list.Conversations[0].ID != "123" {
		t.Errorf("did not receive conversation")
	}
}

func TestListUserConversationsUnread(t *testing.T) {
	testAPI := TestConversationAPI{t: t}
	testAPI.testFunc = func(t *testing.T, params interface{}) {
		if *params.(ConversationListParams).Unread != true {
			t.Errorf("unread was %v, expected true", *params.(ConversationListParams).Unread)
		}
	}
	conversationService := ConversationService{Repository: testAPI}
	contact := Contact{}
	list, _ := conversationService.ListByContact(&contact, SHOW_UNREAD, PageParams{})
	if list.Conversations[0].ID != "123" {
		t.Errorf("did not receive conversation")
	}
}

func TestListUserConversationsAll(t *testing.T) {
	testAPI := TestConversationAPI{t: t}
	testAPI.testFunc = func(t *testing.T, params interface{}) {
		if params.(ConversationListParams).Unread != nil {
			t.Errorf("unread was not nil, was %v", *params.(ConversationListParams).Unread)
		}
	}
	conversationService := ConversationService{Repository: testAPI}
	contact := Contact{}
	list, _ := conversationService.ListByContact(&contact, SHOW_ALL, PageParams{})
	if list.Conversations[0].ID != "123" {
		t.Errorf("did not receive conversation")
	}
}

func TestListAdminConversationsAll(t *testing.T) {
	testAPI := TestConversationAPI{t: t}
	testAPI.testFunc = func(t *testing.T, params interface{}) {
		if params.(ConversationListParams).Open != nil {
			t.Errorf("open was not nil, was %v", *params.(ConversationListParams).Open)
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
	testAPI.testFunc = func(t *testing.T, params interface{}) {
		if *params.(ConversationListParams).Open != true {
			t.Errorf("open was not true, was %v", *params.(ConversationListParams).Open)
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
	testFunc func(t *testing.T, params interface{})
	t        *testing.T
}

func (t TestConversationAPI) list(params ConversationListParams) (ConversationList, error) {
	if t.testFunc != nil {
		t.testFunc(t.t, params)
	}
	return ConversationList{Conversations: []Conversation{{ID: "123"}}, Pages: PageParams{Page: 1, PerPage: 20}}, nil
}

func (t TestConversationAPI) find(id string) (Conversation, error) {
	return Conversation{ID: "123"}, nil
}

func (t TestConversationAPI) read(id string) (Conversation, error) {
	return Conversation{ID: "123"}, nil
}

func (t TestConversationAPI) reply(id string, reply *Reply) (Conversation, error) {
	if t.testFunc != nil {
		t.testFunc(t.t, reply)
	}
	return Conversation{ID: "123"}, nil
}
