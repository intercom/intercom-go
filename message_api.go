package intercom

import (
	"encoding/json"

	"github.com/launchpadcentral/intercom-go/interfaces"
)

// MessageRepository defines the interface for creating and updating Messages through the API.
type MessageRepository interface {
	save(message *MessageRequest) (MessageResponse, error)
}

// MessageAPI implements MessageRepository
type MessageAPI struct {
	httpClient interfaces.HTTPClient
}

func (api MessageAPI) save(message *MessageRequest) (MessageResponse, error) {
	data, err := api.httpClient.Post("/messages", message)
	savedMessage := MessageResponse{}
	if err != nil {
		return savedMessage, err
	}
	err = json.Unmarshal(data, &savedMessage)
	return savedMessage, err
}
