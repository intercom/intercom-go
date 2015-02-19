package intercom

import "testing"

func TestListTags(t *testing.T) {
	tagList, _ := (&TagService{Repository: TestTagAPI{t: t}}).List()
	tags := tagList.Tags
	if tags[0].ID != "24" {
		t.Errorf("Got tag with ID %s, expected 24", tags[0].ID)
	}
}

func TestSaveTag(t *testing.T) {
	tagService := TagService{Repository: TestTagAPI{t: t}}
	tag := Tag{ID: "24", Name: "My Tag"}
	tagService.Save(&tag)
}

func TestDeleteTag(t *testing.T) {
	tagService := TagService{Repository: TestTagAPI{t: t}}
	tagService.Delete("6")
}

func TestTaggingUsers(t *testing.T) {
	tagService := TagService{Repository: TestTagAPI{t: t}}
	taggingList := TaggingList{Name: "My Tag", Users: []Tagging{Tagging{UserID: "245"}}}
	tagService.Tag(&taggingList)
}

type TestTagAPI struct {
	t *testing.T
}

func (t TestTagAPI) list() (TagList, error) {
	return TagList{Tags: []Tag{Tag{ID: "24", Name: "My Tag"}}}, nil
}

func (t TestTagAPI) save(tag *Tag) (Tag, error) {
	if tag.ID != "24" {
		t.t.Errorf("Saved tag expected to have ID 24 but has %s", tag.ID)
	}
	return *tag, nil
}

func (t TestTagAPI) delete(id string) error {
	if id != "6" {
		t.t.Errorf("Delete tag request expected to have ID 6, but has %s", id)
	}
	return nil
}

func (t TestTagAPI) tag(taggingList *TaggingList) (Tag, error) {
	if taggingList.Users[0].UserID != "245" {
		t.t.Errorf("Tagging request expected to have UserID 245 but had %s", taggingList.Users[0].UserID)
	}
	if taggingList.Name != "My Tag" {
		t.t.Errorf("Tagging request expected to have Name My Tag but had %s", taggingList.Name)
	}
	return Tag{}, nil
}
