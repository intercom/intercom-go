package domain

import "fmt"

type User struct {
	ID               string                 `json:"id"`
	Email            string                 `json:"email"`
	UserID           string                 `json:"user_id"`
	SignedUpAt       int64                  `json:"signed_up_at"`
	Name             string                 `json:"name"`
	CustomAttributes map[string]interface{} `json:"custom_attributes"`
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
