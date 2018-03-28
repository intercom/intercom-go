package intercom

import "fmt"

// VisitorService handles interactions with the API through a VisitorRepository.
type VisitorService struct {
	Repository VisitorRepository
}

// Visitor represents a Visitor within Intercom.
// Not all of the fields are writeable to the API, non-writeable fields are
// stripped out from the request. Please see the API documentation for details.
type Visitor struct {
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

// FindByID looks up a Visitor by their Intercom ID.
func (v *VisitorService) FindByID(id string) (Visitor, error) {
	return v.findWithIdentifiers(UserIdentifiers{ID: id})
}

// FindByUserID looks up a Visitor by their UserID (automatically generated server side).
func (v *VisitorService) FindByUserID(userID string) (Visitor, error) {
	return v.findWithIdentifiers(UserIdentifiers{UserID: userID})
}

func (v *VisitorService) findWithIdentifiers(identifiers UserIdentifiers) (Visitor, error) {
	return v.Repository.find(identifiers)
}

// Update Visitor
func (v *VisitorService) Update(visitor *Visitor) (Visitor, error) {
	return v.Repository.update(visitor)
}

// Convert Visitor to Lead
func (v *VisitorService) Convert(visitor *Visitor, lead *Lead) (Lead, error) {
	return v.Repository.convert(visitor, lead)
}

// Convert Visitor to User
func (v *VisitorService) Convert(visitor *Visitor, user *User) (User, error) {
	return v.Repository.convert(visitor, user)
}

// Delete Visitor
func (v *VisitorService) Delete(visitor *Visitor) (Visitor, error) {
	return v.Repository.delete(visitor.ID)
}

func (c Visitor) String() string {
	return fmt.Sprintf("[intercom] visitor { id: %s name: %s, user_id: %s }", c.ID, c.Name, c.UserID )
}
