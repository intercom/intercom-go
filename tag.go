package intercom

type Tag struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type TagList struct {
	Tags []Tag `json:"tags,omitempty"`
}
