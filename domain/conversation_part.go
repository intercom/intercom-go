package domain

import "fmt"

type ConversationPart struct {
	Id          string
	PartType    string
	CreatedAt   int32
	UpdatedAt   int32
	NotifiedAt  int32
	Body        string
	Author      Author
	Attachments []Attachment
	AssignedTo  Admin
}

func (cp ConversationPart) String() string {
	return fmt.Sprintf("[intercom] conversation_part { id: %s }", cp.Id)
}
