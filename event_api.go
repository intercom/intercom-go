package intercom

import "gopkg.in/intercom/intercom-go.v2/interfaces"

// EventRepository defines the interface for working with Events through the API.
type EventRepository interface {
	save(*Event) error
}

// EventAPI implements EventRepository
type EventAPI struct {
	httpClient interfaces.HTTPClient
}

func (api EventAPI) save(event *Event) error {
	_, err := api.httpClient.Post("/events", event)
	return err
}
