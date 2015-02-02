package intercom

type Segment struct {
	ID string `json:"id,omitempty"`
}

type SegmentList struct {
	Segments []Segment `json:"segments,omitempty"`
}
