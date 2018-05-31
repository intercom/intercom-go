package intercom

import "fmt"

// ContactService handles interactions with the API through a ContactRepository.
type ContactService struct {
	Repository ContactRepository
}

// ContactList holds a list of Contacts and paging information
type ContactList struct {
	Pages    PageParams
	Contacts []Contact
	ScrollParam string `json:"scroll_param,omitempty"`
}

// Contact represents a Contact within Intercom.
// Not all of the fields are writeable to the API, non-writeable fields are
// stripped out from the request. Please see the API documentation for details.
type Contact struct {
	ID                     string                 `json:"id,omitempty"`
	Email                  string                 `json:"email,omitempty"`
	Phone                  string                 `json:"phone,omitempty"`
	UserID                 string                 `json:"user_id,omitempty"`
	Name                   string                 `json:"name,omitempty"`
	Avatar                 *UserAvatar            `json:"avatar,omitempty"`
	LocationData           *LocationData          `json:"location_data,omitempty"`
	LastRequestAt          int64                  `json:"last_request_at,omitempty"`
	CreatedAt              int64                  `json:"created_at,omitempty"`
	UpdatedAt              int64                  `json:"updated_at,omitempty"`
	SessionCount           int64                  `json:"session_count,omitempty"`
	LastSeenIP             string                 `json:"last_seen_ip,omitempty"`
	SocialProfiles         *SocialProfileList     `json:"social_profiles,omitempty"`
	UnsubscribedFromEmails *bool                  `json:"unsubscribed_from_emails,omitempty"`
	UserAgentData          string                 `json:"user_agent_data,omitempty"`
	Tags                   *TagList               `json:"tags,omitempty"`
	Segments               *SegmentList           `json:"segments,omitempty"`
	Companies              *CompanyList           `json:"companies,omitempty"`
	CustomAttributes       map[string]interface{} `json:"custom_attributes,omitempty"`
	UpdateLastRequestAt    *bool                  `json:"update_last_request_at,omitempty"`
	NewSession             *bool                  `json:"new_session,omitempty"`
}

type contactListParams struct {
	PageParams
	SegmentID string `url:"segment_id,omitempty"`
	TagID     string `url:"tag_id,omitempty"`
	Email     string `url:"email,omitempty"`
}

// FindByID looks up a Contact by their Intercom ID.
func (c *ContactService) FindByID(id string) (Contact, error) {
	return c.findWithIdentifiers(UserIdentifiers{ID: id})
}

// FindByUserID looks up a Contact by their UserID (automatically generated server side).
func (c *ContactService) FindByUserID(userID string) (Contact, error) {
	return c.findWithIdentifiers(UserIdentifiers{UserID: userID})
}

func (c *ContactService) findWithIdentifiers(identifiers UserIdentifiers) (Contact, error) {
	return c.Repository.find(identifiers)
}

// List all Contacts for App.
func (c *ContactService) List(params PageParams) (ContactList, error) {
	return c.Repository.list(contactListParams{PageParams: params})
}

// List all Contacts for App via Scroll API
func (c *ContactService) Scroll(scrollParam string) (ContactList, error) {
       return c.Repository.scroll(scrollParam)
}

// ListByEmail looks up a list of Contacts by their Email.
func (c *ContactService) ListByEmail(email string, params PageParams) (ContactList, error) {
	return c.Repository.list(contactListParams{PageParams: params, Email: email})
}

// List Contacts by Segment.
func (c *ContactService) ListBySegment(segmentID string, params PageParams) (ContactList, error) {
	return c.Repository.list(contactListParams{PageParams: params, SegmentID: segmentID})
}

// List Contacts By Tag.
func (c *ContactService) ListByTag(tagID string, params PageParams) (ContactList, error) {
	return c.Repository.list(contactListParams{PageParams: params, TagID: tagID})
}

// Create Contact
func (c *ContactService) Create(contact *Contact) (Contact, error) {
	return c.Repository.create(contact)
}

// Update Contact
func (c *ContactService) Update(contact *Contact) (Contact, error) {
	return c.Repository.update(contact)
}

// Convert Contact to User
func (c *ContactService) Convert(contact *Contact, user *User) (User, error) {
	return c.Repository.convert(contact, user)
}

// Delete Contact
func (c *ContactService) Delete(contact *Contact) (Contact, error) {
	return c.Repository.delete(contact.ID)
}

// Get the address for a Contact in order to message them
func (c Contact) MessageAddress() MessageAddress {
	return MessageAddress{
		Type:   "contact",
		ID:     c.ID,
		Email:  c.Email,
		UserID: c.UserID,
	}
}

func (c Contact) String() string {
	return fmt.Sprintf("[intercom] contact { id: %s name: %s, user_id: %s, email: %s }", c.ID, c.Name, c.UserID, c.Email)
}
