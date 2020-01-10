package intercom

import (
	"fmt"
	"strings"
)

// MessageService handles interactions with the API through an MessageRepository.
type MessageService struct {
	Repository MessageRepository
}

// MessageTemplate determines the template used for email messages to Users or Contacts (plain or personal)
type MessageTemplate int

const (
	NO_TEMPLATE MessageTemplate = iota
	PERSONAL_TEMPLATE
	PLAIN_TEMPLATE
)

var templates = [...]string{
	"",
	"personal",
	"plain",
}

func (template MessageTemplate) String() string {
	return templates[template]
}

func (template *MessageTemplate) UnmarshalJSON(b []byte) error {
	str := strings.Trim(string(b), `"`)

	switch {
	case str == "personal":
		*template = PERSONAL_TEMPLATE
	case str == "plain":
		*template = PLAIN_TEMPLATE
	default:
		*template = NO_TEMPLATE
	}

	return nil
}

// MessageRequest represents a Message to be sent through Intercom from/to an Admin, User, or Contact.
type MessageRequest struct {
	MessageType string         `json:"message_type,omitempty"`
	Subject     string         `json:"subject,omitempty"`
	Body        string         `json:"body,omitempty"`
	Template    string         `json:"template,omitempty"`
	From        MessageAddress `json:"from,omitempty"`
	To          MessageAddress `json:"to,omitempty"`
}

// MessageResponse represents a Message to be sent through Intercom from/to an Admin, User, or Contact.
type MessageResponse struct {
	MessageType string          `json:"message_type,omitempty"`
	ID          string          `json:"id"`
	CreatedAt   int64           `json:"created_at,omitempty"`
	Owner       MessageAddress  `json:"owner,omitempty"`
	Subject     string          `json:"subject,omitempty"`
	Body        string          `json:"body,omitempty"`
	Template    MessageTemplate `json:"template,omitempty"`
}

func (m MessageResponse) String() string {
	return fmt.Sprintf("[intercom] message { id: %s, message_type: %s, body: %s }", m.ID, m.MessageType, m.Body)
}

// Save (send) a Message
func (m *MessageService) Save(message *MessageRequest) (MessageResponse, error) {
	return m.Repository.save(message)
}

// NewEmailMessage creates a new *Message of email type.
func NewEmailMessage(template MessageTemplate, from, to MessagePerson, subject, body string) MessageRequest {
	return MessageRequest{MessageType: "email", Template: template.String(), From: from.MessageAddress(), To: to.MessageAddress(), Subject: subject, Body: body}
}

// NewInAppMessage creates a new *Message of InApp (widget) type.
func NewInAppMessage(from, to MessagePerson, body string) MessageRequest {
	return MessageRequest{MessageType: "inapp", From: from.MessageAddress(), To: to.MessageAddress(), Body: body}
}

// NewUserMessage creates a new *Message from a User.
func NewUserMessage(from MessagePerson, body string) MessageRequest {
	return MessageRequest{MessageType: "inapp", From: from.MessageAddress(), Body: body}
}

// A MessagePerson is someone to send a Message to and from.
type MessagePerson interface {
	MessageAddress() MessageAddress
}

type MessageAddress struct {
	Type   string `json:"type,omitempty"`
	ID     string `json:"id,omitempty"`
	Email  string `json:"email,omitempty"`
	UserID string `json:"user_id,omitempty"`
}
