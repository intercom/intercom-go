package intercom

import (
	"testing"
)

func TestContactFindByID(t *testing.T) {
	contact, _ := (&ContactService{Repository: TestContactAPI{t: t}}).FindByID("46adad3f09126dca")
	if contact.ID != "46adad3f09126dca" {
		t.Errorf("Contact not found")
	}
}

func TestContactFindByUserID(t *testing.T) {
	contact, _ := (&ContactService{Repository: TestContactAPI{t: t}}).FindByUserID("134d")
	if contact.UserID != "134d" {
		t.Errorf("Contact not found")
	}
}

func TestContactList(t *testing.T) {
	contactList, _ := (&ContactService{Repository: TestContactAPI{t: t}}).ListByEmail("jamie@example.io", PageParams{})
	contacts := contactList.Contacts
	if contacts[0].ID != "46adad3f09126dca" {
		t.Errorf("Contact not listed")
	}
}

func TestContactListEmail(t *testing.T) {
	contactList, _ := (&ContactService{Repository: TestContactAPI{t: t}}).List(PageParams{})
	contacts := contactList.Contacts
	if contacts[0].ID != "46adad3f09126dca" {
		t.Errorf("Contact not listed")
	}
}

type TestContactAPI struct {
	t *testing.T
}

func (t TestContactAPI) find(params UserIdentifiers) (Contact, error) {
	return Contact{ID: params.ID, Email: params.Email, UserID: params.UserID}, nil
}

func (t TestContactAPI) list(params contactListParams) (ContactList, error) {
	return ContactList{Contacts: []Contact{Contact{ID: "46adad3f09126dca", Email: "jamie@example.io", UserID: "aa123"}}}, nil
}
