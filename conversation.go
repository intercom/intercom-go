package intercom

import "fmt"

type Conversation struct {
	*Resource
	*conversationIdentifiers                      // inline identifiers
	CreatedAt                int32                `json:"created_at,omitempty"`
	UpdatedAt                int32                `json:"updated_at,omitempty"`
	Open                     bool                 `json:"open,omitempty"`
	Read                     bool                 `json:"read,omitempty"`
	User                     User                 `json:"user"`
	Assignee                 Admin                `json:"assignee"`
	Message                  Message              `json:"conversation_message"`
	ConversationParts        ConversationPartList `json:"conversation_parts"`
}

type ConversationList struct {
	Pages         Pages          `json:"pages,omitempty"`
	Conversations []Conversation `json:"conversations"`
}

type conversationIdentifiers struct {
	Id string `json:"id,omitempty" url:"id,omitempty"`
}

type ConversationParams struct {
	Id string
}

func (c Conversation) String() string {
	return fmt.Sprintf("[intercom] conversation { id: %s }", c.Id)
}

// identifiers for a user of a conversation
type userConversationIdentifiers struct {
	IntercomUserId string `json:"intercom_user_id,omitempty" url:"intercom_user_id,omitempty"`
	UserId         string `json:"user_id,omitempty" url:"user_id,omitempty"`
	Email          string `json:"email,omitempty" url:"email,omitempty"`
}

// combined paging parameters, user identifiers, and admin identifiers
type pagedConversationParams struct {
	PageParams
	Type string `json:"type,omitempty" url:"type,omitempty"`
	userConversationIdentifiers
	adminIdentifiers
	Open   *bool `json:"open,omitempty" url:"open,omitempty"` // pointer types as empty value of bool is false
	Unread *bool `json:"unread,omitempty" url:"unread,omitempty"`
}

func (c Conversation) Find(params ConversationParams) (*Conversation, error) {
	if responseBody, err := c.client.Get(fmt.Sprintf("/conversations/%s", params.Id), nil); err != nil {
		return nil, err
	} else {
		conversation := Conversation{}
		return &conversation, c.Unmarshal(&conversation, responseBody.([]byte))
	}
}

func (c Conversation) List(params PageParams) (*ConversationList, error) {
	return c.getConversationsWithParams(params)
}

func (c Conversation) ListForUser(params PageParams, userParams *UserParams) (*ConversationList, error) {
	identifiers := userConversationIdentifiers{userParams.Id, userParams.Email, userParams.UserId}
	return c.getConversationsWithParams(pagedConversationParams{params, "user", identifiers, adminIdentifiers{}, nil, userParams.Unread})
}

func (c Conversation) ListForAdmin(params PageParams, adminParams *AdminParams) (*ConversationList, error) {
	identifiers := adminIdentifiers{adminParams.Id}
	return c.getConversationsWithParams(pagedConversationParams{params, "admin", userConversationIdentifiers{}, identifiers, adminParams.Open, nil})
}

func (c Conversation) getConversationsWithParams(params interface{}) (*ConversationList, error) {
	if responseBody, err := c.client.Get("/conversations", params); err != nil {
		return nil, err
	} else {
		conversationList := ConversationList{}
		return &conversationList, c.Unmarshal(&conversationList, responseBody.([]byte))
	}
}

func (c Conversation) IsUnassigned() bool {
	return c.Assignee.IsNobodyAdmin()
}
