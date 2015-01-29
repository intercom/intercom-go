package domain

import "fmt"

type User struct {
	ID               string                 `json:"id,omitempty"`
	Email            string                 `json:"email,omitempty"`
	UserID           string                 `json:"user_id,omitempty"`
	SignedUpAt       int32                  `json:"signed_up_at,omitempty"`
	Name             string                 `json:"name,omitempty"`
	CustomAttributes map[string]interface{} `json:"custom_attributes,omitempty"`
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
