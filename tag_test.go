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
	_, err := tagService.Save(&tag)
	if err != nil {
		t.Errorf("Failed to save tag: %s", err)
	}
}

func TestDeleteTag(t *testing.T) {
	tagService := TagService{Repository: TestTagAPI{t: t}}
	err := tagService.Delete("6")
	if err != nil {
		t.Errorf("Failed to delete tag: %s", err)
	}
}

func TestTaggingUsers(t *testing.T) {
	tagService := TagService{Repository: TestTagAPI{t: t}}
	taggingList := TaggingList{Name: "My Tag", Contacts: []Tagging{{ContactID: "245"}}}
	_, err := tagService.Tag(&taggingList)
	if err != nil {
		t.Errorf("Failed to tag user tag: %s", err)
	}
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
	if taggingList.Contacts[0].ContactID != "245" {
		t.t.Errorf("Tagging request expected to have UserID 245 but had %s", taggingList.Contacts[0].CompanyID)
	}
	if taggingList.Name != "My Tag" {
		t.t.Errorf("Tagging request expected to have Name My Tag but had %s", taggingList.Name)
	}
	return Tag{}, nil
}
