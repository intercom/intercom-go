package interfaces

import (
	"testing"
	"time"

	"github.com/intercom/intercom-go/domain"
)

func TestSave(t *testing.T) {
	http := TestEventHTTPClient{t: t, expectedURI: "/events"}
	api := EventAPI{httpClient: &http}
	event := domain.Event{UserID: "27", CreatedAt: int32(time.Now().Unix()), EventName: "govent"}
	api.Save(event)
}

func TestSaveFail(t *testing.T) {
	http := TestEventHTTPClient{t: t, expectedURI: "/events", shouldFail: true}
	api := EventAPI{httpClient: &http}
	event := domain.Event{UserID: "444", CreatedAt: int32(time.Now().Unix()), EventName: "govent"}
	err := api.Save(event)
	if herr, ok := err.(HTTPError); ok && herr.Code != "not_found" {
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
		err := HTTPError{StatusCode: 404, Code: "not_found", Message: "User Not Found"}
		return nil, err
	}
	return nil, nil
}
