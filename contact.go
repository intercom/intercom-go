package intercom

import "fmt"

// ContactService handles interactions with the API through a ContactRepository.
type ContactService struct {
	Repository ContactRepository
}

// ContactList holds a list of Contacts and paging information
type ContactList struct {
	Pages       PageParams
	Contacts    []Contact
	ScrollParam string `json:"scroll_param,omitempty"`
}

// Contact represents a Contact within Intercom.
// Not all fields are writeable to the API, non-writeable fields are
// stripped out from the request. Please see the API documentation for details.
type Contact struct {
	ID                     string                 `json:"id,omitempty"`
	Type                   string                 `json:"type,omitempty"`
	Email                  string                 `json:"email,omitempty"`
	Phone                  string                 `json:"phone,omitempty"`
	ExternalID             string                 `json:"external_id,omitempty"`
	Name                   string                 `json:"name,omitempty"`
	Avatar                 *ContactAvatar         `json:"avatar,omitempty"`
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

// LocationData represents the location for a Contact.
type LocationData struct {
	CityName      string  `json:"city_name,omitempty"`
	ContinentCode string  `json:"continent_code,omitempty"`
	CountryName   string  `json:"country_name,omitempty"`
	Latitude      float64 `json:"latitude,omitempty"`
	Longitude     float64 `json:"longitude,omitempty"`
	PostalCode    string  `json:"postal_code,omitempty"`
	RegionName    string  `json:"region_name,omitempty"`
	Timezone      string  `json:"timezone,omitempty"`
	CountryCode   string  `json:"country_code,omitempty"`
}

// SocialProfileList is a list SocialProfiles for a Contact.
type SocialProfileList struct {
	SocialProfiles []SocialProfile `json:"social_profiles,omitempty"`
}

// SocialProfile represents a social account for a Contact.
type SocialProfile struct {
	Name     string `json:"name,omitempty"`
	ID       string `json:"id,omitempty"`
	Username string `json:"username,omitempty"`
	URL      string `json:"url,omitempty"`
}

// ContactIdentifiers are used to identify contacts in Intercom.
type ContactIdentifiers struct {
	ID         string `url:"-"`
	ExternalID string `url:"external_id,omitempty"`
	Email      string `url:"email,omitempty"`
}

// ContactAvatar represents an avatar for a Contact.
type ContactAvatar struct {
	Type     string `json:"type,omitempty"`
	ImageURL string `json:"image_url,omitempty"`
}

type contactListParams struct {
	PageParams
	SegmentID string `url:"segment_id,omitempty"`
	TagID     string `url:"tag_id,omitempty"`
	Email     string `url:"email,omitempty"`
}

type scrollParams struct {
	ScrollParam string `url:"scroll_param,omitempty"`
}

// FindByID looks up a Contact by their Intercom ID.
func (c *ContactService) FindByID(id string) (Contact, error) {
	return c.findWithIdentifiers(ContactIdentifiers{ID: id})
}

// FindByUserID looks up a Contact by their ExternalID (automatically generated server side).
func (c *ContactService) FindByUserID(userID string) (Contact, error) {
	return c.findWithIdentifiers(ContactIdentifiers{ExternalID: userID})
}

func (c *ContactService) findWithIdentifiers(identifiers ContactIdentifiers) (Contact, error) {
	return c.Repository.find(identifiers)
}

// List all Contacts for App.
func (c *ContactService) List(params PageParams) (ContactList, error) {
	return c.Repository.list(contactListParams{PageParams: params})
}

// Scroll List all Contacts for App via Scroll API
func (c *ContactService) Scroll(scrollParam string) (ContactList, error) {
	return c.Repository.scroll(scrollParam)
}

// ListByEmail looks up a list of Contacts by their Email.
func (c *ContactService) ListByEmail(email string, params PageParams) (ContactList, error) {
	return c.Repository.list(contactListParams{PageParams: params, Email: email})
}

// ListBySegment lists Contacts by Segment.
func (c *ContactService) ListBySegment(segmentID string, params PageParams) (ContactList, error) {
	return c.Repository.list(contactListParams{PageParams: params, SegmentID: segmentID})
}

// ListByTag lists Contacts By Tag.
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

// Delete Contact
func (c *ContactService) Delete(contact *Contact) (Contact, error) {
	return c.Repository.delete(contact.ID)
}

// MessageAddress gets the address for a Contact in order to message them
func (c Contact) MessageAddress() MessageAddress {
	return MessageAddress{
		Type:       "contact",
		ID:         c.ID,
		Email:      c.Email,
		ExternalID: c.ExternalID,
	}
}

func (c Contact) String() string {
	return fmt.Sprintf("[intercom] contact { id: %s name: %s, external_id: %s, email: %s }", c.ID, c.Name, c.ExternalID, c.Email)
}
