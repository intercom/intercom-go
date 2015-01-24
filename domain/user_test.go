package domain

import (
	"testing"
)

func TestAddEvent(t *testing.T) {
	user := User{ID: "46adad3f09126dca", Email: "jamie@example.io", UserID: "aa123"}
	event := Event{}
	user.AddEvent(&event)
	if event.ID != "46adad3f09126dca" {
		t.Errorf("ID not set")
	}
	if event.Email != "jamie@example.io" {
		t.Errorf("Email not set")
	}
	if event.UserID != "aa123" {
		t.Errorf("User ID not set")
	}
}

func TestAddNote(t *testing.T) {
	user := User{ID: "46adad3f09126dca", Email: "jamie@example.io", UserID: "aa123"}
	note := Note{}
	user.AddNote(&note)
	if note.User.ID != "46adad3f09126dca" {
		t.Errorf("User not set on Note")
	}
}
