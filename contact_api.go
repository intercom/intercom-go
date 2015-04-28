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
