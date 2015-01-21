package domain

import (
	"fmt"
	"time"
)

type User struct {
	Id         string
	Email      string
	UserId     string
	SignedUp   time.Time
	Name       string
	CustomData map[string]interface{}
}

func (u User) AddEvent(event *Event) {
	event.Id = u.Id
	event.UserId = u.UserId
	event.Email = u.Email
}

func (u User) AddNote(note *Note) {
	note.User = u
}

func (u User) String() string {
	return fmt.Sprintf("[intercom] user { id: %s name: %s, user_id: %s, email: %s }", u.Id, u.Name, u.UserId, u.Email)
}

func (u User) authorType() string {
	return "user"
}
