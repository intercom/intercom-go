package intercom

import (
	"io/ioutil"
	"testing"
)

func TestAPIListSegments(t *testing.T) {
	http := TestSegmentHTTPClient{t: t, fixtureFilename: "fixtures/segments.json", expectedURI: "/segments"}
	api := SegmentAPI{httpClient: &http}
	segmentList, err := api.list()
	if err != nil {
		t.Fatalf(err.Error())
	}
	if segmentList.Segments[0].ID != "5443ac9b316c12246c000005" {
		t.Errorf("Segment list should start with segment 5443ac9b316c12246c000005, but had %s", segmentList.Segments[0].ID)
	}
	if segmentList.Segments[0].PersonType != "user" {
		t.Errorf("Segment list should generate person types from strings %s", segmentList.Segments[0].PersonType)
	}
}

func TestAPIFindSegment(t *testing.T) {
	http := TestSegmentHTTPClient{t: t, fixtureFilename: "fixtures/segment.json", expectedURI: "/segments/5443ac9b316c12246c000005"}
	api := SegmentAPI{httpClient: &http}
	segment, err := api.find("5443ac9b316c12246c000005")
	if err != nil {
		t.Fatalf(err.Error())
	}
	if segment.ID != "5443ac9b316c12246c000005" {
		t.Errorf("Segment should have ID 5443ac9b316c12246c000005, but had %s", segment.ID)
	}
	if segment.PersonType != "contact" {
		t.Errorf("Segment should generate person types from strings %s", segment.PersonType)
	}
}

type TestSegmentHTTPClient struct {
	TestHTTPClient
	t               *testing.T
	fixtureFilename string
	expectedURI     string
}

func (t TestSegmentHTTPClient) Get(uri string, params interface{}) ([]byte, error) {
	if uri != t.expectedURI {
		t.t.Errorf("Wrong endpoint called")
	}
	return ioutil.ReadFile(t.fixtureFilename)
}
