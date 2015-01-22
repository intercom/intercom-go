package domain

import (
	"testing"
)

func TestAddEvent(t *testing.T) {
	user := User{ID: "46adad3f09126dca", Email: "jamie@intercom.io", UserID: "aa123"}
	event := Event{}
	user.AddEvent(&event)
	if "46adad3f09126dca" != event.ID {
		t.Errorf("ID not set")
	}
	if "jamie@intercom.io" != event.Email {
		t.Errorf("Email not set")
	}
	if "aa123" != event.UserID {
		t.Errorf("User ID not set")
	}
}

func TestAddNote(t *testing.T) {
	user := User{ID: "46adad3f09126dca", Email: "jamie@intercom.io", UserID: "aa123"}
	note := Note{}
	user.AddNote(&note)
	if "46adad3f09126dca" != note.User.ID {
		t.Errorf("User not set on Note")
	}
}
