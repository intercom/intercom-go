package intercom

import "fmt"

type SegmentService struct {
	Repository SegmentRepository
}

type Segment struct {
	ID        string `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	CreatedAt int32  `json:"created_at,omitempty"`
	UpdatedAt int32  `json:"updated_at,omitempty"`
}

type SegmentList struct {
	Segments []Segment `json:"segments,omitempty"`
}

func (t *SegmentService) List() (SegmentList, error) {
	return t.Repository.list()
}

func (t *SegmentService) Find(id string) (Segment, error) {
	return t.Repository.find(id)
}

func (s Segment) String() string {
	return fmt.Sprintf("[intercom] segment { id: %s }", s.ID)
}
