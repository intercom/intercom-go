package intercom

import "fmt"

type UserService struct {
	Repository UserRepository
}

type UserList struct {
	Pages PageParams
	Users []User
}

type User struct {
	ID                     string                 `json:"id,omitempty"`
	Email                  string                 `json:"email,omitempty"`
	UserID                 string                 `json:"user_id,omitempty"`
	Name                   string                 `json:"name,omitempty"`
	Avatar                 *UserAvatar            `json:"avatar,omitempty"`
	LocationData           *LocationData          `json:"location_data,omitempty"`
	SignedUpAt             int32                  `json:"signed_up_at,omitempty"`
	RemoteCreatedAt        int32                  `json:"remote_created_at,omitempty"`
	LastRequestAt          int32                  `json:"last_request_at,omitempty"`
	CreatedAt              int32                  `json:"created_at,omitempty"`
	UpdatedAt              int32                  `json:"updated_at,omitempty"`
	SessionCount           int32                  `json:"session_count,omitempty"`
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

type SocialProfileList struct {
	SocialProfiles []SocialProfile `json:"social_profiles,omitempty"`
}

type SocialProfile struct {
	Name     string `json:"name,omitempty"`
	ID       string `json:"id,omitempty"`
	Username string `json:"username,omitempty"`
	URL      string `json:"url,omitempty"`
}

type UserIdentifiers struct {
	ID     string `url:"-"`
	UserID string `url:"user_id,omitempty"`
	Email  string `url:"email,omitempty"`
}

type UserAvatar struct {
	ImageURL string `json:"image_url,omitempty"`
}

type userListParams struct {
	PageParams
	SegmentID string `url:"segment_id,omitempty"`
	TagID     string `url:"tag_id,omitempty"`
}

func (u *UserService) FindByID(id string) (User, error) {
	return u.findWithIdentifiers(UserIdentifiers{ID: id})
}

func (u *UserService) FindByUserID(userID string) (User, error) {
	return u.findWithIdentifiers(UserIdentifiers{UserID: userID})
}

func (u *UserService) FindByEmail(email string) (User, error) {
	return u.findWithIdentifiers(UserIdentifiers{Email: email})
}

func (u *UserService) findWithIdentifiers(identifiers UserIdentifiers) (User, error) {
	return u.Repository.find(identifiers)
}

func (u *UserService) List(params PageParams) (UserList, error) {
	return u.Repository.list(userListParams{PageParams: params})
}

func (u *UserService) ListBySegment(segmentID string, params PageParams) (UserList, error) {
	return u.Repository.list(userListParams{PageParams: params, SegmentID: segmentID})
}

func (u *UserService) ListByTag(tagID string, params PageParams) (UserList, error) {
	return u.Repository.list(userListParams{PageParams: params, TagID: tagID})
}

func (u *UserService) Save(user *User) (User, error) {
	return u.Repository.save(user)
}

func (u *UserService) Delete(id string) (User, error) {
	return u.Repository.delete(id)
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
