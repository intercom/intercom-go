package intercom

import (
	"testing"
)

func TestCompanyFindByID(t *testing.T) {
	company, _ := (&CompanyService{Repository: TestCompanyAPI{t: t}}).FindByID("46adad3f09126dca")
	if company.ID != "46adad3f09126dca" {
		t.Errorf("Company not found")
	}
}

func TestCompanyFindByName(t *testing.T) {
	company, _ := (&CompanyService{Repository: TestCompanyAPI{t: t}}).FindByName("My Co")
	if company.Name != "My Co" {
		t.Errorf("Company not found")
	}
}

func TestCompanyFindByCompanyID(t *testing.T) {
	company, _ := (&CompanyService{Repository: TestCompanyAPI{t: t}}).FindByCompanyID("134d")
	if company.CompanyID != "134d" {
		t.Errorf("Company not found")
	}
}

func TestCompanyList(t *testing.T) {
	companyList, _ := (&CompanyService{Repository: TestCompanyAPI{t: t}}).List(PageParams{})
	companies := companyList.Companies
	if companies[0].ID != "46adad3f09126dca" {
		t.Errorf("Company not listed")
	}
}

func TestCompanyListUsersByID(t *testing.T) {
	companyUserList, _ := (&CompanyService{Repository: TestCompanyAPI{t: t}}).ListUsersByID("46adad3f09126dca", PageParams{})
	users := companyUserList.Users
	if users[0].Companies.Companies[0].ID != "46adad3f09126dca" {
		t.Errorf("User not listed")
	}
}

func TestCompanyListUsersByCompanyID(t *testing.T) {
	companyUserList, _ := (&CompanyService{Repository: TestCompanyAPI{t: t}}).ListUsersByCompanyID("134d", PageParams{})
	users := companyUserList.Users
	if users[0].Companies.Companies[0].CompanyID != "134d" {
		t.Errorf("User not listed")
	}
}

func TestCompanySave(t *testing.T) {
	companyService := CompanyService{Repository: TestCompanyAPI{t: t}}
	company := Company{ID: "46adad3f09126dca", CustomAttributes: map[string]interface{}{"is_cool": true}}
	_, err := companyService.Save(&company)
	if err != nil {
		t.Errorf("Failed to save compant %s", err)
	}
}

type TestCompanyAPI struct {
	t *testing.T
}

func (t TestCompanyAPI) find(params CompanyIdentifiers) (Company, error) {
	return Company{ID: params.ID, Name: params.Name, CompanyID: params.CompanyID}, nil
}

func (t TestCompanyAPI) list(params companyListParams) (CompanyList, error) {
	return CompanyList{Companies: []Company{Company{ID: "46adad3f09126dca", Name: "My Co", CompanyID: "aa123"}}}, nil
}

func (t TestCompanyAPI) listUsers(id string, params companyUserListParams) (UserList, error) {
	return UserList{Users: []User{User{Companies: &CompanyList{Companies: []Company{Company{ID: id, CompanyID: params.CompanyID}}}}}}, nil
}

func (t TestCompanyAPI) scroll(scrollParam string) (CompanyList, error) {
	return CompanyList{Companies: []Company{Company{ID: "46adad3f09126dca", Name: "My Co", CompanyID: "aa123"}}}, nil
}

func (t TestCompanyAPI) save(company *Company) (Company, error) {
	if company.ID != "46adad3f09126dca" {
		t.t.Errorf("Company ID was %s, expected 46adad3f09126dca", company.ID)
	}
	expectedCAs := map[string]interface{}{"is_cool": true}
	if company.CustomAttributes["is_cool"] != expectedCAs["is_cool"] {
		t.t.Errorf("Custom attributes was %v, expected %v", company.CustomAttributes, expectedCAs)
	}
	return Company{}, nil
}
