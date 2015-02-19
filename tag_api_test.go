package intercom

import (
	"io/ioutil"
	"testing"
)

type TestTagHTTPClient struct {
	TestHTTPClient
	t               *testing.T
	fixtureFilename string
	expectedURI     string
}

func TestAPIListTag(t *testing.T) {
	http := TestTagHTTPClient{t: t, fixtureFilename: "fixtures/tags.json", expectedURI: "/tags"}
	api := TagAPI{httpClient: &http}
	tagList, _ := api.list()
	if tagList.Tags[0].ID != "51313" {
		t.Errorf("Tag list should start with tag 51313, but had %s", tagList.Tags[0].ID)
	}
}

func TestAPITagSave(t *testing.T) {
	http := TestTagHTTPClient{t: t, fixtureFilename: "fixtures/tag.json", expectedURI: "/tags"}
	api := TagAPI{httpClient: &http}
	tag := Tag{ID: "60218", Name: "My Tag"}
	savedTag, _ := api.save(&tag)
	if savedTag.ID != "60218" {
		t.Errorf("Expected saved tag with ID 60218, got %s", savedTag.ID)
	}
}

func TestAPITagDelete(t *testing.T) {
	http := TestTagHTTPClient{t: t, expectedURI: "/tags/6"}
	api := TagAPI{httpClient: &http}
	api.delete("6")
}

func TestAPITagTagging(t *testing.T) {
	http := TestTagHTTPClient{t: t, fixtureFilename: "fixtures/tag.json", expectedURI: "/tags"}
	api := TagAPI{httpClient: &http}
	taggingList := TaggingList{Name: "My Tag", Users: []Tagging{Tagging{UserID: "2345"}}}
	savedTag, _ := api.tag(&taggingList)
	if savedTag.ID != "60218" {
		t.Errorf("Expected saved tag with ID 60218, got %s", savedTag.ID)
	}
}

func (t TestTagHTTPClient) Get(uri string, params interface{}) ([]byte, error) {
	if uri != t.expectedURI {
		t.t.Errorf("Wrong endpoint called")
	}
	return ioutil.ReadFile(t.fixtureFilename)
}

func (t TestTagHTTPClient) Post(uri string, body interface{}) ([]byte, error) {
	if uri != t.expectedURI {
		t.t.Errorf("Wrong endpoint called")
	}
	return ioutil.ReadFile(t.fixtureFilename)
}

func (t TestTagHTTPClient) Delete(uri string, body interface{}) ([]byte, error) {
	if uri != t.expectedURI {
		t.t.Errorf("Wrong endpoint called")
	}
	return nil, nil
}
