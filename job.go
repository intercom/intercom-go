package intercom

import "fmt"

// JobService builds jobs to process
type JobService struct {
	Repository JobRepository
}

// The state of a Job
type JobState int

const (
	PENDING JobState = iota
	RUNNING
	COMPLETED
	FAILED
)

var jobStates = [...]string{
	"pending",
	"running",
	"completed",
	"failed",
}

// A JobRequest represents a new job to be sent to Intercom
type JobRequest struct {
	JobData *JobData   `json:"job,omitempty"`
	Items   []*JobItem `json:"items,omitempty"`

	bulkType string
}

// A JobResponse represents a job enqueud on Intercom
type JobResponse struct {
	ID          string            `json:"id,omitempty"`
	AppID       string            `json:"app_id,omitempty"`
	UpdatedAt   int64             `json:"updated_at,omitempty"`
	CreatedAt   int64             `json:"created_at,omitempty"`
	CompletedAt int64             `json:"completed_at,omitempty"`
	ClosingAt   int64             `json:"closing_at,omitempty"`
	Name        string            `json:"name,omitempty"`
	State       string            `json:"job_state,omitempty"`
	Links       map[string]string `json:"links,omitempty"`
}

// JobData is a payload that can be used to identify an existing Job to append to.
type JobData struct {
	ID string `json:"id,omitempty"`
}

// NewUserJob creates a new Job for processing Users.
func (js *JobService) NewUserJob(items ...*JobItem) (JobResponse, error) {
	job := JobRequest{Items: items, bulkType: "users"}
	return js.Repository.save(&job)
}

// NewEventJob creates a new Job for processing Events.
func (js *JobService) NewEventJob(items ...*JobItem) (JobResponse, error) {
	job := JobRequest{Items: items, bulkType: "events"}
	return js.Repository.save(&job)
}

// Append User items to existing Job
func (js *JobService) AppendUsers(id string, items ...*JobItem) (JobResponse, error) {
	job := JobRequest{JobData: &JobData{ID: id}, Items: items, bulkType: "users"}
	return js.Repository.save(&job)
}

// Append Event items to existing Job
func (js *JobService) AppendEvents(id string, items ...*JobItem) (JobResponse, error) {
	job := JobRequest{JobData: &JobData{ID: id}, Items: items, bulkType: "events"}
	return js.Repository.save(&job)
}

// Find existing Job
func (js *JobService) Find(id string) (JobResponse, error) {
	return js.Repository.find(id)
}

func (j JobResponse) String() string {
	return fmt.Sprintf("[intercom] job { id: %s, name: %s}", j.ID, j.Name)
}

func (state JobState) String() string {
	return jobStates[state]
}
