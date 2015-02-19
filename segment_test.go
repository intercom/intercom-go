package intercom

import "testing"

func TestListSegments(t *testing.T) {
	segmentList, _ := (&SegmentService{Repository: TestSegmentAPI{t: t}}).List()
	segments := segmentList.Segments
	if segments[0].ID != "de412cad4" {
		t.Errorf("Got segment with ID %s, expected de412cad4", segments[0].ID)
	}
}

func TestFindSegment(t *testing.T) {
	segment, _ := (&SegmentService{Repository: TestSegmentAPI{t: t}}).Find("de412cad4")
	if segment.ID != "de412cad4" {
		t.Errorf("Got segment with ID %s, expected de412cad4", segment.ID)
	}
}

type TestSegmentAPI struct {
	t *testing.T
}

func (t TestSegmentAPI) list() (SegmentList, error) {
	return SegmentList{Segments: []Segment{Segment{ID: "de412cad4", Name: "My Tag"}}}, nil
}

func (t TestSegmentAPI) find(id string) (Segment, error) {
	return Segment{ID: id}, nil
}
