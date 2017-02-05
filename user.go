package intercom

import "fmt"

// UserService handles interactions with the API through a UserRepository.
type UserService struct {
	Repository UserRepository
}

// UserList holds a list of Users and paging information
type UserList struct {
	Pages PageParams
	Users []User
	ScrollParam string `json:"scroll_param,omitempty"`
}

// User represents a User within Intercom.
// Not all of the fields are writeable to the API, non-writeable fields are
// stripped out from the request. Please see the API documentation for details.
type User struct {
	ID                     string                 `json:"id,omitempty"`
	Email                  string                 `json:"email,omitempty"`
	Phone                  string                 `json:"phone,omitempty"`
	UserID                 string                 `json:"user_id,omitempty"`
	Anonymous              *bool                  `json:"anonymous,omitempty"`
	Name                   string                 `json:"name,omitempty"`
	Pseudonym              string                 `json:"pseudonym,omitempty"`
	Avatar                 *UserAvatar            `json:"avatar,omitempty"`
	LocationData           *LocationData          `json:"location_data,omitempty"`
	SignedUpAt             int64                  `json:"signed_up_at,omitempty"`
	RemoteCreatedAt        int64                  `json:"remote_created_at,omitempty"`
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
	LastSeenUserAgent      string                 `json:"last_seen_user_agent,omitempty"`
}

// LocationData represents the location for a User.
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

// SocialProfile list is a list of SocialProfiles for a User.
type SocialProfileList struct {
	SocialProfiles []SocialProfile `json:"social_profiles,omitempty"`
}

// SocialProfile represents a social account for a User.
type SocialProfile struct {
	Name     string `json:"name,omitempty"`
	ID       string `json:"id,omitempty"`
	Username string `json:"username,omitempty"`
	URL      string `json:"url,omitempty"`
}

// UserIdentifiers are used to identify Users in Intercom.
type UserIdentifiers struct {
	ID     string `url:"-"`
	UserID string `url:"user_id,omitempty"`
	Email  string `url:"email,omitempty"`
}

// UserAvatar represents an avatar for a User.
type UserAvatar struct {
	ImageURL string `json:"image_url,omitempty"`
}

type userListParams struct {
	PageParams
	SegmentID string `url:"segment_id,omitempty"`
	TagID     string `url:"tag_id,omitempty"`
}

type scrollParams struct {
	ScrollParam  string `url:"scroll_param,omitempty"`
}

// FindByID looks up a User by their Intercom ID.
func (u *UserService) FindByID(id string) (User, error) {
	return u.findWithIdentifiers(UserIdentifiers{ID: id})
}

// FindByUserID looks up a User by their UserID (customer supplied).
func (u *UserService) FindByUserID(userID string) (User, error) {
	return u.findWithIdentifiers(UserIdentifiers{UserID: userID})
}

// FindByEmail looks up a User by their Email.
func (u *UserService) FindByEmail(email string) (User, error) {
	return u.findWithIdentifiers(UserIdentifiers{Email: email})
}

func (u *UserService) findWithIdentifiers(identifiers UserIdentifiers) (User, error) {
	return u.Repository.find(identifiers)
}

// List all Users for App.
func (u *UserService) List(params PageParams) (UserList, error) {
	return u.Repository.list(userListParams{PageParams: params})
}

// List all Users for App via Scroll API
func (u *UserService) Scroll(scrollParam string) (UserList, error) {
       return u.Repository.scroll(scrollParam)
}

// List Users by Segment.
func (u *UserService) ListBySegment(segmentID string, params PageParams) (UserList, error) {
	return u.Repository.list(userListParams{PageParams: params, SegmentID: segmentID})
}

// List Users By Tag.
func (u *UserService) ListByTag(tagID string, params PageParams) (UserList, error) {
	return u.Repository.list(userListParams{PageParams: params, TagID: tagID})
}

// Save a User, creating or updating them.
func (u *UserService) Save(user *User) (User, error) {
	return u.Repository.save(user)
}

func (u *UserService) Delete(id string) (User, error) {
	return u.Repository.delete(id)
}

// Get the address for an User in order to message them
func (u User) MessageAddress() MessageAddress {
	return MessageAddress{
		Type:   "user",
		ID:     u.ID,
		Email:  u.Email,
		UserID: u.UserID,
	}
}

func (u User) String() string {
	return fmt.Sprintf("[intercom] user { id: %s name: %s, user_id: %s, email: %s }", u.ID, u.Name, u.UserID, u.Email)
}

func (l LocationData) String() string {
	return fmt.Sprintf("[intercom] location_data { city_name: %s country_name: %s }", l.CityName, l.CountryName)
}

func (s SocialProfile) String() string {
	return fmt.Sprintf("[intercom] social_profile { name: %s id: %s, username: %s }", s.Name, s.ID, s.Username)
}

func (a UserAvatar) String() string {
	return fmt.Sprintf("[intercom] user_avatar { image_url: %s }", a.ImageURL)
}
