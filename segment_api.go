package intercom

import (
	"encoding/json"
	"fmt"

	"gopkg.in/intercom/intercom-go.v2/interfaces"
)

// SegmentRepository defines the interface for working with Segments through the API.
type SegmentRepository interface {
	list() (SegmentList, error)
	find(id string) (Segment, error)
}

// SegmentAPI implements SegmentRepository
type SegmentAPI struct {
	httpClient interfaces.HTTPClient
}

func (api SegmentAPI) list() (SegmentList, error) {
	segmentList := SegmentList{}
	data, err := api.httpClient.Get("/segments", nil)
	if err != nil {
		return segmentList, err
	}
	err = json.Unmarshal(data, &segmentList)
	return segmentList, err
}

func (api SegmentAPI) find(id string) (Segment, error) {
	segment := Segment{}
	data, err := api.httpClient.Get(fmt.Sprintf("/segments/%s", id), nil)
	if err != nil {
		return segment, err
	}
	err = json.Unmarshal(data, &segment)
	return segment, err
}
