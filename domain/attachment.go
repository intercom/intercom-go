package domain

import "fmt"

type Attachment struct {
	Name        string
	URL         string
	ContentType string
	Filesize    int32
}

func (a Attachment) String() string {
	return fmt.Sprintf("[intercom] attachment { name: %s, url: %s }", a.Name, a.URL)
}
