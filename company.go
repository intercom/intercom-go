package intercom

import "fmt"

type CompanyService struct {
	Repository CompanyRepository
}

type CompanyList struct {
	Pages     PageParams
	Companies []Company
}

type Company struct {
	ID               string                 `json:"id,omitempty"`
	CompanyID        string                 `json:"company_id,omitempty"`
	Name             string                 `json:"name,omitempty"`
	RemoteCreatedAt  int32                  `json:"remote_created_at,omitempty"`
	LastRequestAt    int32                  `json:"last_request_at,omitempty"`
	CreatedAt        int32                  `json:"created_at,omitempty"`
	UpdatedAt        int32                  `json:"updated_at,omitempty"`
	SessionCount     int32                  `json:"session_count,omitempty"`
	MonthlySpend     int32                  `json:"monthly_spend,omitempty"`
	UserCount        int32                  `json:"user_count,omitempty"`
	Tags             *TagList               `json:"tags,omitempty"`
	Segments         *SegmentList           `json:"segments,omitempty"`
	Plan             *Plan                  `json:"plan,omitempty"`
	CustomAttributes map[string]interface{} `json:"custom_attributes,omitempty"`
	Remove           *bool                  `json:"-"`
}

type CompanyIdentifiers struct {
	ID        string `url:"-"`
	CompanyID string `url:"company_id,omitempty"`
	Name      string `url:"name,omitempty"`
}

type Plan struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type companyListParams struct {
	PageParams
	SegmentID string `url:"segment_id,omitempty"`
	TagID     string `url:"tag_id,omitempty"`
}

func (c *CompanyService) FindByID(id string) (Company, error) {
	return c.findWithIdentifiers(CompanyIdentifiers{ID: id})
}

func (c *CompanyService) FindByCompanyID(companyID string) (Company, error) {
	return c.findWithIdentifiers(CompanyIdentifiers{CompanyID: companyID})
}

func (c *CompanyService) FindByName(name string) (Company, error) {
	return c.findWithIdentifiers(CompanyIdentifiers{Name: name})
}

func (c *CompanyService) findWithIdentifiers(identifiers CompanyIdentifiers) (Company, error) {
	return c.Repository.find(identifiers)
}

func (c *CompanyService) List(params PageParams) (CompanyList, error) {
	return c.Repository.list(companyListParams{PageParams: params})
}

func (c *CompanyService) ListBySegment(segmentID string, params PageParams) (CompanyList, error) {
	return c.Repository.list(companyListParams{PageParams: params, SegmentID: segmentID})
}

func (c *CompanyService) ListByTag(tagID string, params PageParams) (CompanyList, error) {
	return c.Repository.list(companyListParams{PageParams: params, TagID: tagID})
}

func (c *CompanyService) Save(user *Company) (Company, error) {
	return c.Repository.save(user)
}

func (c Company) String() string {
	return fmt.Sprintf("[intercom] company { id: %s name: %s, company_id: %s }", c.ID, c.Name, c.CompanyID)
}

func (p Plan) String() string {
	return fmt.Sprintf("[intercom] company_plan { id: %s name: %s }", p.ID, p.Name)
}
