package intercom

import (
	"testing"

	"github.com/google/uuid"
)

func TestContactFindByID(t *testing.T) {
	contact, _ := (&ContactService{Repository: TestContactAPI{t: t}}).FindByID("46adad3f09126dca")
	if contact.ID != "46adad3f09126dca" {
		t.Errorf("Contact not found")
	}
}

func TestContactFindByUserID(t *testing.T) {
	contact, _ := (&ContactService{Repository: TestContactAPI{t: t}}).FindByUserID("134d")
	if contact.ExternalID != "134d" {
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

func TestContactCreate(t *testing.T) {
	contactService := ContactService{Repository: TestContactAPI{t: t}}
	contact := Contact{Email: "some@email.com"}
	c, _ := contactService.Create(&contact)
	if c.Email != contact.Email {
		t.Errorf("expected returned contact to have email %s, got %s", contact.Email, c.Email)
	}
}

func TestContactUpdate(t *testing.T) {
	contactService := ContactService{Repository: TestContactAPI{t: t}}
	contact := Contact{Email: "some@email.com"}
	c, _ := contactService.Update(&contact)
	if c.Email != contact.Email {
		t.Errorf("expected returned contact to have email %s, got %s", contact.Email, c.Email)
	}
}

func TestContactDelete(t *testing.T) {
	contactService := ContactService{Repository: TestContactAPI{t: t}}
	contact := Contact{ExternalID: "aaaa", Email: "some@email.com"}
	_, err := contactService.Delete(&contact)
	if err != nil {
		t.Errorf("Failed to delete contact: %s", err)
	}
}

func TestContactMessageAddress(t *testing.T) {
	contact := Contact{ExternalID: "aaaa", Email: "some@email.com"}
	address := contact.MessageAddress()
	if address.ID != "" {
		t.Errorf("Contact address had ID")
	}
	if address.Type != "contact" {
		t.Errorf("Contact address was not of type contact, was %s", address.Type)
	}
	if address.Email != "some@email.com" {
		t.Errorf("Contact address had wrong Email")
	}
	if address.ExternalID != "aaaa" {
		t.Errorf("Contact address had wrong UserID")
	}
}

type TestContactAPI struct {
	t *testing.T
}

func (t TestContactAPI) find(params ContactIdentifiers) (Contact, error) {
	return Contact{ID: params.ID, Email: params.Email, ExternalID: params.ExternalID}, nil
}

func (t TestContactAPI) list(params contactListParams) (ContactList, error) {
	return ContactList{Contacts: []Contact{Contact{ID: "46adad3f09126dca", Email: "jamie@example.io", ExternalID: "aa123"}}}, nil
}

func (t TestContactAPI) scroll(scrollParam string) (ContactList, error) {
	return ContactList{Contacts: []Contact{Contact{ID: "46adad3f09126dca", Email: "jamie@example.io", ExternalID: "aa123"}}}, nil
}

func (t TestContactAPI) create(c *Contact) (Contact, error) {
	return Contact{ID: c.ID, Email: c.Email, ExternalID: uuid.Must(uuid.NewRandom()).String()}, nil
}

func (t TestContactAPI) update(c *Contact) (Contact, error) {
	return Contact{ID: c.ID, Email: c.Email, ExternalID: c.ExternalID}, nil
}

func (t TestContactAPI) delete(id string) (Contact, error) {
	return Contact{ID: id}, nil
}
