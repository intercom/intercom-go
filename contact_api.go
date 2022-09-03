package intercom

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/cameronnewman/intercom-go/interfaces"
)

// ContactRepository defines the interface for working with Contacts through the API.
type ContactRepository interface {
	find(identifiers ContactIdentifiers) (Contact, error)
	list(contactListParams) (ContactList, error)
	scroll(scrollParam string) (ContactList, error)
	create(*Contact) (Contact, error)
	update(*Contact) (Contact, error)
	delete(id string) (Contact, error)
}

// ContactAPI implements ContactRepository
type ContactAPI struct {
	httpClient interfaces.HTTPClient
}

func (api ContactAPI) find(params ContactIdentifiers) (Contact, error) {
	return unmarshalToContact(api.getClientForFind(params))
}

func (api ContactAPI) getClientForFind(params ContactIdentifiers) ([]byte, error) {
	switch {
	case params.ID != "":
		return api.httpClient.Get(fmt.Sprintf("/contacts/%s", params.ID), nil)
	case params.ExternalID != "":
		return api.httpClient.Get("/contacts", params)
	}
	return nil, errors.New("missing Contact Identifier")
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

func (api ContactAPI) scroll(scrollParam string) (ContactList, error) {
	contactList := ContactList{}
	params := scrollParams{ScrollParam: scrollParam}
	data, err := api.httpClient.Get("/contacts/scroll", params)
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
	return unmarshalToContact(api.httpClient.Post("/contacts", &requestContact))
}

func (api ContactAPI) delete(id string) (Contact, error) {
	contact := Contact{}
	data, err := api.httpClient.Delete(fmt.Sprintf("/contacts/%s", id), nil)
	if err != nil {
		return contact, err
	}
	err = json.Unmarshal(data, &contact)
	return contact, err
}

func unmarshalToContact(data []byte, err error) (Contact, error) {
	savedContact := Contact{}
	if err != nil {
		return savedContact, err
	}
	err = json.Unmarshal(data, &savedContact)
	return savedContact, err
}

func (api ContactAPI) buildRequestContact(contact *Contact) requestContact {
	return requestContact{
		ID:                     contact.ID,
		Email:                  contact.Email,
		Phone:                  contact.Phone,
		ExternalID:             contact.ExternalID,
		Name:                   contact.Name,
		LastRequestAt:          contact.LastRequestAt,
		LastSeenIP:             contact.LastSeenIP,
		UnsubscribedFromEmails: contact.UnsubscribedFromEmails,
		Companies:              api.getCompaniesToSendFromContact(contact),
		CustomAttributes:       contact.CustomAttributes,
		UpdateLastRequestAt:    contact.UpdateLastRequestAt,
		NewSession:             contact.NewSession,
	}
}

func (api ContactAPI) getCompaniesToSendFromContact(contact *Contact) []ContactCompany {
	if contact.Companies == nil {
		return []ContactCompany{}
	}
	return requestContact{}.MakeUserCompaniesFromCompanies(contact.Companies.Companies)
}

type requestContact struct {
	ID                     string                 `json:"id,omitempty"`
	Email                  string                 `json:"email,omitempty"`
	Phone                  string                 `json:"phone,omitempty"`
	ExternalID             string                 `json:"external_id,omitempty"`
	Name                   string                 `json:"name,omitempty"`
	SignedUpAt             int64                  `json:"signed_up_at,omitempty"`
	RemoteCreatedAt        int64                  `json:"remote_created_at,omitempty"`
	LastRequestAt          int64                  `json:"last_request_at,omitempty"`
	LastSeenIP             string                 `json:"last_seen_ip,omitempty"`
	UnsubscribedFromEmails *bool                  `json:"unsubscribed_from_emails,omitempty"`
	Companies              []ContactCompany       `json:"companies,omitempty"`
	CustomAttributes       map[string]interface{} `json:"custom_attributes,omitempty"`
	UpdateLastRequestAt    *bool                  `json:"update_last_request_at,omitempty"`
	NewSession             *bool                  `json:"new_session,omitempty"`
	LastSeenUserAgent      string                 `json:"last_seen_user_agent,omitempty"`
}

func (rum requestContact) MakeUserCompaniesFromCompanies(companies []Company) []ContactCompany {
	contactCompanies := make([]ContactCompany, len(companies))
	for i := 0; i < len(companies); i++ {
		contactCompanies[i] = ContactCompany{
			CompanyID: companies[i].CompanyID,
			Name:      companies[i].Name,
			Remove:    companies[i].Remove,
		}
	}
	return contactCompanies
}

// ContactCompany is the Company a Contact belongs to
// used to update Companies on a Contact.
type ContactCompany struct {
	CompanyID string `json:"company_id,omitempty"`
	Name      string `json:"name,omitempty"`
	Remove    *bool  `json:"remove,omitempty"`
}
