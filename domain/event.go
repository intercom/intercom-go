package domain

import (
	"fmt"
	"time"
)

type Event struct {
	Id        string
	UserId    string
	Email     string
	EventName string
	CreatedAt time.Time
	Metadata  map[string]interface{}
}

func (e Event) String() string {
	return fmt.Sprintf("[intercom] event { name: %s, id: %s, user_id: %s, email: %s }", e.EventName, e.Id, e.UserId, e.Email)
}
