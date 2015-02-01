package intercom

import (
	"encoding/json"
	"fmt"

	"github.com/intercom/intercom-go/interfaces"
)

type ConversationRepository interface {
	find(string) (Conversation, error)
}

type ConversationAPI struct {
	httpClient interfaces.HTTPClient
}

func (api ConversationAPI) find(id string) (Conversation, error) {
	conversation := Conversation{}
	data, err := api.httpClient.Get(fmt.Sprintf("/conversations/%s", id), nil)
	if err != nil {
		return conversation, err
	}
	err = json.Unmarshal(data, &conversation)
	return conversation, err
}

// list
// reply
// markRead
// new
