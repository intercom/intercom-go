package intercom

import (
	"io/ioutil"
	"testing"
)

func TestMessageAPISave(t *testing.T) {
	http := TestMessageHTTPClient{t: t, expectedURI: "/messages", fixtureFilename: "fixtures/message.json"}
	api := MessageAPI{httpClient: &http}
	message := NewUserMessage(User{}, "Hey, is the new thing in stock?")
	msg, err := api.save(&message)
	if err != nil {
		t.Error(err)
	}
	if msg.Body != "Hey, is the new thing in stock?" {
		t.Errorf("Message body was not set, was %s", msg.Body)
	}
	if msg.CreatedAt != 1401917202 {
		t.Errorf("Message CreatedAt was not set")
	}
	if msg.Template != PERSONAL_TEMPLATE {
		t.Errorf("Message template was not set, was %s", msg.Template)
	}
}

type TestMessageHTTPClient struct {
	TestHTTPClient
	t               *testing.T
	fixtureFilename string
	expectedURI     string
	lastQueryParams interface{}
}

func (t *TestMessageHTTPClient) Post(uri string, body interface{}) ([]byte, error) {
	if t.expectedURI != uri {
		t.t.Errorf("Wrong endpoint called")
	}
	return ioutil.ReadFile(t.fixtureFilename)
}
