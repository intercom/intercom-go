package intercom

import "fmt"

type Conversation struct {
	Id        string
	CreatedAt int32
	UpdatedAt int32
	Open      bool
	Read      bool
	User      User
	Assignee  Admin
	Message   Message
	Parts     []ConversationPart
}

func (c Conversation) String() string {
	return fmt.Sprintf("[intercom] conversation { id: %s }", c.Id)
}
