package intercom

import "fmt"

type EventService struct {
	Repository EventRepository
}

type Event struct {
	Email     string                 `json:"email,omitempty"`
	UserID    string                 `json:"user_id,omitempty"`
	EventName string                 `json:"event_name,omitempty"`
	CreatedAt int32                  `json:"created_at,omitempty"`
	Metadata  map[string]interface{} `json:"metadata,omitempty"`
}

func (e EventService) Save(event *Event) error {
	return e.Repository.save(event)
}

func (e Event) String() string {
	return fmt.Sprintf("[intercom] event { name: %s, user_id: %s, email: %s }", e.EventName, e.UserID, e.Email)
}
