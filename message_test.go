package intercom

import "testing"

func TestNewEmailMessage(t *testing.T) {
	user := User{}
	admin := Admin{}
	message := NewEmailMessage(PERSONAL_TEMPLATE, admin, user, "subject", "body")
	if message.MessageType != "email" {
		t.Errorf("Message type was not email")
	}
	if message.Template != "personal" {
		t.Errorf("Message template was not personal")
	}
	if message.From.Type != "admin" {
		t.Errorf("Message from was not admin")
	}
	if message.To.Type != "user" {
		t.Errorf("Message to was not user")
	}
	if message.Subject != "subject" {
		t.Errorf("Message subject was not set")
	}
	if message.Body != "body" {
		t.Errorf("Message body was not set")
	}
}

func TestNewInAppMessage(t *testing.T) {
	contact := Contact{}
	admin := Admin{}
	message := NewInAppMessage(admin, contact, "body")
	if message.MessageType != "inapp" {
		t.Errorf("Message type was not inapp")
	}
	if message.Template != "" {
		t.Errorf("Message template was not empty, was %s", message.Template)
	}
	if message.From.Type != "admin" {
		t.Errorf("Message from was not admin")
	}
	if message.To.Type != "contact" {
		t.Errorf("Message to was not contact")
	}
	if message.Subject != "" {
		t.Errorf("Message subject was not empty")
	}
	if message.Body != "body" {
		t.Errorf("Message body was not set")
	}
}

func TestNewUserMessage(t *testing.T) {
	user := User{}
	message := NewUserMessage(user, "body")
	if message.MessageType != "inapp" {
		t.Errorf("Message type was not inapp")
	}
	if message.Template != "" {
		t.Errorf("Message template was not empty, was %s", message.Template)
	}
	if message.From.Type != "user" {
		t.Errorf("Message from was not user")
	}
	if message.To.Type != "" {
		t.Errorf("Message to was not empty")
	}
	if message.Subject != "" {
		t.Errorf("Message subject was not empty")
	}
	if message.Body != "body" {
		t.Errorf("Message body was not set")
	}
}

func TestSaveMessage(t *testing.T) {
	messageService := MessageService{Repository: TestMessageAPI{t: t}}
	message := NewInAppMessage(Admin{}, User{}, "hi there")
	resp, _ := messageService.Save(&message)
	if resp.Owner.Type != "admin" {
		t.Errorf("Owner was not admin")
	}
}

type TestMessageAPI struct {
	t *testing.T
}

func (t TestMessageAPI) save(message *MessageRequest) (MessageResponse, error) {
	if message.MessageType != "inapp" {
		t.t.Errorf("Message not inapp")
	}
	if message.To.Type != "user" {
		t.t.Errorf("Message not sent to user")
	}
	return MessageResponse{MessageType: message.MessageType, Owner: message.From, Body: message.Body}, nil
}
