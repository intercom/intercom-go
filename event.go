package intercom

import "fmt"

// EventService handles interactions with the API through an EventRepository.
type EventService struct {
	Repository EventRepository
}

// An Event represents a new event that happens to a User.
type Event struct {
	ID        string                 `json:"id,omitempty"`
	Email     string                 `json:"email,omitempty"`
	UserID    string                 `json:"user_id,omitempty"`
	EventName string                 `json:"event_name,omitempty"`
	CreatedAt int64                  `json:"created_at,omitempty"`
	Metadata  map[string]interface{} `json:"metadata,omitempty"`
}

// Save a new Event
func (e *EventService) Save(event *Event) error {
	return e.Repository.save(event)
}

func (e Event) String() string {
	return fmt.Sprintf("[intercom] event { name: %s, user_id: %s, email: %s }", e.EventName, e.UserID, e.Email)
}
