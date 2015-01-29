package domain

import (
	"testing"
)

func TestAddNote(t *testing.T) {
	user := User{ID: "46adad3f09126dca", Email: "jamie@example.io", UserID: "aa123"}
	note := Note{}
	user.AddNote(&note)
	if note.User.ID != "46adad3f09126dca" {
		t.Errorf("User not set on Note")
	}
}
