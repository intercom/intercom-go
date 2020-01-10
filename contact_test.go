package intercom

import (
	"testing"

	"github.com/pborman/uuid"
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

func TestContactConvert(t *testing.T) {
	contactService := ContactService{Repository: TestContactAPI{t: t}}
	contact := Contact{UserID: "aaaa", Email: "some@email.com"}
	user := User{ID: "abc13", UserID: "c135"}
	u, _ := contactService.Convert(&contact, &user)
	if u.Email != contact.Email {
		t.Errorf("expected returned user to have email %s, got %s", contact.Email, u.Email)
	}
	if u.UserID != user.UserID {
		t.Errorf("expected returned user to have user id %s, got %s", user.UserID, u.UserID)
	}
}

func TestContactDelete(t *testing.T) {
	contactService := ContactService{Repository: TestContactAPI{t: t}}
	contact := Contact{UserID: "aaaa", Email: "some@email.com"}
	contactService.Delete(&contact)
}

func TestContactMessageAddress(t *testing.T) {
	contact := Contact{UserID: "aaaa", Email: "some@email.com"}
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
	if address.UserID != "aaaa" {
		t.Errorf("Contact address had wrong UserID")
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

func (t TestContactAPI) scroll(scrollParam string) (ContactList, error) {
	return ContactList{Contacts: []Contact{Contact{ID: "46adad3f09126dca", Email: "jamie@example.io", UserID: "aa123"}}}, nil
}

func (t TestContactAPI) create(c *Contact) (Contact, error) {
	return Contact{ID: c.ID, Email: c.Email, UserID: uuid.New()}, nil
}

func (t TestContactAPI) update(c *Contact) (Contact, error) {
	return Contact{ID: c.ID, Email: c.Email, UserID: c.UserID}, nil
}

func (t TestContactAPI) convert(c *Contact, u *User) (User, error) {
	return User{ID: u.ID, Email: c.Email, UserID: u.UserID}, nil
}

func (t TestContactAPI) delete(id string) (Contact, error) {
	return Contact{ID: id}, nil
}
