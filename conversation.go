package intercom

import "fmt"

type ConversationService struct {
	Repository ConversationRepository
}

type Conversation struct {
	ID        string  `json:"id,omitempty"`
	CreatedAt int32   `json:"created_at,omitempty"`
	UpdatedAt int32   `json:"updated_at,omitempty"`
	Open      bool    `json:"open,omitempty"`
	Read      bool    `json:"read,omitempty"`
	User      User    `json:"user,ommitempty"`
	Assignee  Admin   `json:"assignee,omitempty"`
	Message   Message `json:"conversation_message,omitempty"`
	// Parts     []ConversationPart
}

func (c Conversation) String() string {
	return fmt.Sprintf("[intercom] conversation { id: %s }", c.ID)
}

func (c *ConversationService) Find(ID string) (Conversation, error) {
	return c.Repository.find(ID)
}

// Reply
// ListByAdmin
// ListByUser
// ReplyAsAdmin
// ReplyAsUser
// MarkRead
// NewFromAdmin
// NewFromUser
