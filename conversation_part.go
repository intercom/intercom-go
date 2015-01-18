package intercom

import "fmt"

type ConversationPart struct {
	Id          string       `json:"id,omitempty"`
	PartType    string       `json:"part_type,omitempty"`
	CreatedAt   int32        `json:"created_at,omitempty"`
	UpdatedAt   int32        `json:"updated_at,omitempty"`
	NotifiedAt  int32        `json:"updated_at,omitempty"`
	Body        string       `json:"body,omitempty"`
	Author      Author       `json:"author,omitempty"`
	Attachments []Attachment `json:"attachments,omitempty"`
	AssignedTo  Admin        `json:"assigned_to,omitempty"`
}

type ConversationPartList struct {
	Parts []ConversationPart `json:"conversation_parts,omitempty"`
}

func (cp ConversationPart) String() string {
	return fmt.Sprintf("[intercom] conversation_part { id: %s }", cp.Id)
}
