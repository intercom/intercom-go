package intercom

import "fmt"

type Message struct {
	ID          string       `json:"id,omitempty"`
	Subject     string       `json:"subject,omitempty"`
	Body        string       `json:"body,omitempty"`
	Author      Author       `json:"author,omitempty"`
	Attachments []Attachment `json:"attachments"`
}

func (m Message) String() string {
	return fmt.Sprintf("[intercom] message { id: %s }", m.ID)
}
