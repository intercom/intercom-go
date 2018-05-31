package intercom

import (
	"encoding/json"
	"fmt"

	"gopkg.in/intercom/intercom-go.v2/interfaces"
)

// TagRepository defines the interface for working with Tags through the API.
type TagRepository interface {
	list() (TagList, error)
	save(tag *Tag) (Tag, error)
	delete(id string) error
	tag(tagList *TaggingList) (Tag, error)
}

// TagAPI implements TagRepository
type TagAPI struct {
	httpClient interfaces.HTTPClient
}

func (api TagAPI) list() (TagList, error) {
	tagList := TagList{}
	data, err := api.httpClient.Get("/tags", nil)
	if err != nil {
		return tagList, err
	}
	err = json.Unmarshal(data, &tagList)
	return tagList, err
}

func (api TagAPI) save(tag *Tag) (Tag, error) {
	savedTag := Tag{}
	data, err := api.httpClient.Post("/tags", tag)
	if err != nil {
		return savedTag, err
	}
	err = json.Unmarshal(data, &savedTag)
	return savedTag, err
}

func (api TagAPI) delete(id string) error {
	_, err := api.httpClient.Delete(fmt.Sprintf("/tags/%s", id), nil)
	return err
}

func (api TagAPI) tag(taggingList *TaggingList) (Tag, error) {
	savedTag := Tag{}
	data, err := api.httpClient.Post("/tags", taggingList)
	if err != nil {
		return savedTag, err
	}
	err = json.Unmarshal(data, &savedTag)
	return savedTag, err
}
