package intercom

import "fmt"

// CompanyService handles interactions with the API through a CompanyRepository.
type CompanyService struct {
	Repository CompanyRepository
}

// CompanyList holds a list of Companies and paging information
type CompanyList struct {
	Pages       PageParams
	Companies   []Company
	ScrollParam string `json:"scroll_param,omitempty"`
}

// Company represents a Company in Intercom
// Not all of the fields are writeable to the API, non-writeable fields are
// stripped out from the request. Please see the API documentation for details.
type Company struct {
	ID               string                 `json:"id,omitempty"`
	CompanyID        string                 `json:"company_id,omitempty"`
	Name             string                 `json:"name,omitempty"`
	RemoteCreatedAt  int64                  `json:"remote_created_at,omitempty"`
	LastRequestAt    int64                  `json:"last_request_at,omitempty"`
	CreatedAt        int64                  `json:"created_at,omitempty"`
	UpdatedAt        int64                  `json:"updated_at,omitempty"`
	SessionCount     int64                  `json:"session_count,omitempty"`
	MonthlySpend     int64                  `json:"monthly_spend,omitempty"`
	UserCount        int64                  `json:"user_count,omitempty"`
	Tags             *TagList               `json:"tags,omitempty"`
	Segments         *SegmentList           `json:"segments,omitempty"`
	Plan             *Plan                  `json:"plan,omitempty"`
	CustomAttributes map[string]interface{} `json:"custom_attributes,omitempty"`
	Remove           *bool                  `json:"-"`
}

// CompanyIdentifiers to identify a Company using the API
type CompanyIdentifiers struct {
	ID        string `url:"-"`
	CompanyID string `url:"company_id,omitempty"`
	Name      string `url:"name,omitempty"`
}

// The Plan a Company is on
type Plan struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type companyListParams struct {
	PageParams
	SegmentID string `url:"segment_id,omitempty"`
	TagID     string `url:"tag_id,omitempty"`
}

// FindByID finds a Company using their Intercom ID
func (c *CompanyService) FindByID(id string) (Company, error) {
	return c.findWithIdentifiers(CompanyIdentifiers{ID: id})
}

// FindByCompanyID finds a Company using their CompanyID
// CompanyID is a customer-defined field
func (c *CompanyService) FindByCompanyID(companyID string) (Company, error) {
	return c.findWithIdentifiers(CompanyIdentifiers{CompanyID: companyID})
}

// FindByName finds a Company using their Name
func (c *CompanyService) FindByName(name string) (Company, error) {
	return c.findWithIdentifiers(CompanyIdentifiers{Name: name})
}

func (c *CompanyService) findWithIdentifiers(identifiers CompanyIdentifiers) (Company, error) {
	return c.Repository.find(identifiers)
}

// List Companies
func (c *CompanyService) List(params PageParams) (CompanyList, error) {
	return c.Repository.list(companyListParams{PageParams: params})
}

// List Companies by Segment
func (c *CompanyService) ListBySegment(segmentID string, params PageParams) (CompanyList, error) {
	return c.Repository.list(companyListParams{PageParams: params, SegmentID: segmentID})
}

// List Companies by Tag
func (c *CompanyService) ListByTag(tagID string, params PageParams) (CompanyList, error) {
	return c.Repository.list(companyListParams{PageParams: params, TagID: tagID})
}

// List all Companies for App via Scroll API
func (c *CompanyService) Scroll(scrollParam string) (CompanyList, error) {
	return c.Repository.scroll(scrollParam)
}

// Save a new Company, or update an existing one.
func (c *CompanyService) Save(user *Company) (Company, error) {
	return c.Repository.save(user)
}

func (c Company) String() string {
	return fmt.Sprintf("[intercom] company { id: %s name: %s, company_id: %s }", c.ID, c.Name, c.CompanyID)
}

func (p Plan) String() string {
	return fmt.Sprintf("[intercom] company_plan { id: %s name: %s }", p.ID, p.Name)
}
