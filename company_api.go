package intercom

import (
	"encoding/json"
	"errors"
	"fmt"

	"gopkg.in/intercom/intercom-go.v2/interfaces"
)

// CompanyRepository defines the interface for working with Companies through the API.
type CompanyRepository interface {
	find(CompanyIdentifiers) (Company, error)
	list(companyListParams) (CompanyList, error)
	scroll(scrollParam string) (CompanyList, error)
	save(*Company) (Company, error)
}

// CompanyAPI implements CompanyRepository
type CompanyAPI struct {
	httpClient interfaces.HTTPClient
}

type requestCompany struct {
	ID               string                 `json:"id,omitempty"`
	CompanyID        string                 `json:"company_id,omitempty"`
	Name             string                 `json:"name,omitempty"`
	RemoteCreatedAt  int64                  `json:"remote_created_at,omitempty"`
	MonthlySpend     int64                  `json:"monthly_spend,omitempty"`
	Plan             string                 `json:"plan,omitempty"`
	CustomAttributes map[string]interface{} `json:"custom_attributes,omitempty"`
}

func (api CompanyAPI) find(params CompanyIdentifiers) (Company, error) {
	company := Company{}
	data, err := api.getClientForFind(params)
	if err != nil {
		return company, err
	}
	err = json.Unmarshal(data, &company)
	return company, err
}

func (api CompanyAPI) getClientForFind(params CompanyIdentifiers) ([]byte, error) {
	switch {
	case params.ID != "":
		return api.httpClient.Get(fmt.Sprintf("/companies/%s", params.ID), nil)
	case params.CompanyID != "", params.Name != "":
		return api.httpClient.Get("/companies", params)
	}
	return nil, errors.New("Missing Company Identifier")
}

func (api CompanyAPI) list(params companyListParams) (CompanyList, error) {
	companyList := CompanyList{}
	data, err := api.httpClient.Get("/companies", params)
	if err != nil {
		return companyList, err
	}
	err = json.Unmarshal(data, &companyList)
	return companyList, err
}

func (api CompanyAPI) scroll(scrollParam string) (CompanyList, error) {
	companyList := CompanyList{}
	params := scrollParams{ScrollParam: scrollParam }
	data, err := api.httpClient.Get("/companies/scroll", params)
	if err != nil {
		return companyList, err
	}
	err = json.Unmarshal(data, &companyList)
	return companyList, err
}

func (api CompanyAPI) save(company *Company) (Company, error) {
	requestCompany := requestCompany{
		ID:               company.ID,
		Name:             company.Name,
		CompanyID:        company.CompanyID,
		RemoteCreatedAt:  company.RemoteCreatedAt,
		MonthlySpend:     company.MonthlySpend,
		Plan:             api.getPlanName(company),
		CustomAttributes: company.CustomAttributes,
	}

	savedCompany := Company{}
	data, err := api.httpClient.Post("/companies", &requestCompany)
	if err != nil {
		return savedCompany, err
	}
	err = json.Unmarshal(data, &savedCompany)
	return savedCompany, err
}

func (api CompanyAPI) getPlanName(company *Company) string {
	if company.Plan == nil {
		return ""
	}
	return company.Plan.Name
}
