package intercom

import (
	"errors"
	"testing"
	"time"
)

func TestEventSaveFail(t *testing.T) {
	eventService := EventService{Repository: TestEventAPI{t: t, body: failBody}}
	err := eventService.Save(&Event{})
	if err.Error() != "Missing Identifier" {
		t.Errorf("Error not propagated")
	}
}

func failBody(t *testing.T, event Event) error {
	return errors.New("Missing Identifier")
}

func TestEventSave(t *testing.T) {
	eventService := EventService{Repository: TestEventAPI{t: t, body: successBody}}
	event := Event{}
	event.UserID = "27"
	event.EventName = "govent"
	event.CreatedAt = int64(time.Now().Unix())
	event.Metadata = map[string]interface{}{"is_cool": true}
	eventService.Save(&event)
}

func successBody(t *testing.T, event Event) error {
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
	body func(*testing.T, Event) error
}

func (t TestEventAPI) save(event *Event) error {
	return t.body(t.t, *event)
}
