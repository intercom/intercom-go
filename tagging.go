package intercom

// TaggingList is an object used to Tag Users and Companies.
// The Name should be that of the Tag required,
// and Users and Companies are lists of Taggings
type TaggingList struct {
	Name      string    `json:"name,omitempty"`
	Users     []Tagging `json:"users,omitempty"`
	Companies []Tagging `json:"companies,omitempty"`
}

// A Tagging is an object identifying a User or Company to be tagged,
// that can optionally be set to untag.
type Tagging struct {
	ID        string `json:"id,omitempty"`
	UserID    string `json:"user_id,omitempty"`
	Email     string `json:"email,omitempty"`
	CompanyID string `json:"company_id,omitempty"`
	Untag     *bool  `json:"untag,omitempty"`
}
