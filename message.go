package intercom

import "fmt"

type Message struct {
	Id          string       `json:"id,omitempty"`
	Subject     string       `json:"subject,omitempty"`
	Body        string       `json:"body,omitempty"`
	Author      Author       `json:"author,omitempty"`
	Attachments []Attachment `json:"attachments,omitempty"`
}

func (m Message) String() string {
	return fmt.Sprintf("[intercom] message { id: %s }", m.Id)
}
