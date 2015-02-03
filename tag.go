package intercom

import "fmt"

type TagService struct {
	Repository TagRepository
}

type Tag struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type TagList struct {
	Tags []Tag `json:"tags,omitempty"`
}

func (t *TagService) List() (TagList, error) {
	return t.Repository.list()
}

func (t *TagService) Save(tag *Tag) (Tag, error) {
	return t.Repository.save(tag)
}

func (t *TagService) Delete(id string) error {
	return t.Repository.delete(id)
}

func (t *TagService) Tag(taggingList *TaggingList) (Tag, error) {
	return t.Repository.tag(taggingList)
}

func (t Tag) String() string {
	return fmt.Sprintf("[intercom] tag { id: %s name: %s }", t.ID, t.Name)
}
