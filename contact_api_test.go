package intercom

import "testing"

func TestContactAPIFind(t *testing.T) {
	http := TestUserHTTPClient{fixtureFilename: "fixtures/contact.json", expectedURI: "/contacts/54c42e7ea7a765fa7", t: t}
	api := ContactAPI{httpClient: &http}
	contact, err := api.find(UserIdentifiers{ID: "54c42e7ea7a765fa7"})
	if err != nil {
		t.Errorf("Error parsing fixture %s", err)
	}
	if contact.ID != "54c42e7ea7a765fa7" {
		t.Errorf("ID was %s, expected 54c42e7ea7a765fa7", contact.ID)
	}
	if contact.Phone != "+1234567890" {
		t.Errorf("Phone was %s, expected +1234567890", contact.Phone)
	}
	if contact.UserID != "123" {
		t.Errorf("UserID was %s, expected 123", contact.UserID)
	}
}

func TestContactAPIListDefault(t *testing.T) {
	http := TestUserHTTPClient{fixtureFilename: "fixtures/contacts.json", expectedURI: "/contacts", t: t}
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
	http := TestUserHTTPClient{fixtureFilename: "fixtures/contacts.json", expectedURI: "/contacts", t: t}
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
	http := TestUserHTTPClient{fixtureFilename: "fixtures/contact.json", expectedURI: "/contacts", t: t}
	api := ContactAPI{httpClient: &http}
	contact := &Contact{Email: "mycontact@example.io"}
	api.create(contact)
}

func TestContactAPIUpdate(t *testing.T) {
	http := TestUserHTTPClient{fixtureFilename: "fixtures/contact.json", expectedURI: "/contacts", t: t}
	api := ContactAPI{httpClient: &http}
	contact := &Contact{UserID: "123", Email: "mycontact@example.io"}
	api.update(contact)
}

func TestContactAPIConvert(t *testing.T) {
	http := TestUserHTTPClient{fixtureFilename: "fixtures/user.json", expectedURI: "/contacts/convert", t: t}
	api := ContactAPI{httpClient: &http}
	contact := &Contact{UserID: "abc", Email: "mycontact@example.io"}
	user := &User{UserID: "123"}
	returned, _ := api.convert(contact, user)
	if returned.UserID != "123" {
		t.Errorf("Expected UserID %s, got %s", "123", returned.UserID)
	}
}

func TestContactAPIDelete(t *testing.T) {
	http := TestUserHTTPClient{fixtureFilename: "fixtures/contact.json", expectedURI: "/contacts/b123d", t: t}
	api := ContactAPI{httpClient: &http}
	contact := &Contact{ID: "b123d"}
	returned, _ := api.delete(contact.ID)
	if returned.UserID != "123" {
		t.Errorf("Expected UserID %s, got %s", "123", returned.UserID)
	}
}
