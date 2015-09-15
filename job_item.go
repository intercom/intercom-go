package intercom

// A JobItem is an item to be processed as part of a bulk Job
type JobItem struct {
	Method   string      `json:"method"`
	DataType string      `json:"data_type"`
	Data     interface{} `json:"data"`
}

// NewUserJobItem creates a JobItem that holds an User.
// It can take either a JOB_POST (for updates) or JOB_DELETE (for deletes) method.
func NewUserJobItem(user *User, method JobItemMethod) *JobItem {
	return &JobItem{Method: method.String(), DataType: "user", Data: user}
}

// NewEventJobItem creates a JobItem that holds an Event.
func NewEventJobItem(event *Event) *JobItem {
	return &JobItem{Method: JOB_POST.String(), DataType: "event", Data: event}
}

type JobItemMethod int

const (
	JOB_POST JobItemMethod = iota
	JOB_DELETE
)

var jobItemMethods = [...]string{
	"post",
	"delete",
}

func (state JobItemMethod) String() string {
	return jobItemMethods[state]
}
