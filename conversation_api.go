package intercom

import (
	"encoding/json"
	"fmt"

	"gopkg.in/intercom/intercom-go.v2/interfaces"
)

// ConversationRepository defines the interface for working with Conversations through the API.
type ConversationRepository interface {
	find(id string) (Conversation, error)
	list(params conversationListParams) (ConversationList, error)
	read(id string) (Conversation, error)
	reply(id string, reply *Reply) (Conversation, error)
}

// ConversationAPI implements ConversationRepository
type ConversationAPI struct {
	httpClient interfaces.HTTPClient
}

type conversationReadRequest struct {
	Read bool `json:"read"`
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

func (api ConversationAPI) read(id string) (Conversation, error) {
	conversation := Conversation{}
	data, err := api.httpClient.Post(fmt.Sprintf("/conversations/%s", id), conversationReadRequest{Read: true})
	if err != nil {
		return conversation, err
	}
	err = json.Unmarshal(data, &conversation)
	return conversation, err
}

func (api ConversationAPI) reply(id string, reply *Reply) (Conversation, error) {
	conversation := Conversation{}
	data, err := api.httpClient.Post(fmt.Sprintf("/conversations/%s/reply", id), reply)
	if err != nil {
		return conversation, err
	}
	err = json.Unmarshal(data, &conversation)
	return conversation, nil
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
