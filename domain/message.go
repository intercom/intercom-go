package domain

import "fmt"

type Message struct {
	Id          string
	Subject     string
	Body        string
	Author      Author
	Attachments []Attachment
}

func (m Message) String() string {
	return fmt.Sprintf("[intercom] message { id: %s }", m.Id)
}
