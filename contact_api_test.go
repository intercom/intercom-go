package intercom

import (
	"os"
	"testing"
)

func TestContactAPIFind(t *testing.T) {
	http := TestContactHTTPClient{fixtureFilename: "fixtures/contact.json", expectedURI: "/contacts/54c42e7ea7a765fa7", t: t}
	api := ContactAPI{httpClient: &http}
	contact, err := api.find(ContactIdentifiers{ID: "54c42e7ea7a765fa7"})
	if err != nil {
		t.Errorf("Error parsing fixture %s", err)
	}
	if contact.ID != "54c42e7ea7a765fa7" {
		t.Errorf("ID was %s, expected 54c42e7ea7a765fa7", contact.ID)
	}
	if contact.Phone != "+1234567890" {
		t.Errorf("Phone was %s, expected +1234567890", contact.Phone)
	}
	if contact.ExternalID != "123" {
		t.Errorf("ExternalID was %s, expected 123", contact.ExternalID)
	}
}

func TestContactAPIListDefault(t *testing.T) {
	http := TestContactHTTPClient{fixtureFilename: "fixtures/contacts.json", expectedURI: "/contacts", t: t}
	api := ContactAPI{httpClient: &http}
	contactList, _ := api.list(contactListParams{})
	contacts := contactList.Contacts
	if contacts[0].ID != "54c42e7ea7a765fa7" {
		t.Errorf("ID was %s, expected 54c42e7ea7a765fa7", contacts[0].ID)
	}
	pages := contactList.Pages
	if pages.Page != 1 {
		t.Errorf("Page was %d, expected 1", pages.Page)
	}
}

func TestContactAPIListByEmail(t *testing.T) {
	http := TestContactHTTPClient{fixtureFilename: "fixtures/contacts.json", expectedURI: "/contacts", t: t}
	api := ContactAPI{httpClient: &http}
	contactList, _ := api.list(contactListParams{Email: "mycontact@example.io"})
	contacts := contactList.Contacts
	if contacts[0].ID != "54c42e7ea7a765fa7" {
		t.Errorf("ID was %s, expected 54c42e7ea7a765fa7", contacts[0].ID)
	}
	if clParams, ok := http.lastQueryParams.(contactListParams); !ok || clParams.Email != "mycontact@example.io" {
		t.Errorf("Email expected to be mycontact@example.io, but was %s", clParams.Email)
	}
	pages := contactList.Pages
	if pages.Page != 1 {
		t.Errorf("Page was %d, expected 1", pages.Page)
	}
}

func TestContactAPICreate(t *testing.T) {
	http := TestContactHTTPClient{fixtureFilename: "fixtures/contact.json", expectedURI: "/contacts", t: t}
	api := ContactAPI{httpClient: &http}
	contact := &Contact{Email: "mycontact@example.io"}
	_, err := api.create(contact)
	if err != nil {
		t.Errorf("Failed to create contact: %s", err)
	}
}

func TestContactAPIUpdate(t *testing.T) {
	http := TestContactHTTPClient{fixtureFilename: "fixtures/contact.json", expectedURI: "/contacts", t: t}
	api := ContactAPI{httpClient: &http}
	contact := &Contact{ExternalID: "123", Email: "mycontact@example.io"}
	_, err := api.update(contact)
	if err != nil {
		t.Errorf("Failed to update contact %s", err)
	}
}

func TestContactAPIDelete(t *testing.T) {
	http := TestContactHTTPClient{fixtureFilename: "fixtures/contact.json", expectedURI: "/contacts/b123d", t: t}
	api := ContactAPI{httpClient: &http}
	contact := &Contact{ID: "b123d"}
	returned, _ := api.delete(contact.ID)
	if returned.ExternalID != "123" {
		t.Errorf("Expected ExternalID %s, got %s", "123", returned.ExternalID)
	}
}

type TestContactHTTPClient struct {
	TestHTTPClient
	t               *testing.T
	fixtureFilename string
	expectedURI     string
	lastQueryParams interface{}
}

func (t *TestContactHTTPClient) Get(uri string, queryParams interface{}) ([]byte, error) {
	if t.expectedURI != uri {
		t.t.Errorf("URI was %s, expected %s", uri, t.expectedURI)
	}
	t.lastQueryParams = queryParams
	return os.ReadFile(t.fixtureFilename)
}

func (t *TestContactHTTPClient) Post(uri string, body interface{}) ([]byte, error) {
	if t.expectedURI != uri {
		t.t.Errorf("Wrong endpoint called")
	}
	return os.ReadFile(t.fixtureFilename)
}

func (t *TestContactHTTPClient) Delete(uri string, queryParams interface{}) ([]byte, error) {
	if t.expectedURI != uri {
		t.t.Errorf("URI was %s, expected %s", uri, t.expectedURI)
	}
	return os.ReadFile(t.fixtureFilename)
}
