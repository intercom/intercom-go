package domain

import (
	"fmt"
	"time"
)

type User struct {
	ID         string
	Email      string
	UserID     string
	SignedUp   time.Time
	Name       string
	CustomData map[string]interface{}
}

func (u User) AddEvent(event *Event) {
	event.ID = u.ID
	event.UserID = u.UserID
	event.Email = u.Email
}

func (u User) AddNote(note *Note) {
	note.User = u
}

func (u User) String() string {
	return fmt.Sprintf("[intercom] user { id: %s name: %s, user_id: %s, email: %s }", u.ID, u.Name, u.UserID, u.Email)
}

func (u User) authorType() string {
	return "user"
}
