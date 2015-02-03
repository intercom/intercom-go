package intercom

import (
	"io/ioutil"
	"testing"
)

func TestCompanyAPIFind(t *testing.T) {
	http := TestCompanyHTTPClient{fixtureFilename: "fixtures/company.json", expectedURI: "/companies/54c42e7ea7a765fa7", t: t}
	api := CompanyAPI{httpClient: &http}
	company, err := api.find(CompanyIdentifiers{ID: "54c42e7ea7a765fa7"})
	if err != nil {
		t.Errorf("Error parsing fixture %s", err)
	}
	if company.ID != "54c42ed71623d8caa" {
		t.Errorf("ID was %s, expected 54c42ed71623d8caa", company.ID)
	}
	if company.CompanyID != "762" {
		t.Errorf("CompanyID was %s, expected 762", company.CompanyID)
	}
	if company.RemoteCreatedAt != 1413218536 {
		t.Errorf("RemoteCreatedAt was %d, expected %d", company.RemoteCreatedAt, 1413218536)
	}
	if company.CustomAttributes["big_company"] != true {
		t.Errorf("CustomAttributes was %v, expected %v", company.CustomAttributes, map[string]interface{}{"big_company": true})
	}
}

func TestCompanyAPIFindByName(t *testing.T) {
	http := TestCompanyHTTPClient{fixtureFilename: "fixtures/company.json", expectedURI: "/companies", t: t}
	api := CompanyAPI{httpClient: &http}
	company, _ := api.find(CompanyIdentifiers{Name: "Important Company"})
	if company.Name != "Important Company" {
		t.Errorf("Name was %s, expected Important Company", company.Name)
	}
}

func TestCompanyAPIListDefault(t *testing.T) {
	http := TestCompanyHTTPClient{fixtureFilename: "fixtures/companies.json", expectedURI: "/companies", t: t}
	api := CompanyAPI{httpClient: &http}
	companyList, _ := api.list(companyListParams{})
	companies := companyList.Companies
	if companies[0].ID != "54c42ed71623d8caa" {
		t.Errorf("ID was %s, expected 54c42ed71623d8caa", companies[0].ID)
	}
	pages := companyList.Pages
	if pages.Page != 1 {
		t.Errorf("Page was %d, expected 1", pages.Page)
	}
}

func TestCompanyAPISave(t *testing.T) {
	http := TestCompanyHTTPClient{t: t, expectedURI: "/companies"}
	api := CompanyAPI{httpClient: &http}
	company := Company{CompanyID: "27"}
	api.save(&company)
}

type TestCompanyHTTPClient struct {
	TestHTTPClient
	t               *testing.T
	fixtureFilename string
	expectedURI     string
}

func (t TestCompanyHTTPClient) Get(uri string, queryParams interface{}) ([]byte, error) {
	if t.expectedURI != uri {
		t.t.Errorf("URI was %s, expected %s", uri, t.expectedURI)
	}
	return ioutil.ReadFile(t.fixtureFilename)
}

func (t TestCompanyHTTPClient) Post(uri string, body interface{}) ([]byte, error) {
	if uri != "/companies" {
		t.t.Errorf("Wrong endpoint called")
	}
	return nil, nil
}
