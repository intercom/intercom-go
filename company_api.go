package intercom

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/intercom/intercom-go/interfaces"
)

type CompanyRepository interface {
	find(CompanyIdentifiers) (Company, error)
	list(companyListParams) (CompanyList, error)
	save(*Company) (Company, error)
}

type CompanyAPI struct {
	httpClient interfaces.HTTPClient
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

func (api CompanyAPI) save(company *Company) (Company, error) {
	savedCompany := Company{}
	data, err := api.httpClient.Post("/companies", company)
	if err != nil {
		return savedCompany, err
	}
	err = json.Unmarshal(data, &savedCompany)
	return savedCompany, err
}
