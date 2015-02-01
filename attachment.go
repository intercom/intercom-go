package intercom

import "fmt"

type Attachment struct {
	Name        string `json:"name,omitempty"`
	URL         string `json:"url,omitempty"`
	ContentType string `json:"content_type,omitempty"`
	Filesize    int32  `json:"file_size,omitempty"`
}

func (a Attachment) String() string {
	return fmt.Sprintf("[intercom] attachment { name: %s, url: %s }", a.Name, a.URL)
}
