package intercom

type TaggingList struct {
	Name      string    `json:"name,omitempty"`
	Users     []Tagging `json:"users,omitempty"`
	Companies []Tagging `json:"companies,omitempty"`
}

type Tagging struct {
	ID        string `json:"id,omitempty"`
	UserID    string `json:"user_id,omitempty"`
	Email     string `json:"email,omitempty"`
	CompanyID string `json:"company_id,omitempty"`
	Untag     *bool  `json:"untag,omitempty"`
}
