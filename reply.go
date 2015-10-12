package intercom

// A Reply to an Intercom conversation
type Reply struct {
	Type           string   `json:"type"`
	ReplyType      string   `json:"message_type"`
	Body           string   `json:"body,omitempty"`
	AssigneeID     string   `json:"assignee_id,omitempty"`
	AdminID        string   `json:"admin_id,omitempty"`
	IntercomID     string   `json:"intercom_user_id,omitempty"`
	Email          string   `json:"email,omitempty"`
	UserID         string   `json:"user_id,omitempty"`
	AttachmentURLs []string `json:"attachment_urls,omitempty"`
}

// ReplyType determines the type of Reply
type ReplyType int

const (
	CONVERSATION_COMMENT ReplyType = iota
	CONVERSATION_NOTE
	CONVERSATION_ASSIGN
	CONVERSATION_OPEN
	CONVERSATION_CLOSE
)

var replyTypes = [...]string{
	"comment",
	"note",
	"assignment",
	"open",
	"close",
}

func (reply ReplyType) String() string {
	return replyTypes[reply]
}
