package client

import (
	"errors"
	"testing"
	"time"

	"github.com/intercom/intercom-go/domain"
)

func TestEventSaveFail(t *testing.T) {
	event := Event{Repository: TestEventAPI{t: t, body: failBody}}
	err := event.Save()
	if err.Error() != "Missing Identifier" {
		t.Errorf("Error not propagated")
	}
}

func failBody(t *testing.T, event domain.Event) error {
	return errors.New("Missing Identifier")
}

func TestEventSave(t *testing.T) {
	event := Event{Repository: TestEventAPI{t: t, body: successBody}}
	event.UserID = "27"
	event.EventName = "govent"
	event.CreatedAt = int32(time.Now().Unix())
	event.Metadata = map[string]interface{}{"is_cool": true}
	event.Save()
}

func successBody(t *testing.T, event domain.Event) error {
	if event.UserID != "27" {
		t.Errorf("UserID not set")
	}
	if event.EventName != "govent" {
		t.Errorf("EventName not set")
	}
	if event.Metadata["is_cool"] != true {
		t.Errorf("Metadata not set")
	}
	return nil
}

type TestEventAPI struct {
	t    *testing.T
	body func(*testing.T, domain.Event) error
}

func (t TestEventAPI) Save(event domain.Event) error {
	return t.body(t.t, event)
}
