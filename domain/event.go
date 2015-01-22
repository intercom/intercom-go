package domain

import (
	"fmt"
	"time"
)

type Event struct {
	ID        string
	UserID    string
	Email     string
	EventName string
	CreatedAt time.Time
	Metadata  map[string]interface{}
}

func (e Event) String() string {
	return fmt.Sprintf("[intercom] event { name: %s, id: %s, user_id: %s, email: %s }", e.EventName, e.ID, e.UserID, e.Email)
}
