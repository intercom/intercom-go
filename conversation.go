package intercom

// ConversationService handles interactions with the API through an ConversationRepository.
type ConversationService struct {
	Repository ConversationRepository
}

// ConversationList is a list of Conversations
type ConversationList struct {
	Pages         PageParams     `json:"pages"`
	Conversations []Conversation `json:"conversations"`
}

// A Conversation represents a conversation between users and admins in Intercom.
type Conversation struct {
	ID                  string               `json:"id"`
	CreatedAt           int32                `json:"created_at"`
	UpdatedAt           int32                `json:"updated_at"`
	User                User                 `json:"user"`
	Assignee            Admin                `json:"assignee"`
	Open                bool                 `json:"open"`
	Read                bool                 `json:"read"`
	ConversationMessage ConversationMessage  `json:"conversation_message"`
	ConversationParts   ConversationPartList `json:"conversation_parts"`
}

// A ConversationMessage is the message that started the conversation rendered for presentation
type ConversationMessage struct {
	Subject string         `json:"subject"`
	Body    string         `json:"body"`
	Author  MessageAddress `json:"author"`
}

// A ConversationPartList lists the subsequent Conversation Parts
type ConversationPartList struct {
	Parts []ConversationPart `json:"conversation_parts"`
}

// A ConversationPart is a Reply, Note, or Assignment to a Conversation
type ConversationPart struct {
	ID         string         `json:"id"`
	PartType   string         `json:"part_type"`
	Body       string         `json:"body"`
	CreatedAt  int32          `json:"created_at"`
	UpdatedAt  int32          `json:"updated_at"`
	NotifiedAt int32          `json:"notified_at"`
	AssignedTo Admin          `json:"assigned_to"`
	Author     MessageAddress `json:"author"`
}

// The state of Conversations to query
// SHOW_ALL shows all conversations,
// SHOW_OPEN shows only open conversations (only valid for Admin Conversation queries)
// SHOW_CLOSED shows only closed conversations (only valid for Admin Conversation queries)
// SHOW_UNREAD shows only unread conversations (only valid for User Conversation queries)
type ConversationListState int

const (
	SHOW_ALL ConversationListState = iota
	SHOW_OPEN
	SHOW_CLOSED
	SHOW_UNREAD
)

// List all Conversations
func (c *ConversationService) ListAll(pageParams PageParams) (ConversationList, error) {
	return c.Repository.list(conversationListParams{PageParams: pageParams})
}

// List Conversations by Admin
func (c *ConversationService) ListByAdmin(admin *Admin, state ConversationListState, pageParams PageParams) (ConversationList, error) {
	params := conversationListParams{
		PageParams: pageParams,
		Type:       "admin",
		AdminID:    admin.ID.String(),
	}
	if state == SHOW_OPEN {
		params.Open = Bool(true)
	}
	if state == SHOW_CLOSED {
		params.Open = Bool(false)
	}
	return c.Repository.list(params)
}

// List Conversations by User
func (c *ConversationService) ListByUser(user *User, state ConversationListState, pageParams PageParams) (ConversationList, error) {
	params := conversationListParams{
		PageParams:     pageParams,
		Type:           "user",
		IntercomUserID: user.ID,
		UserID:         user.UserID,
		Email:          user.Email,
	}
	if state == SHOW_UNREAD {
		params.Unread = Bool(true)
	}
	return c.Repository.list(params)
}

type conversationListParams struct {
	PageParams
	Type           string `url:"type,omitempty"`
	AdminID        string `url:"admin_id,omitempty"`
	IntercomUserID string `url:"intercom_user_id,omitempty"`
	UserID         string `url:"user_id,omitempty"`
	Email          string `url:"email,omitempty"`
	Open           *bool  `url:"open,omitempty"`
	Unread         *bool  `url:"unread,omitempty"`
	DisplayAs      string `url:"display_as,omitempty"`
}
