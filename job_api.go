package intercom

import (
	"encoding/json"
	"fmt"

	"gopkg.in/intercom/intercom-go.v2/interfaces"
)

// JobRepository defines the interface for working with Jobs.
type JobRepository interface {
	save(job *JobRequest) (JobResponse, error)
	find(id string) (JobResponse, error)
}

// JobAPI implements TagRepository
type JobAPI struct {
	httpClient interfaces.HTTPClient
}

func (api JobAPI) save(job *JobRequest) (JobResponse, error) {
	for i := range job.Items {
		obj := job.Items[i].Data
		switch obj.(type) {
		case *User:
			user := obj.(*User)
			job.Items[i].Data = RequestUserMapper{}.ConvertUser(user)
		}
	}
	savedJob := JobResponse{}
	data, err := api.httpClient.Post(fmt.Sprintf("/bulk/%s", job.bulkType), job)
	if err != nil {
		return savedJob, err
	}
	err = json.Unmarshal(data, &savedJob)
	return savedJob, err
}

func (api JobAPI) find(id string) (JobResponse, error) {
	fetchedJob := JobResponse{}
	data, err := api.httpClient.Get(fmt.Sprintf("/jobs/%s", id), nil)
	if err != nil {
		return fetchedJob, err
	}
	err = json.Unmarshal(data, &fetchedJob)
	return fetchedJob, err
}
