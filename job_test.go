package intercom

import "testing"

func TestNewJob(t *testing.T) {
	repo := &TestJobRepository{t: t}
	repo.f = func(job *JobRequest) {
		if job.Items[0].Method != JOB_POST.String() {
			repo.t.Errorf("Wrong job method")
		}
		u := job.Items[0].Data.(*User)
		if u.Email != "foo@bar.com" {
			repo.t.Errorf("Wrong user email")
		}
	}
	user := User{Email: "foo@bar.com"}
	js := JobService{Repository: repo}
	_, err := js.NewUserJob(NewUserJobItem(&user, JOB_POST))
	if err != nil {
		t.Errorf("Failed to create a new job: %s", err)
	}
}

func TestAppendJob(t *testing.T) {
	repo := &TestJobRepository{t: t}
	js := JobService{Repository: repo}
	newJob, _ := js.NewUserJob()

	repo.f = func(job *JobRequest) {
		if job.Items[0].Method != JOB_POST.String() {
			repo.t.Errorf("Wrong job method")
		}
		u := job.Items[0].Data.(*User)
		if u.Email != "foo@bar.com" {
			repo.t.Errorf("Wrong user email")
		}
	}
	user := User{Email: "foo@bar.com"}

	_, err := js.AppendUsers(newJob.ID, NewUserJobItem(&user, JOB_POST))
	if err != nil {
		t.Errorf("Failed to append user to existing job: %s", err)
	}
}

type TestJobRepository struct {
	t *testing.T
	f func(job *JobRequest)
}

func (api *TestJobRepository) save(job *JobRequest) (JobResponse, error) {
	if api.f != nil {
		api.f(job)
	}
	return JobResponse{}, nil
}

func (api *TestJobRepository) find(id string) (JobResponse, error) {
	return JobResponse{}, nil
}
