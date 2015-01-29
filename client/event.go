package client

import "github.com/intercom/intercom-go/domain"

type Event struct {
	domain.Event
	Repository EventRepository
}

type EventRepository interface {
	Save(domain.Event) error
}

func (e Event) Save() error {
	return e.Repository.Save(e.Event)
}
