package intercom

import (
	"encoding/json"

	"github.com/intercom/intercom-go/interfaces"
)

// ConversationRepository defines the interface for working with Conversations through the API.
type ConversationRepository interface {
	list(conversationListParams) (ConversationList, error)
}

// ConversationAPI implements ConversationRepository
type ConversationAPI struct {
	httpClient interfaces.HTTPClient
}

func (api ConversationAPI) list(params conversationListParams) (ConversationList, error) {
	convoList := ConversationList{}
	data, err := api.httpClient.Get("/conversations", params)
	if err != nil {
		return convoList, err
	}
	err = json.Unmarshal(data, &convoList)
	return convoList, err
}
