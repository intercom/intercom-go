package intercom

import (
	"testing"
	"time"

	"gopkg.in/intercom/intercom-go.v2/interfaces"
)

func TestEventAPISave(t *testing.T) {
	http := TestEventHTTPClient{t: t, expectedURI: "/events"}
	api := EventAPI{httpClient: &http}
	event := Event{UserID: "27", CreatedAt: int64(time.Now().Unix()), EventName: "govent"}
	api.save(&event)
}

func TestEventAPISaveFail(t *testing.T) {
	http := TestEventHTTPClient{t: t, expectedURI: "/events", shouldFail: true}
	api := EventAPI{httpClient: &http}
	event := Event{UserID: "444", CreatedAt: int64(time.Now().Unix()), EventName: "govent"}
	err := api.save(&event)
	if herr, ok := err.(interfaces.HTTPError); ok && herr.Code != "not_found" {
		t.Errorf("Error not returned")
	}
}

type TestEventHTTPClient struct {
	TestHTTPClient
	t           *testing.T
	expectedURI string
	shouldFail  bool
}

func (t TestEventHTTPClient) Post(uri string, event interface{}) ([]byte, error) {
	if uri != "/events" {
		t.t.Errorf("Wrong endpoint called")
	}
	if t.shouldFail {
		err := interfaces.HTTPError{StatusCode: 404, Code: "not_found", Message: "User Not Found"}
		return nil, err
	}
	return nil, nil
}
