package intercom

import (
	"encoding/json"
	"fmt"

	"github.com/intercom/intercom-go/interfaces"
)

type SegmentRepository interface {
	list() (SegmentList, error)
	find(id string) (Segment, error)
}

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
