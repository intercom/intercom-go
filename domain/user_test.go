package domain

import (
	"testing"
)

func TestAddEvent(t *testing.T) {
	user := User{Id: "46adad3f09126dca", Email: "jamie@intercom.io", UserId: "aa123"}
	event := Event{}
	user.AddEvent(&event)
	if "46adad3f09126dca" != event.Id {
		t.Errorf("Id not set")
	}
	if "jamie@intercom.io" != event.Email {
		t.Errorf("Email not set")
	}
	if "aa123" != event.UserId {
		t.Errorf("User ID not set")
	}
}

func TestAddNote(t *testing.T) {
	user := User{Id: "46adad3f09126dca", Email: "jamie@intercom.io", UserId: "aa123"}
	note := Note{}
	user.AddNote(&note)
	if "46adad3f09126dca" != note.User.Id {
		t.Errorf("User not set on Note")
	}
}
