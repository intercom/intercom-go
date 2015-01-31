package intercom

import "github.com/intercom/intercom-go/interfaces"

type EventRepository interface {
	save(*Event) error
}

type EventAPI struct {
	httpClient interfaces.HTTPClient
}

func (api EventAPI) save(event *Event) error {
	_, err := api.httpClient.Post("/events", event)
	return err
}
