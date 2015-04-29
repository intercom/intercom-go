package intercom

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/intercom/intercom-go/interfaces"
)

// ContactRepository defines the interface for working with Contacts through the API.
type ContactRepository interface {
	find(UserIdentifiers) (Contact, error)
	list(contactListParams) (ContactList, error)
	create(*Contact) (Contact, error)
	update(*Contact) (Contact, error)
}

// ContactAPI implements ContactRepository
type ContactAPI struct {
	httpClient interfaces.HTTPClient
}

func (api ContactAPI) find(params UserIdentifiers) (Contact, error) {
	contact := Contact{}
	data, err := api.getClientForFind(params)
	if err != nil {
		return contact, err
	}
	err = json.Unmarshal(data, &contact)
	return contact, err
}

func (api ContactAPI) getClientForFind(params UserIdentifiers) ([]byte, error) {
	switch {
	case params.ID != "":
		return api.httpClient.Get(fmt.Sprintf("/contacts/%s", params.ID), nil)
	case params.UserID != "":
		return api.httpClient.Get("/contacts", params)
	}
	return nil, errors.New("Missing Contact Identifier")
}

func (api ContactAPI) list(params contactListParams) (ContactList, error) {
	contactList := ContactList{}
	data, err := api.httpClient.Get("/contacts", params)
	if err != nil {
		return contactList, err
	}
	err = json.Unmarshal(data, &contactList)
	return contactList, err
}

func (api ContactAPI) create(contact *Contact) (Contact, error) {
	requestContact := api.buildRequestContact(contact)
	return unmarshalToContact(api.httpClient.Post("/contacts", &requestContact))
}

func (api ContactAPI) update(contact *Contact) (Contact, error) {
	requestContact := api.buildRequestContact(contact)
	return unmarshalToContact(api.httpClient.Patch("/contacts", &requestContact))
}

func unmarshalToContact(data []byte, err error) (Contact, error) {
	savedContact := Contact{}
	if err != nil {
		return savedContact, err
	}
	err = json.Unmarshal(data, &savedContact)
	return savedContact, err
}

func (api ContactAPI) buildRequestContact(contact *Contact) requestUser {
	return requestUser{
		ID:                     contact.ID,
		Email:                  contact.Email,
		UserID:                 contact.UserID,
		Name:                   contact.Name,
		LastRequestAt:          contact.LastRequestAt,
		LastSeenIP:             contact.LastSeenIP,
		UnsubscribedFromEmails: contact.UnsubscribedFromEmails,
		CustomAttributes:       contact.CustomAttributes,
		UpdateLastRequestAt:    contact.UpdateLastRequestAt,
		NewSession:             contact.NewSession,
	}
}
