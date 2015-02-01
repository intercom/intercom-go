package intercom

import "testing"

func TestConversationFind(t *testing.T) {
	conversationService := ConversationService{Repository: TestConversationAPI{t: t}}
	conversation, _ := conversationService.Find("213")
	if conversation.ID != "213" {
		t.Errorf("Conversation not found")
	}
}

type TestConversationAPI struct {
	t *testing.T
}

func (t TestConversationAPI) find(ID string) (Conversation, error) {
	return Conversation{ID: ID}, nil
}
